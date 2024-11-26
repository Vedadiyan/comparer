package comparer

import (
	"fmt"
	"strings"
)

type Byte byte

func (i Byte) Compare(v any) int {
	ref := byte(i)
	switch t := v.(type) {
	case int:
		{
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
			v := byte(t)
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
