package cal

import (
	"testing"
)

func TestUpadder(t *testing.T) {
	res := addUpper(10)
	if res != 45 {
		t.Fatalf("addUpper(10) error")
	}
	t.Logf("success")
}
