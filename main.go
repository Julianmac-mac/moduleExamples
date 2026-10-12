package main

import (
	"fmt"
	"strings"

	"github.com/Julianmac-mac/moduleExamples/slices"
	"rsc.io/quote/v3"
)

//La funcionalidad de trabajar con paquetes es exportar funcionalidades

func main() {
	list := []string{"EDteam", "Golang", "gophers", quote.HelloV3()} //quote.HelloV3 devuelve un saludo de bienvenida

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})

	fmt.Println(quote.Concurrency())
}
