package handlers

import (
	"github.com/microservices/data"
	//"encoding/json"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		p.GetProducts(rw, r)
		return
	}

	//handle an update

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func(p *Products) GetProducts(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}