package main

import (
	"strings"

	"github.com/Julianmac-mac/moduleExamples/slices"
	"rsc.io/quote"
)

//La funcionalidad de trabajar con paquetes es exportar funcionalidades

func main() {
	list := []string{"EDteam", "Golang", "gophers", quote.Hello()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "g")
	})
}
