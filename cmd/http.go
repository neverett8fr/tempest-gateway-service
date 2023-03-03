package cmd

import (
	"fmt"
	"log"
	"net/http"
	"tempest-gateway-service/pkg/config"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartServer(conf *config.Service, router *mux.Router) error {
	// srv := &http.Server{
	// 	Handler:      router,
	// 	Addr:         fmt.Sprintf("%v:%v", conf.Host, conf.Port),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Printf("Server started on port: %v", conf.Port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		c.Handler(router),
	))
	log.Printf("Server started on port: %v", conf.Port)

	// log.Fatal(srv.ListenAndServe())
	return nil
}
