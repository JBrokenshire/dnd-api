package helpers

import (
	"fmt"
	"github.com/JBrokenshire/dnd-api/pkg/try"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"testing"
)

func ClearTable(db *gorm.DB, tableName string) {
	err := db.Exec(fmt.Sprintf("DELETE FROM %v", tableName)).Error
	if err != nil {
		log.Fatalf("You can't clear that table. Err: %v", err)
	}
	err = db.Exec(fmt.Sprintf("ALTER TABLE %v AUTO_INCREMENT = 1", tableName)).Error
	if err != nil {
		log.Fatalf("Error setting autoincrement. Err: %v", err)
	}
}

func ExecuteTestCase(t *testing.T, e *echo.Echo, testCase *try.TestCase) {
	// Perform any setup needed for the test case
	if testCase.Setup != nil {
		testCase.Setup(testCase)
	}
	req, err := try.GenerateRequest(testCase)
	if err != nil {
		t.Fatalf("Unable to generate request")
	}
	res := try.ExecuteRequest(e, req)
	// Perform any teardown needed for test case
	if testCase.Teardown != nil {
		testCase.Teardown(testCase, res)
	}
}
