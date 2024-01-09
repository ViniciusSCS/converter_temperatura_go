package main

import (
	"converter-graus/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.ConverterGraus()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
