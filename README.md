# ponytail-for-python

**The lazy senior Python developer, in your terminal.**

A single-file Claude Code / Cursor skill that makes AI write Python like it has been maintaining the codebase for five years. No unnecessary deps. No premature abstraction. No ceremonial docstrings on 3-line functions.

<p align="center">
  <img src="assets/hero.gif" alt="Same prompt, left without ponytail (4 layers of abstraction), right with ponytail (12 lines)." width="720">
</p>

<p align="center">
  <a href="https://github.com/Danush-Aries/ponytail-for-python/releases/latest"><img src="https://img.shields.io/github/v/release/Danush-Aries/ponytail-for-python?style=for-the-badge&label=install" alt="latest release"></a>
  <a href="https://github.com/Danush-Aries/ponytail-for-python/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue?style=for-the-badge" alt="MIT"></a>
  <a href="https://github.com/Danush-Aries/ponytail-for-python/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/Danush-Aries/ponytail-for-python/ci.yml?branch=main&style=for-the-badge" alt="CI"></a>
</p>

---

## Install

```bash
curl -sSL https://raw.githubusercontent.com/Danush-Aries/ponytail-for-python/main/install.sh | sh
```

That drops the `pnt` binary in `~/.local/bin` and installs the skill in the current directory.

Prefer to see what runs? [Read the installer](install.sh) — it is 60 lines of POSIX `sh`.

Or grab a binary from [Releases](https://github.com/Danush-Aries/ponytail-for-python/releases/latest) and run:

```bash
pnt install                 # writes .claude/skills/ponytail-python/SKILL.md in cwd
pnt install --target cursor # writes .cursor/rules/ponytail-python.mdc in cwd
pnt install --target both   # both
pnt install --global        # ~/.claude/skills/ponytail-python/SKILL.md
```

Reload your editor. That is the whole thing.

---

## What it does

| | With ponytail | Without |
|---|---|---|
| Small script | 12 lines with `dataclass` and `pathlib` | 4 classes, 3 abstract bases, a `Factory` |
| Error handling | narrow `except`, act on it | `except Exception: pass`, print to stderr |
| Deps | stdlib first, third-party justified | `pip install httpx requests urllib3 aiohttp` |
| Docstrings | on public API only | on every 3-line helper |
| Comments | explain the surprise | narrate what the code already says |

---

## The rulebook

<details>
<summary>Eight rules (click to expand)</summary>

1. **No unnecessary dependencies.** Check stdlib first. `pathlib`, `dataclasses`, `functools`, `sqlite3`, `tomllib`, `zoneinfo`, `http.server` are all there.
2. **No premature abstraction.** No `Manager` / `Factory` / `Service` for a single call site. Extract at three call sites, not before.
3. **No ceremony.** No docstring on 3-line functions. No type hints on `i: int = 0`. No `# TODO: refactor later`.
4. **Idiomatic Python.** List comps over `map+lambda`. `dataclass` over hand-written `__init__`. `pathlib.Path` over string concat. f-strings.
5. **Delete dead code.** Zero callers → remove it. Git remembers.
6. **Trust internal calls.** Validate at the boundary. Do not `isinstance`-guard between your own modules.
7. **Error handling is code, not decoration.** Catch narrowly, act on it. No bare `except`. No `try` around code that cannot throw.
8. **Comments explain _why_, never _what_.** Write a comment only when a future reader would be surprised.

Full skill file: [skill.md](skill.md).

</details>

---

## How it works

`pnt install` writes the skill to a location your editor already looks at:

- **Claude Code** — `.claude/skills/ponytail-python/SKILL.md` (project) or `~/.claude/skills/ponytail-python/SKILL.md` (`--global`). Claude auto-loads it.
- **Cursor** — `.cursor/rules/ponytail-python.mdc` (project). Cursor auto-loads it.

The rules go against the default habit of generating verbose, over-abstracted code. Nothing about ponytail is stateful — the skill is just a Markdown file. Uninstall by deleting it.

---

## Why

Language-model coding assistants default to *ceremony*. Empty docstrings. Layered abstractions. Try/except around lines that cannot throw. `Manager` classes for one-off scripts. This is what the training data rewarded.

It is not what a five-year Python engineer would write.

Ponytail-for-Python is a distilled version of the instincts a senior dev applies without thinking. Loaded once, applied every prompt.

---

## Contributing

Rules that catch real patterns beat rules that sound clever. If you have an example of a pattern the current skill missed, open an issue with:

- The prompt that produced it
- What the model wrote
- What you actually wanted

PRs to `skill.md` are welcome. Do not add third-party Go dependencies to the CLI without a good reason — the point is that this stays under 200 lines of Go and one Markdown file.

---

## Credit

The Python-specific fork of [Ponytail](https://github.com/mrgoonie/ponytail), the viral "lazy senior developer" coding skill.

MIT-licensed. Do whatever.
