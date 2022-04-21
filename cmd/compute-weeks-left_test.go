package cmd

import (
	"testing"
)

func TestComputeWeeksLeftBeforeNextYear(t *testing.T) {
	weeksLeft, _ := ComputeWeeksLeftBeforeNextYear()
	if weeksLeft != "~ 46 weeks left this year" {
		t.Errorf("\nexpected: [~ 46 weeks left this year]\nreceived: [%s]", weeksLeft)
	}
}
