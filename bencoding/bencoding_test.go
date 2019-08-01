package bencoding

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestBencoding(t *testing.T) {
	if err := quick.Check(func(i int) bool {
		return t.Run("test int", func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, new(int)); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i string) bool {
		return t.Run("test string", func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, new(string)); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i []int) bool {
		return t.Run("test list", func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, new([]int)); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i map[string][]int) bool {
		return t.Run("test dict", func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, new(map[string][]int)); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	type testStruct struct {
		ValA int    `bencode:"val_a,omitempty"`
		ValB string `bencode:"val_b"`
		ValC []int
		ValD map[string][]int
	}
	if err := quick.Check(func(i testStruct) bool {
		return t.Run("test struct", func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, new(testStruct)); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i bool) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i int8) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i int16) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i int32) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i int64) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i uint) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i uint8) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i uint16) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i uint32) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i uint64) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, &i); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	if err := quick.Check(func(i []byte) bool {
		return t.Run(fmt.Sprintf("test %T", i), func(t *testing.T) {
			var data []byte
			var err error
			if data, err = Marshal(i); err != nil {
				t.Error(err)
			}
			if err := Unmarshal(data, new([]byte)); err != nil {
				t.Error(err)
			}
		})
	}, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	t.Run("test float32", func(t *testing.T) {
		if _, err := Marshal(float32(1.1)); err == nil {
			t.Error("float32 should return error on Marshal")
		}
		if err := Unmarshal([]byte("asd"), new(float32)); err == nil {
			t.Error("float32 should return error on Marshal")
		}
	})
	t.Run("test float64", func(t *testing.T) {
		if _, err := Marshal(float64(1.1)); err == nil {
			t.Error("float64 should return error on Marshal")
		}
		if err := Unmarshal([]byte("asd"), new(float64)); err == nil {
			t.Error("float64 should return error on Unmarshal")
		}
	})
	t.Run("test complex64", func(t *testing.T) {
		if _, err := Marshal(complex64(1)); err == nil {
			t.Error("complex64 should return error on Marshal")
		}
		if err := Unmarshal([]byte("asd"), new(complex64)); err == nil {
			t.Error("complex64 should return error on Unmarshal")
		}
	})
	t.Run("test complex128", func(t *testing.T) {
		if _, err := Marshal(complex128(1)); err == nil {
			t.Error("complex128 should return error on Marshal")
		}
		if err := Unmarshal([]byte("asd"), new(complex128)); err == nil {
			t.Error("complex128 should return error on Unmarshal")
		}
	})
	t.Run("nil data", func(t *testing.T) {
		if _, err := Marshal(nil); err == nil {
			t.Error("nil data should return error")
		}
		if err := Unmarshal(nil, new(int)); err == nil {
			t.Error("nil data should return error")
		}
		if err := Unmarshal([]byte("aa"), nil); err == nil {
			t.Error("nil data should return error")
		}
	})
}
