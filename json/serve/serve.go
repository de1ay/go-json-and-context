package serve

import (
	"log"
	"net/http"
)

func proceedRequest(r *http.Request) ([]byte, int, error) {
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

func Slide39() {
	log.Print("serving the example for slide 39, you can open http://localhost:9000")
	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
