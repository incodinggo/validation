package valid

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	required = "required"
	enum     = "enum"
	min      = "min"
	max      = "max"
	rang     = "range"
	regex    = "regex"
	arr      = "arr"
)

// ValidateError 验证错误提示
type ValidateError struct {
	Field   string // 验证不通过的属性
	Valid   string // 不通过的条件
	ErrCode int    // 不通过的错误码
}

func (ve ValidateError) Error() string {
	return fmt.Sprintf("Unverified <%s>:%s", ve.Field, ve.Valid)
}

// validateInterface 校验封装
type validateInterface interface {
	validate() bool
}

type validateModel struct {
	required bool
	fieldT   reflect.StructField
	fieldV   interface{}
	fieldE   reflect.Value
	error    int64
}

func Check(reqModel interface{}) error {

	typeV := reflect.ValueOf(reqModel)
	if typeV.Kind() == reflect.Ptr {
		typeV = typeV.Elem()
	}
	typeT := typeV.Type()

	return validateCheck(typeT, typeV)
}

func validateCheck(typeT reflect.Type, typeV reflect.Value) error {
	for i := 0; i < typeT.NumField(); i++ {
		fieldT := typeT.Field(i)
		fieldV := typeV.Field(i)
		if !fieldV.CanInterface() {
			continue
		}

		//如果是匿名结构体,需要递归判断
		if fieldT.Anonymous && fieldT.Type.Kind() == reflect.Struct {
			err := validateCheck(fieldT.Type, fieldV)
			if err != nil {
				return err
			}
			continue
		}

		// 是否存在校验字段
		validCond := fieldT.Tag.Get("valid")
		if len(validCond) == 0 {
			continue
		}

		// 如果校验出错，直接返回。不需要判断所有条件
		if err := validate(validCond, fieldT, typeV.FieldByName(fieldT.Name).Interface(), fieldV); err != nil {
			return err
		}
	}
	return nil
}

func validate(validCond string, fieldT reflect.StructField, fieldV interface{}, fieldE reflect.Value) error {

	// 是否必须
	_validateModel := validateModel{fieldT: fieldT, fieldV: fieldV, fieldE: fieldE}

	validSlice := strings.Split(validCond, ";")
	if strings.Index(validCond, "|") != -1 {
		validSlice = append(validSlice, strings.Split(validCond, "|")...)
	}

	for _, valid := range validSlice {
		slice := strings.Split(valid, "#")
		v := valid
		e := ""
		if len(slice) > 1 {
			v = slice[0]
			e = slice[1]
		}

		if len(v) == 0 {
			continue
		}

		var valid validateInterface
		num, err := strconv.ParseFloat(e, 64)
		if err != nil {
			err = nil
			num = 0
		}
		if strings.Index(v, required) == 0 {
			// 必填
			_validateModel.required = true
			valid = &validateRequired{validateModel: _validateModel}
		} else if strings.Index(v, rang) == 0 {
			// range
			valid = &validateRange{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, min) == 0 {
			// min
			valid = &validateMin{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, max) == 0 {
			// max
			valid = &validateMax{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, enum) == 0 {
			// enum
			valid = &validateEnum{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, regex) == 0 {
			// regex
			valid = &validateRegex{condition: v, validateModel: _validateModel}
		} else if strings.Index(v, arr) == 0 {
			// array/slice
			valid = &validateArr{condition: v, validateModel: _validateModel}
		} else {
			fmt.Println("unsupported syntax:", v)
			continue
		}

		if !valid.validate() {
			return ValidateError{
				Field:   fieldT.Name,
				Valid:   validCond,
				ErrCode: int(num),
			}
		}
	}

	return nil
}

func getRegIntValue(cond string) (values []int64) {
	reg, _ := regexp.Compile(`[-0-9]+`)
	regs := reg.FindAllString(cond, -1)

	for _, v := range regs {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Println("need [int:int] or [int],but get string")
			values = append(values, -1)
			continue
		}
		values = append(values, value)
	}
	return
}

func getRegStrIntValue(cond string) (values []string) {
	reg, _ := regexp.Compile(`[-0-9]+`)
	regs := reg.FindAllString(cond, -1)
	for _, v := range regs {
		if v == "-" {
			values = append(values, v)
			continue
		}
		values = append(values, v)
	}
	return
}
