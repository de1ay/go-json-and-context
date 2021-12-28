package channel

import (
	"errors"
	"log"
	"net/http"
	"time"
)

func proceedRequest(r *http.Request) ([]byte, int, error) {
	time.Sleep(5 * time.Minute)
	return []byte("{\"status\":\"OK\"}"), http.StatusOK, nil
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, statusCode, err := proceedRequest(r)
	if err != nil {
		log.Print(err)
	}
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func startServer() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":9000", nil)
}

func performRequest(ch chan error) {
	_, err := http.Get("http://localhost:9000")
	ch <- err
}

func Slide49() {
	log.Print("serving \"slow\" server on http://localhost:9000")
	go startServer()

	ch := make(chan error)
	go performRequest(ch)
	go func(ch chan error) {
		time.Sleep(5 * time.Second)
		ch <- errors.New("request timed out")
	}(ch)

	err := <-ch
	log.Println(err)
}
