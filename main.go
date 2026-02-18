package main

import (
	"api-fizzbuzz/api"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func getKeysFromFile(path string) (map[string]struct{}, error) {
	fichier, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fichier.Close()
	res := map[string]struct{}{}
	scan := bufio.NewScanner(fichier)
	for scan.Scan() {
		cle := strings.TrimSpace(scan.Text())
		if cle != "" {
			res[cle] = struct{}{}
		}
	}
	return res, nil
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	flagCles := pflag.StringSliceP("key", "k", []string{}, "Clé d'API")
	flagFichCles := pflag.String("keyfile", "", "Fichier de clés")
	flagPort := pflag.Uint16P("port", "p", 8080, "Port exposé par l'API")

	pflag.Parse()

	cles := map[string]struct{}{}
	if *flagFichCles != "" {
		var err error
		cles, err = getKeysFromFile(*flagFichCles)
		if err != nil {
			log.Fatal("Impossible d'ouvrir le fichier de clés : ", err)
		}
	}
	for _, cle := range *flagCles {
		cles[cle] = struct{}{}
	}

	r := api.SetupAPI(cles)
	r.Run(fmt.Sprintf(":%d", *flagPort))
}
