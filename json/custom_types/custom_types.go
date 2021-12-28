package customtypes

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type unixTimestamp struct {
	time.Time
}

func (ut unixTimestamp) MarshalJSON() ([]byte, error) {
	s := strconv.Itoa(int(ut.Unix()))
	return []byte(s), nil
}

func (ut *unixTimestamp) UnmarshalJSON(dat []byte) error {
	unix, err := strconv.Atoi(string(dat))
	if err != nil {
		return err
	}
	*ut = unixTimestamp{time.Unix(int64(unix), 0)}
	return nil
}

type Person struct {
	Name      string        `json:"name"`
	CreatedAt unixTimestamp `json:"created_at"`
}

func Slide45() {
	jData := `
	{
		"name": "Nikita",
		"created_at": 0
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
