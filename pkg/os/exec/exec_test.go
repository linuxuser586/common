package exec

import (
	"context"
	"testing"
)

const msg = "fake message"

func TestCommand(t *testing.T) {
	cmd, ctrl := FakeCommand(t)
	defer ctrl.Finish()
	cmd.EXPECT().CombinedOutput().Return([]byte(msg), nil)
	c := Command("fake")
	r, err := c.CombinedOutput()
	if err != nil {
		t.Errorf("got: %v, want: %v", err, msg)
	}
	if string(r) != msg {
		t.Errorf("got: %v, want: %v", r, msg)
	}
}
func TestCommandContext(t *testing.T) {
	cmd, ctrl := FakeCommandContext(t)
	defer ctrl.Finish()
	cmd.EXPECT().CombinedOutput().Return([]byte(msg), nil)
	c := CommandContext(context.Background(), "fake")
	r, err := c.CombinedOutput()
	if err != nil {
		t.Errorf("got: %v, want: %v", err, msg)
	}
	if string(r) != msg {
		t.Errorf("got: %v, want: %v", r, msg)
	}
}
