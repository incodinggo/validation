package valid

import (
	"reflect"
	"strings"
)

type validateRequired struct {
	validateModel
}

//Judgments based:
//If it is the following value type
//[Int,Int8,Int16,Int32,Int64,Uint,Uint8,Uint16,Uint32,Uint64,Uintptr,Float32,Float64,Complex64,Complex128] Trigger by value 0
//[string] Trigger by length 0 (after trim space)
//[ptr,slice,array] Trigger by nil
func (m *validateRequired) validate() (result bool) {
	switch m.fieldT.Type.Kind() {
	case reflect.String:
		result = len(strings.TrimSpace(m.fieldV.(string))) > 0
	case reflect.Ptr:
		result = !reflect.ValueOf(m.fieldV).IsNil()
	default:
		result = !reflect.ValueOf(m.fieldV).IsZero()
	}
	return result
}
