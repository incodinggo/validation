package valid_test

import (
	"encoding/json"
	"fmt"
	"github.com/incodinggo/validation/valid"
	"testing"
)

func TestMax(t *testing.T) {
	type Inc struct {
		IA bool `valid:""`
	}

	var req struct {
		Inc
		A int               `valid:"max[10]#3"`
		B int8              `valid:"max[10]#3"`
		C int16             `valid:"max[10]#3"`
		D int32             `valid:"max[10]#3"`
		E int64             `valid:"max[10]#3"`
		F string            `valid:"max[10]#3"`
		G *string           `valid:"max[10]#3"`
		H *int              `valid:"max[10]#3"`
		I []int             `valid:"max[10]#3"`
		J *[]int            `valid:"max[10]#3"`
		K []string          `valid:"max[10]#3"`
		L []*string         `valid:"max[10]#3"`
		M struct{ M1 int }  `valid:"max[10]#3"`
		N *struct{ M1 int } `valid:"max[10]#3"`
	}
	var reqStr = `{"IA":true,"A":10,"B":10,"C":10,"D":10,"E":10,"F":"string1","G":"string2","H":10,"I":[],"J":[],"K":[],"L":[],"M":{"M1":1},"N":{}}`
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
