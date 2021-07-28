package internal

import "testing"

func TestImportantNumberIs100(t *testing.T) {
	if ImportantNumber != 100 {
		t.Errorf("ImportantNumber was incorrect, got: %d wanted: 100", ImportantNumber)
	}
}
