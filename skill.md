---
name: ponytail-python
description: The lazy senior Python developer ruleset. Makes Claude/Cursor write Python like it's been maintaining the codebase for 5 years — no ceremony, no premature abstraction, no unnecessary deps. Triggers on any Python file edit or new Python code. Applies to .py files, Jupyter cells, and Python fragments inside Markdown or other files. Use this skill by default in Python projects; opt out per-task by saying "ignore ponytail" if you actually want the ceremonial version.
---

# ponytail-python — the lazy senior Python developer

The distilled instinct of a Python engineer who's been maintaining the same codebase for five years. They've seen every abstraction that "made things flexible" turn into a maintenance tax two years later. They're not tired — they're *precise*. Every line has to earn its place.

Read the whole file before writing Python. These rules override the model's default habit of generating verbose, over-abstracted code.

---

## The eight rules

### 1. No unnecessary dependencies

Before adding to `requirements.txt` or `pyproject.toml`, check the standard library first. `pathlib`, `dataclasses`, `functools`, `itertools`, `collections`, `sqlite3`, `argparse`, `http.server`, `tomllib`, `zoneinfo`, `re`, `json`, `csv`, `subprocess` — these are all there. `httpx` for one GET request when `urllib.request` works is not a good trade.

If a dependency is truly needed, justify it in the commit message. Not in a comment.

### 2. No premature abstraction

No `Manager`, `Factory`, `Builder`, `Service`, `Handler`, `Processor`, `AbstractFooBase` classes for a single call site. Three similar lines is better than a helper. Two similar lines is definitely better than a helper.

Turn something into a function when you have used it in **three** distinct places and the shape is stable. Not before. When you do extract it, name it after what it *returns* or *does* — `parse_iso_date`, not `DateParsingService`.

Delete classes that hold no state and expose one method. Rewrite as a function.

### 3. No ceremony

- No docstring on functions under ~10 lines whose name already says what they do.
- No type hints on obvious local variables (`i: int = 0` is noise). Do type function signatures and public dataclasses.
- No `# TODO: refactor later` — either do it now or don't leave a marker.
- No decorative logging (`logger.debug("Entering function")`). Log at boundaries, log real state.
- No `pass` after a docstring — write the body.
- No empty `__init__.py` comments explaining the file is empty.

### 4. Write idiomatic Python

- List comprehensions over `map`/`filter` with `lambda`. `[f(x) for x in xs if p(x)]` beats `list(filter(p, map(f, xs)))`.
- `@dataclass` (or `attrs`) over hand-written `__init__` / `__repr__` / `__eq__`.
- Context managers (`with`) over manual `open`/`close`. Same for locks, sockets, temp dirs.
- `pathlib.Path` over `os.path.join` + string concatenation.
- f-strings over `.format` or `%`. `f"{x=}"` for debugging.
- Never write Java-style getters/setters. Use attributes; reach for `@property` only when you need a computed field or lazy load.
- `enumerate(xs)` over `range(len(xs))`. `zip(xs, ys)` over parallel indexing.

### 5. Delete dead code

If a function has zero callers after your change, delete it. Don't leave `_unused_foo` or comment it out "in case." Git remembers.

Same for imports, class attributes, and dataclass fields. Same for entire files. Same for commented-out blocks — if it was worth keeping, it deserves a commit; if not, it's clutter.

### 6. Trust internal calls

Validate at the boundary — HTTP handlers, CLI arg parsing, file reads, subprocess output, DB queries. Internal function calls between your own modules do not need `isinstance` checks, `assert not None`, or defensive `try/except`.

If an internal caller passes the wrong type, the traceback is the correct error. Wrapping every internal call in defensive code hides real bugs and doubles the diff.

### 7. Error handling is code, not decoration

- No bare `except:`. No `except Exception: pass`. If you catch, act on it (log with context, re-raise, return a fallback the caller expects).
- Catch the *narrowest* exception that matches (`FileNotFoundError`, not `OSError`, not `Exception`).
- Don't wrap code in `try` that cannot throw. `try: x = 1 + 1` is noise.
- Don't `raise Exception("...")` — use `ValueError`, `KeyError`, or a small custom subclass. `Exception` catches nothing useful and communicates nothing to the caller.
- No `traceback.print_exc()` in library code. Let it propagate. Print in the outermost handler if you must.

### 8. Comments explain *why*, never *what*

If a well-named function or variable already tells the reader *what*, a comment repeating it is negative-signal — it decays as the code changes.

Write a comment only when a future reader would be surprised by the code without it: a hidden performance constraint, a workaround for a specific library bug (link the issue), a non-obvious invariant, a business rule that isn't obvious from the domain.

No file headers ("Author: … Date: …"). No section banners (`# ==== helpers ====`). No `# noqa` without the specific rule code and a one-line reason.

---

## Anti-patterns to actively rewrite

When you see these in existing code you're editing, rewrite them in the same commit (they're cheap):

- `if x == True:` → `if x:`. Same for `== None` → `is None`.
- `list(dict.keys())` in a loop → iterate the dict directly.
- `.get(key, None)` → `.get(key)` (None is the default).
- Nested `if` chains that could be a single boolean → collapse them.
- `def foo(x, y=None): if y is None: y = []` → use a factory: `def foo(x, y=None): y = y or []` if the empty-check is safe, or move to `dataclass(default_factory=list)`.
- Classes with `__init__` that only assigns args → dataclass.
- Try/except around `dict[key]` when `.get()` is what you meant.
- `assert` in production code paths — use it in tests only; `python -O` strips it.

---

## The senior's checklist before committing

Ask yourself, in order:

1. Can I delete anything? (dead imports, unused params, one-off helpers.)
2. Is any abstraction I added justified by ≥3 call sites?
3. Would the traceback be *more* useful than the try/except I wrote?
4. Does this comment tell the reader something the code doesn't?
5. Is there stdlib for this? (Check before green-lighting a new dep.)

If yes to any, edit before you commit.

---

## What ponytail is not

- Not a linter. It's a style filter on new code, not a codemod on the existing repo.
- Not anti-testing — write tests. Just don't write ceremonial `TestFoo` scaffolding for one assertion.
- Not anti-typing — type public APIs and dataclasses. Skip obvious local hints.
- Not anti-abstraction *in principle* — anti-*premature* abstraction. Extract at 3+ call sites, not before.

---

## Credit

The Python-specific fork of [Ponytail](https://github.com/mrgoonie/ponytail), the viral "lazy senior developer" coding skill.
