//go:build linux

package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLegacyMOTDWarningIsRemovedWithoutChangingOriginalContent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "motd")
	const original = "Welcome to the server.\n"
	managed := original + motdWarningStart + "\nlegacy warning\n\n" + motdWarningEnd
	if err := os.WriteFile(path, []byte(managed), 0640); err != nil {
		t.Fatal(err)
	}

	if err := removeInstalledMOTDWarning(path); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != original {
		t.Fatalf("MOTD cleanup changed original content: %q", data)
	}
	info, err := os.Stat(path)
	if err != nil || info.Mode().Perm() != 0640 {
		t.Fatalf("MOTD permissions changed: %v", err)
	}
}
