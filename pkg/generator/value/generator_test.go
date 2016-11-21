package value

import (
	"testing"
)

// test implementation satisfies the interface
func TestGeneratorInterface(t *testing.T) {
	t.Parallel()
	var _ ValueGenerator = (*ConstantValueGenerator)(nil)
}
