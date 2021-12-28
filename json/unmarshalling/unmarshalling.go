package unmarshalling

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int

	tag string
}

func Slide40() {
	jData := `
	{
		"Name": "Nikita",
		"Age": 22,
		"Tag": "some_tag",
		"ExcessiveField": "some_info"
	}
	`

	p := Person{}

	err := json.Unmarshal([]byte(jData), &p)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(p) // will print {Nikita 22 }
}
