package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	srvAddr = ":8080"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {

	router := mux.NewRouter()

	router.MethodNotAllowedHandler = http.HandlerFunc(sendNotFound)
	router.NotFoundHandler = http.HandlerFunc(sendBadMethod)

	router.HandleFunc("/healthz", healthHandler).Methods("GET")
	router.HandleFunc("/api/v1/koalas", koalaHandler).Methods("GET")

	methodsOk := handlers.AllowedMethods([]string{"GET"})

	go func() {
		log.Infof("server listening at %s", srvAddr)
		if err := http.ListenAndServe(srvAddr, handlers.CORS(methodsOk)(router)); err != nil {
			log.Panicf("error while serving: %s", err)
		}
	}()

	// allows clean handling of kill signals to the pid
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
	<-sigC

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK: 200")
}

func koalaHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse("koalas are great!", w, r)
}

func sendResponse(response string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str := `{"Response": "` + response + `"}`
	fmt.Fprint(w, str)
	return
}

func sendNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	str := `{"Error":"Path not allowed"}`
	fmt.Fprint(w, str)
	return
}

func sendBadMethod(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	str := `{"Error":"Method not allowed"}`
	fmt.Fprint(w, str)
	return
}
