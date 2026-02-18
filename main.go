package main

import (
	"api-fizzbuzz/api"

	"github.com/spf13/pflag"
)

func main() {
	flagCles := pflag.StringSliceP("key", "k", []string{}, "Clé d'API")

	pflag.Parse()

	cles := map[string]struct{}{}
	for _, cle := range *flagCles {
		cles[cle] = struct{}{}
	}

	r := api.SetupAPI(cles)
	r.Run(":8080")
}
