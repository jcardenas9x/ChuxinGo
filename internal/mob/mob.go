package mob

import (
	"encoding/json"
	"os"
	"fmt"
	"log"
)

type Mob struct {
	Name string
	Type string
	Size string
	Elem string
	Elem_level integer
}

func readMobTable() {
	data, err := os.ReadFile("basicMobTable.json")
	if err != nil {
		log.Fatal("There was an issue opening the basicMobTable file: ", err)
	}

	var payload []Mob

	err = json.Unmarshall([]byte(data), &payload)

	if err != nul {
		log.Fatal("There was an issue parsing the basicMobTable: ", err)
	}

	fmt.PrintF("Mobs: %+v", payload)
}

