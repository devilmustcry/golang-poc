package controller_test

import (
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/wiremock/go-wiremock"
	"golang-poc/wiremock/first-party/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SetPortContainerCustomizer struct {
}

func (SetPortContainerCustomizer) Customize(req *testcontainers.GenericContainerRequest) {
	req.ExposedPorts = []string{"8081/tcp"}
}

func TestHandleHello(t *testing.T) {
	//ctx := context.Background()
	wc := wiremock.NewClient("http://localhost:9090")
	err := wc.StubFor(
		wiremock.Get(wiremock.URLEqualTo("/hello")).
			WillReturnResponse(wiremock.NewResponse().
				WithJSONBody(map[string]string{"result": "Hello, world!"}).
				WithHeader("Content-Type", "application/json").
				WithStatus(http.StatusOK),
			),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("It should return what that is mocked within wiremock", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/hello", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controller.HandleHello)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
		fmt.Println(resp.Body.String())
	})
	wc.Reset()

	stubRule := wiremock.Get(wiremock.URLEqualTo("/hello")).
		WillReturnResponse(wiremock.NewResponse().
			WithJSONBody(map[string]string{"result": "Hello, world2"}).
			WithHeader("Content-Type", "application/json").
			WithStatus(http.StatusOK),
		)
	err = wc.StubFor(
		stubRule,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("It should return what that is mocked within wiremock", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/hello", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(controller.HandleHello)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
		fmt.Println(resp.Body.String())
	})
	wc.DeleteStubByID(stubRule.UUID())
	//statusCode, out, err := SendHttpGet(container, "/hello", nil)
	//if err != nil {
	//	t.Fatal(err, "Failed to get a response")
	//}
	//
	//if statusCode != 200 {
	//	t.Fatalf("expected HTTP-200 but got %d", statusCode)
	//}
	//
	//if string(out) != `{"result":"Hello, world!"}` {
	//	t.Fatalf("expected 'Hello, world!' but got %v", out)
	//}
}
