package valid

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type validateEnum struct {
	validateModel
	condition string
}

func (m *validateEnum) validate() (result bool) {
	// 汉字 字母 数字
	reg, _ := regexp.Compile(`[\p{Han}\w]*[^,\[\]]`)
	conds := reg.FindAllString(m.condition, -1)
	if len(conds) < 2 {
		fmt.Println("Incorrect Expression:", m.fieldT.Name, " -> ", m.condition)
		return
	}

	switch m.fieldT.Type.Kind() {
	case reflect.Ptr:
		m.fieldE = m.fieldE.Elem()
		result = m.Enum(conds)
	default:
		result = m.Enum(conds)
	}
	return
}

func (m *validateEnum) Enum(conds []string) (result bool) {
	var val string
	switch m.fieldE.Kind() {
	case reflect.String:
		val = strings.TrimSpace(m.fieldE.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val = fmt.Sprint(m.fieldE.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val = fmt.Sprint(m.fieldE.Uint())
	case reflect.Float32, reflect.Float64:
		val = fmt.Sprint(m.fieldE.Float())
	default:
		fmt.Printf("Check Tag [enum] Unsupported Param %v.(%v)\n", m.fieldT.Name, m.fieldT.Type)
		return true
	}
	for _, cond := range conds {
		if cond == val {
			return true
		}
	}
	return
}
