package json

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	test := `{
				"id": 1,
				"name": "qwerty",
				"birthday": "12.09.2018"
			}`

	if _, err := newConf([]byte(test)); err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	test = `{
				"persons": [
					{
						"id": 1,
						"name": "qwerty",
						"birthday": "12.09.2018"
					},
					{
						"id": 2,
						"name": "asdfgh",
						"birthday": "12.12.2010"
					}
				]
			}`

	if _, err := newConf([]byte(test)); err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	test = `[
				{
					"id": 1,
					"name": "qwerty",
					"birthday": "12.09.2018"
				},
				{
					"id": 2,
					"name": "asdfgh",
					"birthday": "12.12.2010"
				}
			]`

	if _, err := newConf([]byte(test)); err == nil {
		t.Errorf("Got err parsing %v", test)
	}

	test = `{
				"id": 1,
				"name": "qwerty",
				"birthday": "12.09.2018",
				"address": 
				{
					"city": "Moscow",
					"street": "Lenina str."
				}
			}`

	if _, err := newConf([]byte(test)); err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	test = `{
				qwe: "qwe"
			}`

	if _, err := newConf([]byte(test)); err == nil {
		t.Errorf("Got err parsing %v", test)
	}
}

func TestSection(t *testing.T) {
	test := `{
		"id": 1,
		"name": "qwerty",
		"birthday": "12.09.2018",
		"address": 
		{
			"city": "Moscow",
			"street": "Lenina str."
		}
	}`

	config, err := newConf([]byte(test))
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	_, err = config.Section("address")
	if err != nil {
		t.Errorf("Got err parsing section from %v", test)
	}
}

func TestGetString(t *testing.T) {
	test := `{
		"id": 1,
		"name": "qwerty",
		"birthday": "12.09.2018",
		"address": 
		{
			"city": "Moscow",
			"street": "Lenina str."
		}
	}`

	config, err := newConf([]byte(test))
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	// ignoring error - definitely this type
	Config := config.(Config)

	name, err := Config.getString("name")
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	if name != "qwerty" {
		t.Errorf("Got %v, expected qwerty", name)
	}
}

func TestGetInt(t *testing.T) {
	test := `{
		"id": 1,
		"name": "qwerty",
		"birthday": "12.09.2018",
		"address": 
		{
			"city": "Moscow",
			"street": "Lenina str."
		}
	}`

	config, err := newConf([]byte(test))
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	// ignoring error - definitely this type
	Config := config.(Config)

	id, err := Config.getInt("id")
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	if id != 1 {
		t.Errorf("Got %v, expected 1", id)
	}
}

func TestGetTime(t *testing.T) {
	test := `{
		"id": 1,
		"name": "qwerty",
		"birthday": "12.09.2018",
		"address": 
		{
			"city": "Moscow",
			"street": "Lenina str."
		}
	}`

	config, err := newConf([]byte(test))
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	// ignoring error - definitely this type
	Config := config.(Config)

	birthday, err := Config.getTime("birthday")
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	if birthday.Day() != 12 || birthday.Month() != 9 || birthday.Year() != 2018 {
		t.Errorf("Got %v, expected 12.09.2018", birthday)
	}
}

func TestGetDuration(t *testing.T) {
	test := `{
		"id": 1,
		"name": "qwerty",
		"birthday": "12.09.2018",
		"to_pay_in": "30m"
	}`

	config, err := newConf([]byte(test))
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	// ignoring error - definitely this type
	Config := config.(Config)

	want := 30 * time.Minute

	toPayIn, err := Config.getDuration("to_pay_in")
	if err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	if toPayIn != want {
		t.Errorf("Got %v, expected %v", toPayIn, want)
	}
}

func TestGetStringSlice(t *testing.T) {
	test := `{
		"id": 1,
		"name": "qwerty",
		"birthday": "12.09.2018",
		"to_pay_in": "30m",
		"nicknames": [
			"The Most Brilliant",
			"Mr Awesome",
			"Strange Guy"
		]
	}`

	config, err := newConf([]byte(test))
	if err != nil {
		t.Errorf("Got err parsing %v", err)
	}

	// ignoring error - definitely this type
	Config := config.(Config)

	want := []string{"The Most Brilliant", "Mr Awesome", "Strange Guy"}

	nicknames, err := Config.getStringSlice("nicknames")
	if err != nil {
		t.Errorf("Got err parsing %v", err)
	}

	if len(nicknames) != len(want) || nicknames[0] != want[0] ||
		nicknames[1] != want[1] || nicknames[2] != want[2] {
		t.Errorf("Got %v, want %v", nicknames, want)
	}
}
