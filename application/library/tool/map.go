package tool

import (
	"encoding/json"
	"strconv"
)

func MapGetValString(m map[string]interface{}, k, defaultV string) string {
	if m != nil {
		if v, exist := m[k]; exist {
			if vs, ok := v.(string); ok {
				return vs
			}
		}
	}

	return defaultV
}

func MapGetValMap(m map[string]interface{}, k string, defaultV map[string]interface{}) map[string]interface{} {
	if v, exist := m[k]; exist {
		if vm, ok := v.(map[string]interface{}); ok {
			return vm
		}
	}

	return defaultV
}

func MapGetValInt64(m map[string]interface{}, k string, defaultV int64) int64 {
	if v, exist := m[k]; exist {
		switch vc := v.(type) {
		case string:
			if uidInt, err := strconv.ParseInt(vc, 10, 64); err == nil {
				return uidInt
			}
		case int:
			return int64(vc)
		case int64:
			return vc
		case float64:
			return int64(vc)
		case json.Number:
			if vint, err := strconv.Atoi(string(vc)); err == nil {
				return int64(vint)
			}
		}
	}

	return defaultV
}

func MapGetValInt(m map[string]interface{}, k string, defaultV int) int {
	if v, exist := m[k]; exist {
		switch vc := v.(type) {
		case string:
			if vint, err := strconv.Atoi(vc); err == nil {
				return vint
			}
		case int:
			return vc
		case int64:
			return int(vc)
		case float64:
			return int(vc)
		case json.Number:
			if vint, err := strconv.Atoi(string(vc)); err == nil {
				return vint
			}
		}
	}

	return defaultV
}
