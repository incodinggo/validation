package valid

import (
	"fmt"
	"reflect"
	"strconv"
)

type validateArr struct {
	validateModel
	condition string
}

func (m *validateArr) validate() (result bool) {
	roles := getRegStrIntValue(m.condition)
	if len(roles) < 2 {
		fmt.Println("Incorrect Expression:", m.fieldT.Name, " -> ", m.condition)
		return
	}

	switch m.fieldT.Type.Kind() {
	case reflect.Ptr:
		m.fieldE = m.fieldE.Elem()
		result = m.Arr(roles[0], roles[1])
	default:
		result = m.Arr(roles[0], roles[1])
	}
	return
}

func (m *validateArr) Arr(min, max string) (result bool) {
	result = true
	switch m.fieldE.Kind() {
	case reflect.Slice:
		l := m.fieldE.Len()
		if min != "-" {
			c, err := strconv.ParseInt(min, 10, 64)
			if err != nil {
				fmt.Println("need [int:int] or [int],but get string")
				return
			}
			result = result && (int64(l) >= c)
		}
		if max != "-" {
			c, err := strconv.ParseInt(max, 10, 64)
			if err != nil {
				fmt.Println("need [int:int] or [int],but get string")
				return
			}
			result = result && (int64(l) <= c)
		}
	default:
		fmt.Printf("Check Tag [arr] Unsupported Param %v.(%v)\n", m.fieldT.Name, m.fieldT.Type)
		return true
	}
	return
}
