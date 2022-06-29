package valid_test

import (
	"encoding/json"
	"fmt"
	"github.com/incodinggo/validation/valid"
	"testing"
)

func TestRange(t *testing.T) {
	type Inc struct {
		IA bool `valid:""`
	}

	var req struct {
		Inc
		A int               `valid:"range[1,100]#2"`
		B int8              `valid:"range[1,100]#2"`
		C int16             `valid:"range[1,100]#2"`
		D int32             `valid:"range[1,100]#2"`
		E int64             `valid:"range[1,100]#2"`
		F string            `valid:"range[1,100]#2"`
		G *string           `valid:"range[1,100]#2"`
		H *int              `valid:"range[1,100]#2"`
		I []int             `valid:"range[1,100]#2"`
		J *[]int            `valid:"range[1,100]#2"`
		K []string          `valid:"range[1,100]#2"`
		L []*string         `valid:"range[1,100]#2"`
		M struct{ M1 int }  `valid:"range[1,100]#2"`
		N *struct{ M1 int } `valid:"range[0,100]#2"`
	}
	var reqStr = `{"IA":true,"A":1,"B":1,"C":1,"D":1,"E":1,"F":"string1","G":"string2","H":1,"I":[],"J":[],"K":[],"L":[],"M":{"M1":1},"N":{}}`
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
