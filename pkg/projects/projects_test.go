package projects

import "testing"

func TestNew(t *testing.T) {
	svc := New()
	if svc == nil {
		t.Fatal("expected non-nil service")
	}
}
