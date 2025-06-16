# Commitlint

**A lightweight, fast, and cross-platform CLI tool for linting Git commit messages.**

Linting commit messages helps maintain a consistent commit history, which is critical for readability, automation, and collaboration across teams. `commitlint` ensures your commits follow a defined convention, making your Git logs cleaner and easier to work with.

## Features

- ⚡ **Fast and lightweight**: Built for performance.
- 💻 **Cross-platform support**: Works on Windows, macOS, and Linux.
- 📦 **Easy installation**: Install with a single command per OS.
- 🔧 **Project-level integration**: Add to your project with `commitlint install`.

---

## Getting Started

### Step 1: Install the CLI

Choose the appropriate method for your operating system:

#### Windows

1. Go to the [Releases](https://github.com/jurienhamaker/commitlint/releases) page.
2. Download the Windows installer (`.exe`) for the latest version.
3. Run the installer to add `commitlint` to your system.

#### macOS

Install using Homebrew:

```bash
brew install --cask jurienhamaker/commitlint-go/commitlint-go
```

#### Linux (Arch-based systems)

Install from the AUR (using an AUR helper like `yay` or `paru`):

```bash
yay -S commitlint-go-bin
```

---

### Step 2: Initialize in Your Repository

In your project directory, run:

```bash
commitlint install
```

This command adds `commitlint` to your repository by installing necessary hooks and configuration to lint commit messages before they are created or pushed.

---

## Usage

Once installed, `commitlint` automatically lints commit messages via Git hooks. If a commit message does not meet the expected format, the commit is rejected with a descriptive error.

You can also run it manually on a commit message:

```bash
echo "commit message" | commitlint lint
```

---

## Configuration

Currently, `commitlint` ships with a opinionated set of rules based on the conventional commit format.
These rules can be changed in the `.commitlint/commitlint.yml` file

**Available rules:**

```yaml
body-case: [2, "always", "sentencecase"]
body-empty: [2, "always"]
body-full-stop: [2, "always"]
body-leading-blank: [2, "always"]
body-max-length: [2, "always", 100]
body-max-line-length: [2, "always", 100]
body-min-length: [2, "always", 5]
footer-empty: [2, "always"]
footer-leading-blank: [2, "always"]
footer-max-length: [2, "always", 100]
footer-max-line-length: [2, "always", 100]
footer-min-length: [2, "always", 5]
header-case: [2, "always", "lowercase"]
header-full-stop: [2, "always", "."]
header-max-length: [2, "always", 5]
header-min-length: [2, "always", 100]
header-trim: [2]
references: [2, "always", ["#"]]
scope-case:
  [
    2,
    "always",
    [
      "upper-case",
      "lower-case",
      "pascal-case",
      "kebab-case",
      "snake-case",
      "sentence-case",
    ],
  ]
scope-empty: [2, "never"]
scope-enum: [2, "always", "scope"]
scope-max-length: [2, "always", 100]
scope-min-length: [2, "always", 5]
signed-off-by: [2, "always"]
subject-case:
  [
    2,
    "always",
    [
      "upper-case",
      "lower-case",
      "pascal-case",
      "kebab-case",
      "snake-case",
      "sentence-case",
    ],
  ]
subject-empty: [2, "never"]
subject-full-stop: [2, "always"]
subject-max-length: [2, "always", 100]
subject-min-length: [2, "always", 5]
type-case: [2, "always", "lowercase"]
type-empty: [2, "never"]
type-enum:
  [
    2,
    "always",
    [
      "build",
      "chore",
      "ci",
      "docs",
      "feat",
      "fix",
      "perf",
      "refactor",
      "revert",
      "style",
      "test",
    ],
  ]
type-max-length: [2, "always", 100]
type-min-length: [2, "always", 2]
```

**Available cases:**

```
upper-case
uppercase
lower-case
lowercase
pascal-case
pascalcase
kebab-case
kebabcase
snake-case
snakecase
sentence-case
sentencecase
```

---

## Contributing

We welcome contributions! Please fork the repository and submit a pull request. Be sure your commits follow the linting rules — after all, that's the point!

---

## Troubleshooting

- **CLI not found after install?**  
  Make sure the install location is in your system’s `PATH`.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Links

- [GitHub Repository](https://github.com/jurienhamaker/commitlint)
- [Releases](https://github.com/jurienhamaker/commitlint/releases)
- [Conventional Commits](https://www.conventionalcommits.org/)
