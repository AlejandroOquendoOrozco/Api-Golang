package api

import (
	"context"
	"encoding/json"
	"net/http"
	"paquetes/services"
)
type apiServer struct{

  svc services.Service

}

func NewApiService(svc services.Service)*apiServer{

	return &apiServer{

		svc: svc,
	}
}
func(api *apiServer)Start(listenAdd string )error{

	http.HandleFunc("/",api.handleGetCatFact)
	return http.ListenAndServe(listenAdd,nil)
}

func WriteJSON(w http.ResponseWriter,s int,v any)error{

	w.WriteHeader(s)
	w.Header().Add("Content-Type","application/json")
	return json.NewEncoder(w).Encode(v)
}

func (api *apiServer)handleGetCatFact(w http.ResponseWriter, r *http.Request){

	fact,err:=api.svc.GetCatFact(context.Background())

	if err!=nil{

       WriteJSON(w,http.StatusUnprocessableEntity,map[string]any{"error":err.Error()})


	}
	WriteJSON(w,http.StatusOK,fact)
}

