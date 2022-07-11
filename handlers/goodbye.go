package handlers

import (
	"net/http"
	"log"
)


type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye{
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	g.l.Println("Goodbye")
	rw.Write([]byte("Goodbye"))
}