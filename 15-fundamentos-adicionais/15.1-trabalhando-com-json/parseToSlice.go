package main

import (
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
	cachorroEmJson := `{"nome":"bob","raca":"vira lata","idade":18}`

	var c cachorro

	//o Unmarshal aceia apenas slices de byte
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

}
