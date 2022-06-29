package valid_test

import (
	"encoding/json"
	"fmt"
	"github.com/incodinggo/validation/valid"
	"testing"
)

func TestArr(t *testing.T) {
	type Inc struct {
		IA bool `valid:""`
	}

	var req struct {
		Inc
		A int               `valid:"arr[1 3]#5"`
		B int8              `valid:"arr[1 3]#5"`
		C int16             `valid:"arr[1 3]#5"`
		D int32             `valid:"arr[1 3]#5"`
		E int64             `valid:"arr[1 3]#5"`
		F string            `valid:"arr[1 3]#5"`
		G *string           `valid:"arr[1 3]#5"`
		H *int              `valid:"arr[1 3]#5"`
		I []int             `valid:"arr[1 3]#5"`
		J *[]int            `valid:"arr[1 3]#5"`
		K []string          `valid:"arr[1 3]#5"`
		L []*string         `valid:"arr[1 3]#5"`
		M struct{ M1 int }  `valid:"arr[1 3]#5"`
		N *struct{ M1 int } `valid:"arr[1 3]#5"`
		o *struct{ M1 int } `valid:"arr[1 3]#5"`
	}
	var reqStr = `{"IA":true,"A":20,"B":10,"C":30,"D":10,"E":20,"F":"A","G":"B","H":20,"I":[1],"J":[1,2,3],"K":["a","b","c"],"L":["A","B","C"],"M":{"M1":1},"N":{}}`
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
