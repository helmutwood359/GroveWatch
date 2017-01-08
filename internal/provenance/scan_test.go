package provenance

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

// writeTree creates a small deterministic workspace for testing and returns its
// root. It includes a nested directory and an ignored directory.
func writeTree(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
