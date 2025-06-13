/*
	Unit tests for cron package
	Check that cron expressions can only be constructed with valid values if using the CronFromString function
*/
package cron
import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"


// TestCronRegex tests the cron regex patterns
func TestCronRegex(t *testing.T) {
	testCase := "0-59 0-23 1-31 1-12 0-6"
	
	cron := CronFromString(testCase)
	fmt.Println(cron)
	assert.NotNil(t, cron, "Cron should not be nil")
}