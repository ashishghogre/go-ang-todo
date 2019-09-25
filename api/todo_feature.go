package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

var response *http.Response

func getURL(path string) string {
	return "http://localhost:8083" + path
}

func sendingToWithBody(method, path string, body *gherkin.DocString) (err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Trying to send %s %s with body: \n%s\n", method, path, body.Content)
	response, err = http.Post(getURL(path), "application/json", strings.NewReader(body.Content))
	return
}

func shouldGetResponseStatusCodeOf(expectedStatusCode int) error {
	if expectedStatusCode != response.StatusCode {
		return fmt.Errorf("Expected status code %d, recieved %d", expectedStatusCode, response.StatusCode)
	}
	return nil
}

func sendingToShouldReturnItemsContaining(method, path string, expectedText *gherkin.DocString) (err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Trying to send %s %s", method, path)
	response, err = http.Get(getURL(path))
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(response.Body)
	if !strings.Contains(string(body), expectedText.Content) {
		return fmt.Errorf("expected text %s, but got %s", expectedText.Content, string(body))
	}
	return nil
}

func stepDefinitions(s *godog.Suite) {
	s.Step(`^sending "([^"]*)" to "([^"]*)" with body:$`, sendingToWithBody)
	s.Step(`^should get response status code of "([^"]*)"$`, shouldGetResponseStatusCodeOf)
	s.Step(`^sending "([^"]*)" to "([^"]*)" should return items containing:$`, sendingToShouldReturnItemsContaining)
}
