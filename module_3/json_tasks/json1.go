package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	var ss myStruct

	if err := json.Unmarshal(data, &ss); err != nil {
		fmt.Println(err)
		return
	}
	data = bytes.Trim(data, "{") // испортим json удалив '{'

	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}

	fmt.Printf("%v", ss) // {John Connor 35 true [15 11 37]}

}
