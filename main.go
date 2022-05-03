package main

import (
	"fmt"
	"os"

	"training.go/GenCert/cert"
	"training.go/GenCert/html"
)

func main() {
	c, err := cert.New("Golang programming", "Alexis Prout", "2018-06-21")
	if err != nil {
		fmt.Printf("Une erreur est survenue durant la création du certificat: %v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = html.New("output")

	if err != nil {
		fmt.Printf("Une erreur est survenue durant la géneration du pdf")
		os.Exit(1)
	}

	saver.Save(*c)
}
