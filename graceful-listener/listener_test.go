package listener

import "testing"

func TestListen(t *testing.T) {
	_, err := Listen()
	if err == nil {
		t.Errorf("listen without passed socket: %v", err)
	}
}
