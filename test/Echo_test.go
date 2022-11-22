package test

import (
	"github.com/TelephoneTan/GoHTTPEcho/server"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	go server.Echo()
	s := "hello, world"
	println("=== Echo ===============")
	println(s)
	res, err := http.Post("http://localhost:3000", "text/plain", strings.NewReader(s))
	if err != nil {
		log.Fatal("request error: ", err)
	}
	all, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal("read error: ", err)
	}
	println("--- Response ---------------")
	println(string(all))
}
