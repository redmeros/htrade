package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveConfigPath(t *testing.T) {
	paths := TryResolveConfigPath()
	assert.Greater(t, len(paths), 1, paths)
}
