# Gommit

**Gommit** is a CLI tool that generates **Git commit messages using AI** based on your repository's **diff and commit history**.

It is designed to help developers keep **clean, consistent and conventional commit messages** without wasting time writing them manually.

The tool analyzes:

* `git diff`
* optional `git log` history
* your commit pattern prompt

Then generates a commit message following your configured style.

---

# Features

* AI-generated commit messages
* Optional commit history context
* Automatic commit execution
* Optional push after commit
* `--force` push support
* Configurable commit pattern prompt
* Lightweight CLI built with Go

---

# Installation

## Releases

Go to the last release and get the installer bin.

## Build from source

```bash
git clone https://github.com/Yyax13/Gommit
cd Gommit

go build -o gommit
```

Move binary somewhere in your `$PATH` if desired:

```bash
sudo mv gommit /usr/local/bin/
```

---

# Configuration

Gommit automatically creates a configuration file if it does not exist.

Default location:

```
~/.config/gommit/settings.yaml
```

Example config:

```yaml
GeminiApiKey: ""
UseHist: false
CommitPatternPrompt: ""
OverwriteDefaultCommitPatternPrompt: false
```

### Fields

| Field                                 | Description                                       |
| ------------------------------------- | ------------------------------------------------- |
| `GeminiApiKey`                        | API key used for AI generation                    |
| `UseHist`                             | Include `git log` history when generating commits |
| `CommitPatternPrompt`                 | Custom prompt defining commit style               |
| `OverwriteDefaultCommitPatternPrompt` | If true, replaces the built-in commit prompt      |

---

# Usage

Generate a commit message:

```bash
gommit gen
```

Output example:

```
Generated commit message:
feat(auth): add JWT validation middleware
```

By default **it only prints the message**.

---

# Flags

## Commit automatically

```bash
gommit gen --commit
```

or

```bash
gommit gen -c
```

---

## Commit and push

```bash
gommit gen -c -p
```

---

## Push to specific branch

```bash
gommit gen -c -p --branch main
```

---

## Force push

```bash
gommit gen -c -p --force
```

---

# Examples

Generate message only

```bash
gommit gen
```

Generate and commit

```bash
gommit gen -c
```

Generate, commit and push

```bash
gommit gen -c -p
```

Force push

```bash
gommit gen -c -p --force
```

---

# How it works

1. Gommit reads your **git diff**
2. Optionally loads **commit history**
3. Sends context to the configured AI model
4. Generates a commit message
5. Optionally runs:

```
git commit
git push
```

---

# Tech Stack

* Go
* Cobra (CLI framework)
* Viper (configuration loader)

---

# Contributing

Contributions are welcome.

You can help by:

* improving prompts
* adding providers
* improving git analysis
* fixing bugs

---

# License

[MIT License](LICENSE)