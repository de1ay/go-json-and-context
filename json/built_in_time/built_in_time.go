package builtintime

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func Slide46() {
	jData := `
	{
		"name": "Nikita",
		"created_at": "1970-01-01T00:00:00Z"
	}
	`

	p := Person{}

	err := json.Unmarshal([]byte(jData), &p)
	if err != nil {
		fmt.Print(err)
	}
	// will print "1970-01-01 00:00:00 +0000 UTC"
	fmt.Println(p.CreatedAt.UTC())
}
