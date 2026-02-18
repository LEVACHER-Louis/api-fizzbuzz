package main

import "api-fizzbuzz/api"

func main() {
	cles := map[string]struct{}{"testcle":struct{}{}}
	r := api.SetupAPI(cles)
	r.Run(":8080")
}
