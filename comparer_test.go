package comparer

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		a any
		b any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Basic tests for different types
		{"Int Equal", args{10, 10}, 0},
		{"Int Greater", args{20, 10}, 1},
		{"Int Less", args{10, 20}, -1},
		{"Int32 Test", args{int32(10), 20}, -1},
		{"Int64 Test", args{int64(20), 10}, 1},
		{"Int16 Test", args{int16(10), 10}, 0},
		{"Int8 Test", args{int8(20), 10}, 1},
		{"Uint Test", args{uint(10), 20}, -1},
		{"Uint64 Test", args{uint64(20), 10}, 1},
		{"Uint32 Test", args{uint32(10), 10}, 0},
		{"Uint16 Test", args{uint16(20), 10}, 1},
		{"Byte Test", args{byte(10), 20}, -1},
		{"Float32 Test", args{float32(20.0), 10}, 1},
		{"Float64 Test", args{float64(10.0), 20}, -1},

		// Cross-type comparison tests
		{"Int vs Float64", args{10, 10.0}, 0},
		{"Int vs String", args{10, "10"}, 0},
		{"Float32 vs Int64", args{float32(10.0), int64(10)}, 0},

		// Edge cases
		{"Custom Type", args{struct{ val int }{10}, "test"}, strings.Compare(fmt.Sprintf("%v", struct{ val int }{10}), "test")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Template for type-specific test function
func createTestCases[T any](_ T) []struct {
	name string
	v    any
	want int
} {
	return []struct {
		name string
		v    any
		want int
	}{
		{"vs Int", 10, 0},
		{"vs Int32", int32(10), 0},
		{"vs Int64", int64(10), 0},
		{"vs Int16", int16(10), 0},
		{"vs Int8", int8(10), 0},
		{"vs Uint", uint(10), 0},
		{"vs Uint64", uint64(10), 0},
		{"vs Uint32", uint32(10), 0},
		{"vs Uint16", uint16(10), 0},
		{"vs Byte", byte(10), 0},
		{"vs Float32", float32(10.0), 0},
		{"vs Float64", float64(10.0), 0},
		{"vs String", "10", 0},
		{"Greater Than", 5, 1},
		{"Less Than", 15, -1},
		{"Less Than", 15, -1},
		{"vs any", []string{}, -1},
	}
}

func TestInt(t *testing.T) {
	i := int(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Int.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	i := int32(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Int32.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	i := int64(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Int64.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	i := int16(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Int16.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	i := int8(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Int8.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	i := uint(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("UInt.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	i := uint64(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("UInt64.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	i := uint32(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("UInt32.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	i := uint16(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("UInt16.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte(t *testing.T) {
	i := byte(10)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Byte.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	i := float32(10.0)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Float32.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	i := float64(10.0)
	for _, tt := range createTestCases(i) {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(i, tt.v); got != tt.want {
				t.Errorf("Float64.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Optional: Add benchmark tests
func BenchmarkCompare(b *testing.B) {
	benchmarks := []struct {
		name string
		a    any
		b    any
	}{
		{"Int", 10, 20},
		{"Float64", 10.0, 20.0},
		{"String", "10", "20"},
		{"Cross Type", 10, 20.0},
		{"Custom Type", struct{ val int }{10}, struct{ val int }{20}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Compare(bm.a, bm.b)
			}
		})
	}
}
