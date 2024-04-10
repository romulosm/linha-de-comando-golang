package main

import (
	"fmt"
	"log"
	"modulo/app"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}

}
