package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSubdir(t *testing.T) {
	t.Run("wildleap", func(t *testing.T) {
		subdir := GetSubdir("wildleap.dynamicmultimediaga.com")
		assert.Equal(t, "wildleap", subdir)
	})
	t.Run("main", func(t *testing.T) {
		subdir := GetSubdir("dynamicmultimediaga.com")
		assert.Equal(t, "main", subdir)
	})
}
