# Workshop: Let's Write a JSX Compiler from Scratch in Go

## Pre-Talk Setup (do this before the talk!)

**Prerequisites:**
- Go 1.22+: https://go.dev/dl/ — verify with `go version`
- Git — verify with `git --version`
- VS Code + Go extension (recommended)

**Setup:**
```bash
git clone https://github.com/YOUR_USERNAME/jsxc
cd jsxc
git checkout start
go run .
```

**Expected output:**
```
panic: not implemented
```

If you see that panic → **your setup is correct.** The skeleton is running and waiting for you to implement it.

---

## Catch-Up Commands

If you fall behind at any stage, run the matching command to jump to the current position. Then verify your output and continue from there.

| Just finished | Catch-up command | Verify with |
|---|---|---|
| Stage 1 (Tokens) | `git stash && git checkout checkpoint/1-tokens` | `go run .` → panics at NextToken (expected) |
| Stage 2 (Lexer) | `git stash && git checkout checkpoint/2-lexer` | `STAGE=2 go run .` → prints token stream |
| Stage 3 (Parser) | `git stash && git checkout checkpoint/3-parser` | `STAGE=3 go run .` → prints AST |
| Stage 4 (Generator) | `git stash && git checkout checkpoint/4-generator` | `go run .` → writes output.html |

---

## Verification Commands

Use these at each stage to confirm your implementation is working:

```bash
# After Stage 2 — should print the full token stream
STAGE=2 go run .

# After Stage 3 — should print the nested AST struct
STAGE=3 go run .

# After Stage 4 — should write output.html
go run .
# Then open output.html in your browser
```

---

## What We're Building

```
Input:  <div><h2>Genetec</h2><button id="myBtn">Click me</button></div>

Output: React.createElement("div", null,
          React.createElement("h2", null, "Genetec"),
          React.createElement("button", { id: "myBtn" }, "Click me")
        )
```

The 3-stage pipeline:
```
Raw string → [LEXER] → tokens → [PARSER] → AST → [GENERATOR] → JS string
```

---

## Stage Guide

### Stage 1: Tokens (`token.go`)
Define the 8 token types. Look at the input string — what distinct characters or groups need a type?

```
< → OpenAngle    > → CloseAngle   / → Slash        = → Equals
div → Identifier  "myBtn" → String  Genetec → Text   (end) → EOF
```

### Stage 2: Lexer (`lexer.go`)
Implement `NextToken()`. The key insight: use `inTag bool` to know whether a character is markup or text content.

```
[outside tag] --<--> [inside tag]
```

### Stage 3: Parser (`parser.go`)
Implement `parseElement()`, `parseChildren()`, `parseProps()`.

The insight: `parseChildren()` calls `parseElement()` which calls `parseChildren()`. The call stack IS the parse tree.

### Stage 4: Generator (`generator.go`)
Implement `generate()`. Walk the AST and emit `React.createElement(...)`.

The insight: lowercase tag → string literal. Uppercase → identifier reference (same rule as Babel).

---

## Going Further

After the talk, explore the `main` branch for the complete solution with comments.

Ideas to extend the compiler:
- **Self-closing tags**: `<img src="x.png" />` — one `if` branch in `parseElement()`
- **JSX expressions**: `<div>{name}</div>` — new token types + AST node
- **Fragments**: `<>...</>` — special case in the parser
- **Multiple root elements**: change `parseElement()` return type

---

## Resources

- [Babel AST Explorer](https://astexplorer.net/) — see what real JSX compiles to
- [Go Tour](https://go.dev/tour/) — learn Go interactively
- [Crafting Interpreters](https://craftinginterpreters.com/) — free book on compiler construction
