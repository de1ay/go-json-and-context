package timeout

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

func Slide51() {
	log.Print("serving \"slow\" server on http://localhost:9000")
	go startServer()

	bCtx := context.Background()
	timeoutContext, cancel := context.WithTimeout(bCtx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		timeoutContext,
		http.MethodGet,
		"http://localhost:9000",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = http.DefaultClient.Do(req)
	log.Println(err)
}
