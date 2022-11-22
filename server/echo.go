package server

import (
	"github.com/TelephoneTan/GoHTTPGzipServer/gzip"
	"io"
	"log"
	"net/http"
)

func Echo() {
	h := http.NewServeMux()
	h.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		all, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = writer.Write(all)
	})
	gzipH := (&gzip.Handler{Handler: h}).Init()
	log.Fatal(http.ListenAndServe("localhost:3000", gzipH))
}
