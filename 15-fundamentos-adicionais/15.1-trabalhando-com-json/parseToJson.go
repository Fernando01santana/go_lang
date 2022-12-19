package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `idade:"idade"`
}

func main() {
	doguinho := cachorro{"bob", "viralata", 10}
	jsonTransform(doguinho)
}

// map to json or struct to json
func jsonTransform(dog cachorro) {
	cachorroEmJson, erro := json.Marshal(dog)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorroEmJson))
}
