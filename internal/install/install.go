package install

import (
	"fmt"
	"os"
	"path/filepath"
)

type Target string

const (
	TargetClaude Target = "claude"
	TargetCursor Target = "cursor"
	TargetBoth   Target = "both"
)

type Options struct {
	Target Target
	Global bool
	Root   string
	Skill  []byte
}

func (o Options) Install() ([]string, error) {
	if len(o.Skill) == 0 {
		return nil, fmt.Errorf("skill content is empty")
	}
	root := o.Root
	if o.Global {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("resolve home: %w", err)
		}
		root = home
	}
	if root == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("resolve cwd: %w", err)
		}
		root = wd
	}

	var written []string
	targets := targetsFor(o.Target)
	for _, t := range targets {
		path, err := targetPath(root, t, o.Global)
		if err != nil {
			return written, err
		}
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return written, fmt.Errorf("create dir %s: %w", filepath.Dir(path), err)
		}
		if err := os.WriteFile(path, o.Skill, 0o644); err != nil {
			return written, fmt.Errorf("write %s: %w", path, err)
		}
		written = append(written, path)
	}
	return written, nil
}

func targetsFor(t Target) []Target {
	switch t {
	case TargetBoth:
		return []Target{TargetClaude, TargetCursor}
	case TargetCursor:
		return []Target{TargetCursor}
	default:
		return []Target{TargetClaude}
	}
}

func targetPath(root string, t Target, global bool) (string, error) {
	switch t {
	case TargetClaude:
		if global {
			return filepath.Join(root, ".claude", "skills", "ponytail-python", "SKILL.md"), nil
		}
		return filepath.Join(root, ".claude", "skills", "ponytail-python", "SKILL.md"), nil
	case TargetCursor:
		if global {
			return "", fmt.Errorf("cursor rules are per-project; --global only works with --target claude")
		}
		return filepath.Join(root, ".cursor", "rules", "ponytail-python.mdc"), nil
	}
	return "", fmt.Errorf("unknown target %q", t)
}
