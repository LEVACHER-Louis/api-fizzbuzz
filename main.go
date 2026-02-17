package main

import "api-fizzbuzz/api"

func main() {
	cles := []string{"testcle"}
	r := api.SetupAPI(cles)
	r.Run(":8080")
}
