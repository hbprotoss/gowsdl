package main

import (
	"fmt"
	"gowsdl/enterprise"
)

func main() {
	var service = enterprise.NewHelloService()
	response, err := service.Hello("haha")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
