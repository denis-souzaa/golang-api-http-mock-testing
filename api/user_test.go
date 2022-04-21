package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func (m *MockHttp) Do(_ *http.Request) (*http.Response, error) {
	return m.response, m.err
}

type MockHttp struct {
	response *http.Response
	err      error
}

func TestGetRandomUser(t *testing.T) {
	mock := MockHttp{
		response: &http.Response{
			Body: ioutil.NopCloser(strings.NewReader("")),
		},
	}
	service := Service{
		HttpClient: &mock,
	}

	r := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()

	service.GetRandomUser(w, r)

	got := w.Result()

	if !reflect.DeepEqual(http.StatusOK, got.StatusCode) {
		t.Errorf("wanted: %d, got: %d", http.StatusOK, got.StatusCode)
	}
}
