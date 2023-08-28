package main

import (
	"log"
	"paquetes/services"
	"paquetes/api"
)
func main(){

svc:=services.NewCatFactService("https://catfact.ninja/fact")
svc=services.NewLogginService(svc)
apiserver:=api.NewApiService(svc)
log.Fatal(apiserver.Start(":3000"))


}