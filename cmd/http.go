package cmd

import (
	"fmt"
	"log"
	"net/http"
	"tempest-gateway-service/pkg/config"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartServer(conf *config.Service, router *mux.Router) error {

	// methodsAllowed := handlers.AllowedMethods([]string{"*"})
	// originsAllowed := handlers.AllowedOrigins([]string{"*"})
	// headersAllowed := handlers.AllowedHeaders([]string{"*"})

	// srv := &http.Server{
	// 	Handler:      router,
	// 	Addr:         fmt.Sprintf("%v:%v", conf.Host, conf.Port),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Printf("Server started on port: %v", conf.Port)

	handler := cors.Default().Handler(router)

	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// log.Fatal(http.ListenAndServe(
	// 	fmt.Sprintf("%v:%v", conf.Host, conf.Port),
	// 	handlers.CORS(
	// 		methodsAllowed,
	// 		originsAllowed,
	// 		headersAllowed,
	// 	)(router),
	// ))
	log.Fatal(srv.ListenAndServe())
	log.Printf("Server started on port: %v", conf.Port)

	// log.Fatal(srv.ListenAndServe())
	return nil
}
