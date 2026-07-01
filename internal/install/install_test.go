package install

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInstall_Claude(t *testing.T) {
	dir := t.TempDir()
	opts := Options{
		Target: TargetClaude,
		Root:   dir,
		Skill:  []byte("hello"),
	}
	written, err := opts.Install()
	if err != nil {
		t.Fatalf("install: %v", err)
	}
	if len(written) != 1 {
		t.Fatalf("want 1 path, got %d", len(written))
	}
	want := filepath.Join(dir, ".claude", "skills", "ponytail-python", "SKILL.md")
	if written[0] != want {
		t.Errorf("path = %s, want %s", written[0], want)
	}
	data, err := os.ReadFile(want)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(data) != "hello" {
		t.Errorf("content mismatch")
	}
}

func TestInstall_Both(t *testing.T) {
	dir := t.TempDir()
	opts := Options{Target: TargetBoth, Root: dir, Skill: []byte("x")}
	written, err := opts.Install()
	if err != nil {
		t.Fatalf("install: %v", err)
	}
	if len(written) != 2 {
		t.Fatalf("want 2 paths, got %d", len(written))
	}
	hasClaude, hasCursor := false, false
	for _, p := range written {
		if strings.Contains(p, ".claude/skills/ponytail-python/SKILL.md") {
			hasClaude = true
		}
		if strings.Contains(p, ".cursor/rules/ponytail-python.mdc") {
			hasCursor = true
		}
	}
	if !hasClaude || !hasCursor {
		t.Errorf("missing target: claude=%v cursor=%v", hasClaude, hasCursor)
	}
}

func TestInstall_EmptySkillErrors(t *testing.T) {
	_, err := Options{Target: TargetClaude, Root: t.TempDir()}.Install()
	if err == nil {
		t.Fatal("want error on empty skill")
	}
}

func TestInstall_CursorGlobalErrors(t *testing.T) {
	_, err := Options{Target: TargetCursor, Global: true, Skill: []byte("x")}.Install()
	if err == nil {
		t.Fatal("want error: cursor + global is not supported")
	}
}
