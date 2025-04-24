# Glyphra 🧠📦

![Go Version](https://img.shields.io/github/go-mod/go-version/danny-molnar/glyphra)
![License](https://img.shields.io/github/license/danny-molnar/glyphra)
![Last Commit](https://img.shields.io/github/last-commit/danny-molnar/glyphra)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/danny-molnar/glyphra/go.yml?label=CI)

![Made with ❤️](https://img.shields.io/badge/built_with-love-blueviolet?style=flat-square)

**Glyphra** is an AI-powered CLI tool that explains Terraform plans — because reading giant diffs should *not* be a full-time job.

Built with Go, powered by Hugging Face, and documented via its own blog, Glyphra is your dev sidekick when infra gets noisy.

---

## 🚀 What It Does

- Reads Terraform plan output
- Summarizes it using an LLM (like `flan-t5-large`)
- Outputs a human-readable summary
- Can generate Markdown blog posts from summaries

---

## 📦 Installation

You can build it from source:

```bash
git clone https://github.com/danny-molnar/glyphra.git
cd glyphra
make build
./glyphra --help
```

---

## 🔧 Requirements

- Go 1.21+
- Hugging Face API token (free tier is fine)

You can store the token in `.envrc` or export it directly:

```bash
export HUGGINGFACE_TOKEN=your_token_here
```

---

## 🛠 Usage

### Summarize a plan:

```bash
make plan        # Generates a fake Terraform plan
make run         # Runs Glyphra on plan.out
```

### Save to blog:

```bash
make blog        # Summarizes and saves to Markdown blog post
```

---

## 🧪 Development

- `make test` — run Go tests
- `make lint` — run basic linters
- `make clean` — remove build artifacts
- `make docker-build` — build Docker image

---

## 📚 Blog

Check out the dev journey:
👉 https://danny-molnar.github.io/glyphra

Source files live in the `/docs` folder and use the Minimal Mistakes theme with a dark skin.

---

## 🤔 Why?

Because infrastructure shouldn't be cryptic.
And AI can make tooling more helpful *and* more fun.

---

## 💡 What's Next?

- Support for `terraform show -json`
- GitHub PR summarization (`glyphra pr`)
- Local AI model support via Ollama
- Auto-tagging and blog post scheduling

---

## 🧑‍💻 Author

Built by [@danny-molnar](https://github.com/danny-molnar)
Infra dev turned CLI engineer with a love for 🍜, AI, and dark mode everything.

---

## 📄 License

MIT License — use it, remix it, build cool tools with it.
See [LICENSE](./LICENSE) for details.
