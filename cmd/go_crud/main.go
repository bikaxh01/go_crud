package main

import (
	"fmt"
	"net/http"

	"github.com/bikaxh01/go_crud/internal/config"
)

func main() {

	ctx := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("OK"))
	})

	server := http.Server{
		Addr:    ctx.Address,
		Handler: router,
	}
	fmt.Println("Running at ........",ctx.Address)
	server.ListenAndServe()

}
