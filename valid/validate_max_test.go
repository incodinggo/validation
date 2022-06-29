package valid_test

import (
	"encoding/json"
	"fmt"
	"github.com/incodinggo/validation/valid"
	"testing"
)

func TestMin(t *testing.T) {
	type Inc struct {
		IA bool `valid:""`
	}

	var req struct {
		Inc
		A int               `valid:"min[10]#3"`
		B int8              `valid:"min[10]#3"`
		C int16             `valid:"min[10]#3"`
		D int32             `valid:"min[10]#3"`
		E int64             `valid:"min[10]#3"`
		F string            `valid:"min[10]#3"`
		G *string           `valid:"min[10]#3"`
		H *int              `valid:"min[10]#3"`
		I []int             `valid:"min[10]#3"`
		J *[]int            `valid:"min[10]#3"`
		K []string          `valid:"min[10]#3"`
		L []*string         `valid:"min[10]#3"`
		M struct{ M1 int }  `valid:"min[10]#3"`
		N *struct{ M1 int } `valid:"min[10]#3"`
	}
	var reqStr = `{"IA":true,"A":100,"B":100,"C":100,"D":100,"E":100,"F":"string1string1","G":"string2string2","H":11,"I":[],"J":[],"K":[],"L":[],"M":{"M1":1},"N":{}}`
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
