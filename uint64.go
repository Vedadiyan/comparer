package comparer

import (
	"fmt"
	"strings"
)


type UInt64 uint64

func (i UInt64) Compare(v any) int {
	ref := uint64(i)
	switch t := v.(type) {
	case int:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case int32:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case int64:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case int16:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case int8:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case uint:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case uint64:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case uint32:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case uint16:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case byte:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case float32:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case float64:
		{
			v := uint64(t)
			if ref == v {
				return 0
			}
			if ref > v {
				return 1
			}
			return -1
		}
	case string:
		{
			return strings.Compare(fmt.Sprintf("%d", ref), t)
		}
	}
	return strings.Compare(fmt.Sprintf("%d", ref), fmt.Sprintf("%v", v))
}
