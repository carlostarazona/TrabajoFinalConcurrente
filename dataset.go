package main

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://datos.ins.gob.pe/dataset/f39f21ee-c6e3-4b24-8bac-512839e07d74/resource/4197e197-7a32-4024-9d8d-ee53b3a4b410/download/ninos_madre_de_dios_estado_nutricional_2017.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	_,err = os.Stdout.Write(body)

	if err != nil {
		log.Fatal(err)
	}
}