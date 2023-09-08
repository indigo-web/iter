package iter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExtract(t *testing.T) {
	iter := Slice([]string{"hello", "world"})
	want := []string{"hello", "world"}
	require.Equal(t, want, Extract(iter, nil))
}
