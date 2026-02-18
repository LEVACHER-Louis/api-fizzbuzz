package main

import (
	"api-fizzbuzz/api"
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	flagCles := pflag.StringSliceP("key", "k", []string{}, "Clé d'API")
	flagPort := pflag.Uint16P("port", "p", 8080, "Port exposé par l'API")

	pflag.Parse()

	cles := map[string]struct{}{}
	for _, cle := range *flagCles {
		cles[cle] = struct{}{}
	}

	r := api.SetupAPI(cles)
	r.Run(fmt.Sprintf(":%d", *flagPort))
}
