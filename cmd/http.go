package cmd

import (
	"fmt"
	"log"
	"net/http"
	"tempest-gateway-service/pkg/config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartServer(conf *config.Service, router *mux.Router) error {
	// srv := &http.Server{
	// 	Handler:      router,
	// 	Addr:         fmt.Sprintf("%v:%v", conf.Host, conf.Port),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Printf("Server started on port: %v", conf.Port)

	methodsAllowed := handlers.AllowedMethods([]string{"*"})
	originsAllowed := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		handlers.CORS(methodsAllowed, originsAllowed)(router),
	))
	log.Printf("Server started on port: %v", conf.Port)

	// log.Fatal(srv.ListenAndServe())
	return nil
}
