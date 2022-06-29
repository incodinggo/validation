package valid

import (
	"fmt"
	"reflect"
	"strings"
)

type validateMax struct {
	validateModel
	condition string
}

func (m *validateMax) validate() (result bool) {
	regValues := getRegIntValue(m.condition)
	var max int64
	if len(regValues) > 0 {
		max = regValues[0]
	}
	switch m.fieldT.Type.Kind() {
	case reflect.Ptr:
		m.fieldE = m.fieldE.Elem()
		result = m.Max(max)
	default:
		result = m.Max(max)
	}

	return
}

func (m *validateMax) Max(min int64) (result bool) {
	switch m.fieldE.Kind() {
	case reflect.String:
		val := m.fieldE.String()
		vLen := len(strings.TrimSpace(val))
		result = vLen <= int(min)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := m.fieldE.Int()
		result = val <= min
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := m.fieldE.Uint()
		result = val <= uint64(min)
	case reflect.Float32, reflect.Float64:
		val := m.fieldE.Float()
		result = val <= float64(min)
	default:
		result = true
		fmt.Printf("Check Tag [max] Unsupported Param %v.(%v) With Value [%v]\n", m.fieldT.Name, m.fieldT.Type, m.fieldV)
	}
	return result
}
