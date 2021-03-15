package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T)  {
	t.Run("returns Code Education string", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		Greeting(response, request)

		got := response.Body.String()
		want := "<b>Code.education Rocks!!</b>\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
func TestHello(t *testing.T)  {
	t.Run("returns Hello string", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		Hello(response, request)

		got := response.Body.String()
		want := "Hello!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}