package comparer

import (
	"fmt"
	"strings"
)

func Compare(a, b any) int {
	switch t := a.(type) {
	case int:
		{
			return Int(t).Compare(b)
		}
	case int32:
		{
			return Int32(t).Compare(b)
		}
	case int64:
		{
			return Int64(t).Compare(b)
		}
	case int16:
		{
			return Int16(t).Compare(b)
		}
	case int8:
		{
			return Int8(t).Compare(b)
		}
	case uint:
		{
			return UInt(t).Compare(b)
		}
	case uint64:
		{
			return UInt64(t).Compare(b)
		}
	case uint32:
		{
			return UInt32(t).Compare(b)
		}
	case uint16:
		{
			return UInt16(t).Compare(b)
		}
	case byte:
		{
			return Byte(t).Compare(b)
		}
	case float32:
		{
			return Float32(t).Compare(b)
		}
	case float64:
		{
			return Float64(t).Compare(b)
		}
	default:
		{
			return strings.Compare(fmt.Sprintf("%v", a), fmt.Sprintf("%v", b))
		}
	}
}
