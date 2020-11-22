package main

import (
	"fmt"
	"io"
	"log"
)

type lectorescritor struct {
	contenido []byte
	indice    int
	tamaño    int
}

func nuevolectorescritor(b []byte) *lectorescritor {

	return &lectorescritor{contenido: b}
}

func main() {

	contenidoorigen := []byte("Hola a todos, soy César UwU")
	origen := nuevolectorescritor(contenidoorigen)

	buf := make([]byte, 4)

	cotenidodestino := []byte{}
	destino := nuevolectorescritor(cotenidodestino)

	for {
		n1, err := origen.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		n2, _ := destino.Write(buf[:n1])

		fmt.Print(string(destino.contenido[:n2]))
	}
}

func (le *lectorescritor) Read(b []byte) (n int, err error) {

	if le.indice >= len(le.contenido) {

		n = 0
		err = io.EOF
		return

	}

	for i := 0; i < len(b); i++ {

		if le.indice == len(le.contenido) {
			break
		}

		b[i] = le.contenido[le.indice]
		le.indice++
		n++

	}

	return

}

func (le *lectorescritor) Write(b []byte) (n int, err error) {

	le.tamaño = len(le.contenido)

	n1 := le.tamaño

	le.contenido = append(le.contenido, b...)

	le.tamaño = len(le.contenido)

	n2 := le.contenido

	n = n2 - n1

	if n != len(b) {
		log.Fatal(err)
	}

	return
}
