package serialize

import (
	"testing"
)
// test implementation satisfies the interface
func TestSerializerInterface(t *testing.T) {
	t.Parallel()
	var _ Serializer = (*DebugSerializer)(nil)
}
