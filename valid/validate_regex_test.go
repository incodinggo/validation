package valid_test

import (
	"encoding/json"
	"fmt"
	"github.com/incodinggo/validation/valid"
	"testing"
)

func TestRegex(t *testing.T) {
	type Inc struct {
		IA bool `valid:""`
	}

	var req struct {
		Inc
		A int               `valid:"regex[10]#3"`
		B int8              `valid:"regex[10]#3"`
		C int16             `valid:"regex[10]#3"`
		D int32             `valid:"regex[10]#3"`
		E int64             `valid:"regex[10]#3"`
		F string            `valid:"regex[^[1-2]\d*$]#3"`
		G *string           `valid:"regex[email]#3"`
		H *int              `valid:"regex[pwd]#3"`
		I []int             `valid:"regex[10]#3"`
		J *[]int            `valid:"regex[10]#3"`
		K []string          `valid:"regex[10]#3"`
		L []*string         `valid:"regex[10]#3"`
		M struct{ M1 int }  `valid:"regex[10]#3"`
		N *struct{ M1 int } `valid:"regex[10]#3"`
	}
	var reqStr = `{"IA":true,"A":10,"B":10,"C":10,"D":10,"E":10,"F":"string1_1A","G":"string1_1A","H":10,"I":[],"J":[],"K":[],"L":[],"M":{"M1":1},"N":{}}`
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
