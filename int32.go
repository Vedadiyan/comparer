package comparer

import (
	"fmt"
	"strings"
)


type Int32 int32

func (i Int32) Compare(v any) int {
	ref := int32(i)
	switch t := v.(type) {
	case int:
		{
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
			v := int32(t)
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
