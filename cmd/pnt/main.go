package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/Danush-Aries/ponytail-for-python/internal/install"
)

//go:embed skill.md
var skill []byte

var version = "dev"

const usage = `pnt — install the ponytail-for-python skill

Usage:
  pnt install [--target claude|cursor|both] [--global]
  pnt version
  pnt help

Flags:
  --target   Which host to install into (default: claude)
  --global   Install for the current user instead of the current repo
             (only valid with --target claude)

Examples:
  pnt install                       # writes .claude/skills/ponytail-python/SKILL.md in cwd
  pnt install --target cursor       # writes .cursor/rules/ponytail-python.mdc in cwd
  pnt install --target both         # writes both
  pnt install --global              # writes to ~/.claude/skills/ponytail-python/SKILL.md
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usage)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "install":
		if err := runInstall(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	case "version", "--version", "-v":
		fmt.Println(version)
	case "help", "--help", "-h":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n%s", os.Args[1], usage)
		os.Exit(1)
	}
}

func runInstall(args []string) error {
	fs := flag.NewFlagSet("install", flag.ExitOnError)
	target := fs.String("target", "claude", "claude | cursor | both")
	global := fs.Bool("global", false, "install for the current user (~/.claude)")
	if err := fs.Parse(args); err != nil {
		return err
	}

	written, err := install.Options{
		Target: install.Target(*target),
		Global: *global,
		Skill:  skill,
	}.Install()
	if err != nil {
		return err
	}
	fmt.Println("ponytail-python installed:")
	for _, p := range written {
		fmt.Println("  ", p)
	}
	fmt.Println("\nreload your editor to pick up the skill.")
	return nil
}
