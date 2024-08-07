package structural

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ResponseSimulator struct {
	Data string `json:"data"`
}

type HttpRequestSimulator struct{}

func (h *HttpRequestSimulator) get(url string) (string, error) {
	data := fmt.Sprintf("Simulated response data from %v", url)

	res := ResponseSimulator{
		Data: data,
	}

	resp, err := json.Marshal(res)
	if err != nil {
		return "", err
	}

	fmt.Printf("simulated fetch from %v\n", url)
	return string(resp), nil
}

type ResponseParserSimulator struct{}

func (r *ResponseParserSimulator) parse(response string) (res ResponseSimulator, err error) {
	res = ResponseSimulator{}

	fmt.Printf("Parsing response...\n")

	err = json.Unmarshal([]byte(response), &res)

	if err != nil {
		return
	}
	return

}

type ErrorHandlerSimulator struct{}

func (e *ErrorHandlerSimulator) handle(err error) error {
	return errors.New(fmt.Sprintf("Simulated error handling: %v\n", err))
}

type ApiFacadeSimulator struct {
	httpRequest    HttpRequestSimulator
	responseParser ResponseParserSimulator
	errorHandler   ErrorHandlerSimulator
}

func NewApiFacadeSimulator() ApiFacadeSimulator {
	return ApiFacadeSimulator{
		httpRequest:    HttpRequestSimulator{},
		responseParser: ResponseParserSimulator{},
		errorHandler:   ErrorHandlerSimulator{},
	}
}

func (a *ApiFacadeSimulator) fetchUserData(url string) (res ResponseSimulator, err error) {
	resp, err := a.httpRequest.get(url)

	if err != nil {
		err = a.errorHandler.handle(err)
		return
	}

	res, err = a.responseParser.parse(resp)

	if err != nil {
		err = a.errorHandler.handle(err)
		return
	}

	return
}

func clientCodeSimulator() {
	apiFacade := NewApiFacadeSimulator()
	url := "https://api.example.com/simulated-users/1"
	userData, err := apiFacade.fetchUserData(url)

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(userData)
}
