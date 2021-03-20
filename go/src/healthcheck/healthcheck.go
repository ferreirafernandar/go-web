package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	_, err := http.Get(fmt.Sprintf("http://localhost:8000/healthcheck"))
	if err != nil {
		os.Exit(1)
	}
}