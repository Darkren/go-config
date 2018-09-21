package json

import (
	"testing"
)

func TestNew(t *testing.T) {
	test := `{
				"id": 1,
				"name": "qwerty",
				"birthday": "12.09.2018"
			}`

	if _, err := new([]byte(test)); err != nil {
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

	if _, err := new([]byte(test)); err != nil {
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

	if _, err := new([]byte(test)); err == nil {
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

	if _, err := new([]byte(test)); err != nil {
		t.Errorf("Got err parsing %v", test)
	}

	test = `{
				qwe: "qwe"
			}`

	if _, err := new([]byte(test)); err == nil {
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

	config, err := new([]byte(test))
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

	config, err := new([]byte(test))
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

	config, err := new([]byte(test))
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

	config, err := new([]byte(test))
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
