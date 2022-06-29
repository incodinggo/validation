package valid

import (
	"fmt"
	"reflect"
	"strings"
)

type validateRange struct {
	validateModel
	condition string
}

func (m *validateRange) validate() (result bool) {
	if strings.Contains(m.condition, "[:") || strings.Contains(m.condition, ":]") {
		fmt.Println("Incorrect Expression:", m.fieldT.Name, m.condition)
		return
	}

	regValues := getRegIntValue(m.condition)
	var min, max = regValues[0], regValues[1]

	switch m.fieldT.Type.Kind() {
	case reflect.Ptr:
		m.fieldE = m.fieldE.Elem()
		result = m.Range(min, max)
	default:
		result = m.Range(min, max)
	}

	return result
}

func (m *validateRange) Range(min, max int64) (result bool) {
	switch m.fieldE.Kind() {
	case reflect.String:
		val := m.fieldE.String()
		vLen := len(strings.TrimSpace(val))
		result = vLen >= int(min) && vLen <= int(max)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val := m.fieldE.Int()
		result = val >= min && val <= max
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := m.fieldE.Uint()
		result = val >= uint64(min) && val <= uint64(max)
	case reflect.Float32, reflect.Float64:
		val := m.fieldE.Float()
		result = val >= float64(min) && val <= float64(max)
	default:
		result = true
		fmt.Printf("Check Tag [rang] Unsupported Param %v.(%v) With Value [%v]\n", m.fieldT.Name, m.fieldT.Type, m.fieldV)
	}
	return result
}
