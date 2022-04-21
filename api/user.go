package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type Service struct {
	HttpClient HttpClientInterface
}

func (s *Service) GetRandomUser(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://randomuser.me/api"

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := s.HttpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
