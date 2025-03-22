package fontforge

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVersion_Success(t *testing.T) {
	t.Parallel()

	v, err := newVersion("003.001")
	require.NoError(t, err)
	require.NotNil(t, v)

	assert.Equal(t, "3.1.0", v.String())
}
