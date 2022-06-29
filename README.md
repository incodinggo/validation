Validation 校验
----
## 校验使用说明  Instructions
### 使用方式  How to use
**作用于struct的字段，使用时添加想要使用的验证方式的字段、参数，到待验证字段tag。**  
**Acting on the fields of the struct, add the fields and parameters of the verification method you want to use when using them, and add them to the field tag to be verified.**  
### 示例  Example
```go
type Inc struct {
    IA int `valid:"required#1"`
}

type Example struct {
    Inc
    A int               `valid:"required#1"`
    B int8              `valid:"min[10]#2"`
    C int16             `valid:"max[1024]#3"`
    D int32             `valid:"required#1;min[10]#2"`
    E int64             `valid:"required#1;max[1024]#3"`
    F string            `valid:"required#1;range[1,99]#4"`
    G *string           `valid:"required#1;range[0:8]#4"`
    H *int              `valid:"required#1;range[0:8]#5"`
    I []int             `valid:"required#1"`    //unsupported others method
    J *[]int            `valid:"required#1"`    //unsupported others method
    K []string          `valid:"required#1"`    //unsupported others method
    L []*string         `valid:"required#1"`    //unsupported others method
    M struct{ M1 int }  `valid:"required#1"`    //unsupported others method
    N *struct{ M1 int } `valid:"required#1"`    //unsupported others method
}
```
`valid`  
**go tag 验证识别符**  
*go tag verification identifier*  

`required;range...`  
**使用的验证方式(参考下方支持校验方式及说明表)，多个验证方式使用;分割**  
*The used verification method (refer to the support verification method and description table below), use ‘;’ to split multiple authentication methods*  

`required#1234`  
**方式#错误码，如果当前条件被触发结果会返回该错误码**  
*Method#ErrorCode, which will be returned if current condition is on Trigger*  

`range[1,10]`  
**方式[范围] 或 方式[值] 使用该方式指定某些方式的范围、固定值和正则表达式等**
*mode[range] or mode[value] Use this mode to specify ranges, fixed values, regular expressions, etc. for certain modes*
```go
var reqStr = `{"IA":1,"A":1025,"B":-127,"C":1,"D":1,"E":1,"F":"string1","G":"string2","H":1,"I":[],"J":[],"K":[],"L":[],"M":{"M1":1},"N":{}}`
err := json.Unmarshal([]byte(reqStr), &req)
if err != nil {
	panic(err)
}
err = valid.Check(req)
if err != nil {
	ve := err.(valid.ValidateError)
	fmt.Println(ve.Field, "->", ve.Valid, ve.ErrCode)
	return
}
```
## 支持校验方式及说明 Support verification methods and descriptions

| 校验方式     | 适用字段类型                               | 样式                          | 说明                                                                                  |
|:---------|:-------------------------------------|:----------------------------|:------------------------------------------------------------------------------------|
| required | string number struct slice array ptr | required                    | 验证非0值或非空字符串或指针nil,如果可能包含0值或空字符串,可使用ptr                                              |                 
| enum     | string number ptr                    | enum[1,2,3]<br/>enum[a,B,c] | 输入在指定的值中                                                                            |           
| min      | string number ptr                    | min[10]                     | string最小长度(会去除前后空白进行验证)<br/>number最小值                                               |             
| max      | string number ptr                    | max[1024]                   | string最大长度(会去除前后空白进行验证)<br/>number最大值                                               |                 
| range    | string number ptr                    | range[0:10]                 | 长度在指定的范围内,会去除前后空白进行验证                                                               |
| regex    | string ptr(*string)                  | regex[pwd]<br/>regex[自定义正则] | 验证是否符合正则规范可以自定义或使用预制:<br/>pwd，pwdH，pwdHS，email，phone，phoneCN，phone，id，dt，date，ts，ip |## 支持校验方式及说明

| Validation Method | Applicable Field Type                | Style                                | Description                                                                                                                                             |
|:------------------|:-------------------------------------|:-------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------|
| required          | string number struct slice array ptr | required                             | validate non-zero value or non-empty string or pointer nil, if it may contain 0 value or empty string, use ptr                                          |                 
| enum              | string number ptr                    | enum[1,2,3]<br/>enum[a,B,c]          | Enter in the specified value                                                                                                                            |           
| min               | string number ptr                    | min[10]                              | string: Minimum length (the leading and trailing blanks will be removed for verification)<br/>number: minimum                                           |             
| max               | string number ptr                    | max[1024]                            | string: Maximum length (the leading and trailing blanks will be removed for verification)<br/>number: maximum                                           |                 
| range             | string number ptr                    | range[0:10]                          | If the length is within the specified range, the leading and trailing blanks will be removed for verification.                                          |
| regex             | string ptr(*string)                  | regex[pwd]<br/>regex[custom regular] | Validation for compliance with regular specifications can be customized or prefabricated:<br/>pwd，pwdH，pwdHS，email，phone，phoneCN，phone，id，dt，date，ts，ip |