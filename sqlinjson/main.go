package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {

	jq := gojsonq.New().File("./data.json")
	res := jq.From("EmpId")
	fmt.Println(res)	
}