package valid_test

import (
	"encoding/json"
	"fmt"
	"github.com/incodinggo/validation/valid"
	"testing"
)

func TestEnum(t *testing.T) {
	type Inc struct {
		IA bool `valid:""`
	}

	var req struct {
		Inc
		A int               `valid:"enum[10,20,30]#4"`
		B int8              `valid:"enum[10,20,30]#4"`
		C int16             `valid:"enum[10,20,30]#4"`
		D int32             `valid:"enum[10,20,30]#4"`
		E int64             `valid:"enum[10,20,30]#4"`
		F string            `valid:"enum[A,B,C]#4"`
		G *string           `valid:"enum[A,B,C]#4"`
		H *int              `valid:"enum[10,20,30]#4"`
		I []int             `valid:"enum[10,20,30]#4"`
		J *[]int            `valid:"enum[10,20,30]#4"`
		K []string          `valid:"enum[10,20,30]#4"`
		L []*string         `valid:"enum[10,20,30]#4"`
		M struct{ M1 int }  `valid:"enum[10,20,30]#4"`
		N *struct{ M1 int } `valid:"enum[10,20,30]#4"`
	}
	var reqStr = `{"IA":true,"A":20,"B":10,"C":30,"D":10,"E":20,"F":"A","G":"B","H":20,"I":[],"J":[],"K":[],"L":[],"M":{"M1":1},"N":{}}`
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
