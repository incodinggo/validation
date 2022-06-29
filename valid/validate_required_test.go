package valid_test

import (
	"encoding/json"
	"fmt"
	"github.com/incodinggo/validation/valid"
	"testing"
)

func TestRequired(t *testing.T) {
	type Inc struct {
		IA bool `valid:"required#1"`
	}

	var req struct {
		Inc
		A int               `valid:"required#1"`
		B int8              `valid:"required#1"`
		C int16             `valid:"required#1"`
		D int32             `valid:"required#1"`
		E int64             `valid:"required#1"`
		F string            `valid:"required#1"`
		G *string           `valid:"required#1"`
		H *int              `valid:"required#1"`
		I []int             `valid:"required#1"`
		J *[]int            `valid:"required#1"`
		K []string          `valid:"required#1"`
		L []*string         `valid:"required#1"`
		M struct{ M1 int }  `valid:"required#1"`
		N *struct{ M1 int } `valid:"required#1"`
	}
	var reqStr = `{"IA":true,"A":1025,"B":-127,"C":1,"D":1,"E":1,"F":"string1","G":"string2","H":1,"I":[],"J":[],"K":[],"L":[],"M":{"M1":1},"N":{}}`
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
	fmt.Println("success")
}
