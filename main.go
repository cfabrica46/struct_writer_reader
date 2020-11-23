package main

import (
	"fmt"
	"io"
	"log"

	"github.com/cfabrica46/writers_and_readers/struct_writer_reader/utilidades"
)

func main() {

	contenidoorigen := []byte("Hola a todos, soy César UwU")
	origen := utilidades.Nuevolectorescritor(contenidoorigen)

	buf := make([]byte, 4)

	cotenidodestino := []byte{}
	destino := utilidades.Nuevolectorescritor(cotenidodestino)

	for {
		n1, err := origen.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		destino.Write(buf[:n1])

	}
	fmt.Println(string(destino.String()))
}
