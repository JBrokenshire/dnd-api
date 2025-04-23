package service_tests

import (
	"dnd-api/tests/helpers"
	"os"
	"testing"
)

var (
	ts *helpers.TestServer
)

func TestMain(m *testing.M) {
	ts = helpers.NewTestServer("../../.env")
	// Close the database connection
	defer ts.S.Db.Close()

	// Run the test
	code := m.Run()

	os.Exit(code)
}
