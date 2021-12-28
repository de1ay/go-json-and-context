package cancelfunc

import (
	"context"
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

func Slide52() {
	log.Print("serving \"slow\" server on http://localhost:9000")
	go startServer()

	bCtx := context.Background()
	cancelContext, cancel := context.WithCancel(bCtx)
	defer cancel()

	req, err := http.NewRequestWithContext(
		cancelContext,
		http.MethodGet,
		"http://localhost:9000",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	_, err = http.DefaultClient.Do(req)
	log.Println(err)
}
