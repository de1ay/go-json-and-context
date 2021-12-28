package omitempty

import (
	"encoding/json"
	"log"
)

type Example2 struct {
	StringField string
}

type Example struct {
	StringField string `json:",omitempty"`
	Number      int
	Flag        bool
	Arr         []string
	InnerStruct Example2
}

func Slide36() {
	ex2 := Example2{"String"}
	ex := Example{
		"", // field is empty
		20,
		true,
		[]string{},
		ex2,
	}

	out, err := json.Marshal(ex)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(out))
}
