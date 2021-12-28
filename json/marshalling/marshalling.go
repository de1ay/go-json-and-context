package marshalling

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int

	tag string
}

func Slide38() {
	p := Person{
		"Nikita",
		22,
		"some_tag",
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(data))
}
