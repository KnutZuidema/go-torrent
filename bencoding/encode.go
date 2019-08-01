package bencoding

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Marshaler interface {
	MarshalBEncode() ([]byte, error)
}

func Marshal(value interface{}) ([]byte, error) {
	value = normalize(value)
	switch v := value.(type) {
	case Marshaler:
		return v.MarshalBEncode()
	case string:
		return encodeString(v)
	case int64:
		return encodeInt(v)
	case uint64:
		return encodeUint(v)
	case bool:
		if v {
			return encodeInt(1)
		}
		return encodeInt(0)
	default:
		rv := dereference(reflect.ValueOf(v))
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			var list []interface{}
			for i := 0; i < rv.Len(); i++ {
				list = append(list, rv.Index(i).Interface())
			}
			return encodeList(list)
		case reflect.Map:
			if rv.Type().Key().Kind() != reflect.String {
				return nil, ErrNonStringKey
			}
			dict := make(map[string]interface{})
			iter := rv.MapRange()
			for iter.Next() {
				dict[iter.Key().Interface().(string)] = iter.Value().Interface()
			}
			return encodeDict(dict)
		case reflect.Struct:
			return encodeDict(structToMap(rv))
		default:
			return nil, ErrInvalidType
		}
	}
}

func encodeString(value string) ([]byte, error) {
	return []byte(strconv.Itoa(len(value)) + string(StringSeparatorToken) + value), nil
}

func encodeInt(value int64) ([]byte, error) {
	return []byte(string(IntToken) + strconv.FormatInt(value, 10) + string(EndToken)), nil
}

func encodeUint(value uint64) ([]byte, error) {
	return []byte(string(IntToken) + strconv.FormatUint(value, 10) + string(EndToken)), nil
}

func encodeList(value []interface{}) ([]byte, error) {
	res := []byte(string(ListToken))
	for _, val := range value {
		d, err := Marshal(val)
		if err != nil {
			return nil, err
		}
		res = append(res, d...)
	}
	res = append(res, EndToken)
	return res, nil
}

func encodeDict(value map[string]interface{}) ([]byte, error) {
	res := []byte(string(DictToken))
	type keyValuePair struct {
		key   string
		value interface{}
	}
	var sorted []keyValuePair
	for key, val := range value {
		sorted = append(sorted, keyValuePair{
			key:   key,
			value: val,
		})
	}
	sort.Slice(sorted, func(i, j int) bool {
		if len(sorted[i].key) == len(sorted[j].key) {
			for k := 0; k < len(sorted[i].key); k++ {
				if sorted[i].key[k] != sorted[j].key[k] {
					return sorted[i].key[k] < sorted[j].key[k]
				}
			}
			return true
		}
		return len(sorted[i].key) < len(sorted[j].key)
	})
	for _, pair := range sorted {
		key, _ := encodeString(pair.key)
		val, err := Marshal(pair.value)
		if err != nil {
			return nil, err
		}
		res = append(res, append(key, val...)...)
	}
	res = append(res, EndToken)
	return res, nil
}

type tag struct {
	Name      string
	OmitEmpty bool
	Skip      bool
}

func newTag(f reflect.StructField) *tag {
	v := f.Tag.Get(StructTagKey)
	if v == "" {
		return &tag{
			Name: f.Name,
		}
	}
	split := strings.Split(v, ",")
	if split[0] == "-" {
		return &tag{
			Skip: true,
		}
	}
	t := &tag{}
	if split[0] == "" {
		t.Name = f.Name
	}
	for _, option := range split[1:] {
		if option == OptionOmitEmpty {
			t.OmitEmpty = true
		}
	}
	return t
}

func structToMap(rv reflect.Value) map[string]interface{} {
	dict := make(map[string]interface{})
	typ := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		if typ.Field(i).PkgPath != "" {
			continue
		}
		t := newTag(typ.Field(i))
		if t.Skip {
			continue
		}
		if t.OmitEmpty && isEmpty(rv) {
			continue
		}
		if typ.Field(i).Anonymous && t.Name == "" {
			for k, v := range structToMap(rv.Field(i)) {
				dict[k] = v
			}
		}
		v := rv.Field(i).Interface()
		dict[t.Name] = v
	}
	return dict
}

func isEmpty(rv reflect.Value) bool {
	return reflect.DeepEqual(rv.Interface(), reflect.Zero(rv.Type()).Interface())
}

func normalize(i interface{}) interface{} {
	switch v := i.(type) {
	case []byte:
		return string(v)
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	}
	return i
}