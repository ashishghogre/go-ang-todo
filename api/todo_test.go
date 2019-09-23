package todo

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
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

func createItem(w http.ResponseWriter, r *http.Request) {

}

func FeatureContext(s *godog.Suite) {
	s.Step(`^sending "([^"]*)" to "([^"]*)" with body:$`, sendingToWithBody)
	s.Step(`^should get response status code of "([^"]*)"$`, shouldGetResponseStatusCodeOf)

	s.BeforeScenario(func(interface{}) {
		go func() {
			http.HandleFunc("/item", createItem)
			log.Fatal(http.ListenAndServe(":8083", nil))
		}()
	})
}
func TestMain(m *testing.M) {
	flag.Parse()

	status := godog.RunWithOptions("godogs", FeatureContext, godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "progress", // can define default values
		Paths:  flag.Args(),
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
