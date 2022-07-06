package valid

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

//some simple RE
var regexpMap map[string]func(string) bool

func init() {
	regexpMap = map[string]func(string) bool{
		"pwd":     regexp.MustCompile(`^[\w_.,]+$`).MatchString,                                                                         //简单密码(只能包含字母、数字和下划线)
		"email":   regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`).MatchString,                                      //邮箱
		"phone":   regexp.MustCompile(`^[0-9]{9,13}$`).MatchString,                                                                      //手机号
		"phoneCN": regexp.MustCompile(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`).MatchString,       //手机号(china)
		"phone+":  regexp.MustCompile(`^\+[0-9]{9,13}$`).MatchString,                                                                    //手机号带+0112345678901
		"id":      regexp.MustCompile(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`).MatchString,                                           //证件号
		"dt":      regexp.MustCompile(`^\d{4}(-)(1[0-2]|0?\d)(-)([0-2]\d|0+\d|30|31)\s+(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$`).MatchString, //日期时间2006-01-02 15:04:05
		"date":    regexp.MustCompile(`^\d{4}(-)(1[0-2]|0?\d)(-)([0-2]\d|0+\d|30|31)$`).MatchString,                                     //日期2006-01-02
		"ts":      regexp.MustCompile(`^[a-zA-Z]+/[a-zA-Z]+$`).MatchString,                                                              //时区Asia/Shanghai
		"ip":      regexp.MustCompile(`((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}`).MatchString,      //IP
	}
}

type validateRegex struct {
	validateModel
	condition string
}

func (m *validateRegex) validate() (result bool) {
	if strings.EqualFold(m.condition, "regex[]") {
		fmt.Println("Incorrect Expression:", m.fieldT.Name, m.condition)
		return
	}
	cond := getCondition(m.condition)
	switch m.fieldT.Type.Kind() {
	case reflect.Ptr:
		m.fieldE = m.fieldE.Elem()
		result = m.Regex(cond)
	default:
		result = m.Regex(cond)
	}
	return
}

func (m *validateRegex) Regex(cond string) (result bool) {
	switch m.fieldE.Kind() {
	case reflect.String:
		val := strings.TrimSpace(m.fieldE.String())
		re, ok := regexpMap[cond]
		if !ok {
			re = regexp.MustCompile(cond).MatchString
		}
		ok = re(val)
		if ok {
			result = true
		}
	default:
		result = true
		fmt.Printf("Check Tag [regex] Unsupported Param %v.(%v)\n", m.fieldT.Name, m.fieldT.Type)
	}
	return result
}

//^[\w_.,]+$
func getCondition(cond string) (value string) {
	value = strings.TrimLeft(cond, "regex[")
	value = strings.TrimRight(value, "]")
	return
}
