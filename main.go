package main

import (
	"fmt"
	"math"
	"strings"
)

type Info struct {
	Age  int
	Name string
}

func main() {
	info := Info{
		Age:  56,
		Name: "VLAD VALACAS",
	}

	userInfo := float64(info.Age) / 10
	incUSer := math.Ceil(userInfo) * 10
	info.Age = int(math.Ceil(float64(incUSer)))
	info.Name = strings.Title(info.Name)

	fmt.Println(info)
}
