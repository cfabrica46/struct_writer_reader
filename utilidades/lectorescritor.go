package utilidades

import (
	"io"
	"log"
)

//Lectorescritor es exportado...
type Lectorescritor struct {
	contenido []byte
	indice    int
	tamaño    int
}

//Nuevolectorescritor es exportado...
func Nuevolectorescritor(b []byte) *Lectorescritor {

	return &Lectorescritor{contenido: b}
}

func (le *Lectorescritor) Read(b []byte) (n int, err error) {

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

func (le *Lectorescritor) Write(b []byte) (n int, err error) {

	le.tamaño = len(le.contenido)

	n1 := le.tamaño

	le.contenido = append(le.contenido, b...)

	le.tamaño = len(le.contenido)

	n2 := le.tamaño

	n = n2 - n1

	if n != len(b) {
		log.Fatal(err)
	}

	return
}

func (le *Lectorescritor) String() (contenido string) {

	contenido = string(le.contenido)

	return
}
