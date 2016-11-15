package serialize

import (
	"testing"
)
// test implementation satisfies the interface
func TestGeneratorInterface(t *testing.T) {
	t.Parallel()
	var _ Serializer = (*DebugSerializer)(nil)
}
