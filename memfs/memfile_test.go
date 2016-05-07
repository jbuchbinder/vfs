package memfs

import (
	"github.com/jbuchbinder/vfs"
	"testing"
)

func TestFileInterface(t *testing.T) {
	_ = vfs.File(NewMemFile("", nil, nil))
}
