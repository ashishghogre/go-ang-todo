package main

import (
	"flag"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
)

func FeatureContext(s *godog.Suite) {
	stepDefinitions(s)
	s.BeforeScenario(func(interface{}) {
		go func() { startServer(8083,"todo_test.db") }()
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
