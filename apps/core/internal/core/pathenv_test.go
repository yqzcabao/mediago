package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPathWithDirectoryPrependsDependencyDirectory(t *testing.T) {
	separator := string(os.PathListSeparator)
	depsDir := filepath.Join("/app", "deps")
	current := strings.Join([]string{"/usr/local/bin", "/usr/bin", "/bin"}, separator)

	want := strings.Join([]string{depsDir, "/usr/local/bin", "/usr/bin", "/bin"}, separator)
	if got := PathWithDirectory(current, depsDir); got != want {
		t.Fatalf("PathWithDirectory() = %q, want %q", got, want)
	}
}

func TestPathWithDirectoryKeepsSingleEntryWhenAlreadyPresent(t *testing.T) {
	separator := string(os.PathListSeparator)
	depsDir := filepath.Join("/app", "deps")
	current := strings.Join([]string{"/usr/bin", depsDir, "/bin"}, separator)

	want := strings.Join([]string{depsDir, "/usr/bin", "/bin"}, separator)
	got := PathWithDirectory(current, depsDir)
	if got != want {
		t.Fatalf("PathWithDirectory() = %q, want %q", got, want)
	}
	if again := PathWithDirectory(got, depsDir); again != got {
		t.Fatalf("PathWithDirectory() is not idempotent: %q then %q", got, again)
	}
}

func TestPathWithDirectoryHandlesEmptyValues(t *testing.T) {
	depsDir := filepath.Join("/app", "deps")
	if got := PathWithDirectory("", depsDir); got != depsDir {
		t.Fatalf("PathWithDirectory() = %q, want %q", got, depsDir)
	}
	if got := PathWithDirectory("/usr/bin", "  "); got != "/usr/bin" {
		t.Fatalf("PathWithDirectory() = %q, want the original PATH", got)
	}
	if got := PathWithDirectory("", ""); got != "" {
		t.Fatalf("PathWithDirectory() = %q, want empty", got)
	}
}

func TestPathWithDirectoryDropsEmptyEntries(t *testing.T) {
	separator := string(os.PathListSeparator)
	depsDir := filepath.Join("/app", "deps")

	want := strings.Join([]string{depsDir, "/usr/bin"}, separator)
	if got := PathWithDirectory(separator+"/usr/bin"+separator, depsDir); got != want {
		t.Fatalf("PathWithDirectory() = %q, want %q", got, want)
	}
}
