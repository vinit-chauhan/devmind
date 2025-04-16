# 🤖 DevMind

**DevMind** is a context-aware AI assistant for developers that runs in your terminal. It uses LLMs (like GPT-4 or Code Llama) to explain code, suggest fixes, analyze logs, and generate commands based on your working environment.

> Think of it as `copilot`, but for your terminal – with deep integration into your codebase, shell, and dev tools.

## 🚀 Features
- 🧠 Explain source code files (`.go`, `.py`, `.js`...)
- 🪵 Analyze and explain logs / error messages
- 🧪 Suggest shell fixes for Git errors, Docker issues, etc.
- 🧬 Contextual awareness: env vars, open files, Git diffs
- 🔌 Pluggable backends: OpenAI, Ollama, Claude
- 💻 CLI interface and optional local WebView UI
- 🔧 Local inference with LLMs via Ollama (e.g., Code Llama)

## 📦 Example Usage
```bash
$ devmind explain ./internal/api/handler.go
"This file defines a REST API handler for managing user login tokens..."

$ devmind fix ./crash.log
"Postgres connection refused. Try checking DB_HOST or restart Docker."

$ devmind run 'generate a Dockerfile for this Go project'
"Here is a Dockerfile optimized for Go 1.21 and multi-stage builds..."
```

## 📁 Project Structure
```
devmind/
├── cmd/                  # CLI commands
│   └── root.go
├── pkg/
│   ├── llm/              # LLM integrations (OpenAI, Ollama)
│   ├── context/          # Env, Git diff, AST extractors
│   ├── prompt/           # Prompt templates
│   └── render/           # TUI/WebView output
├── internal/
│   └── util/             # Shell wrappers, file utils
├── scripts/              # Setup / inference tools
├── models/               # Local models (GGUF/Ollama)
├── webui/                # Optional HTML UI
└── main.go
```

## ⚙️ Requirements
| Component         | Required For            |
| ----------------- | ----------------------- |
| Go 1.21+          | CLI Tool                |
| OpenAI API Key    | GPT-3.5/4 access        |
| Ollama            | Local LLM inference     |
| Git               | Context diff extraction |
| Docker (optional) | For containerized setup |

## 💡 Features
| Capability              | Description |
|-------------------------|-------------|
| ✅ Code summarization   | Explains Go/Python/JS files and diffs |
| ✅ Log diagnosis        | Parses and interprets crash/error logs |
| ✅ Script generation    | Creates shell commands, Dockerfiles, Makefiles |
| ✅ Git context awareness| Uses diffs and staged files as context |
| ✅ Env introspection    | Optionally includes sanitized env vars |
| ✅ LLM flexibility      | Use OpenAI, Ollama, or local HTTP models |
| ✅ Token-aware chunking | Dynamically breaks large files into model-friendly input |
| ✅ Interactive REPL     | Chat with your project locally |
| 🔒 Privacy-first        | No telemetry; local-only mode with Ollama available |


## 🧠 Supported LLMs
- ✅ OpenAI (gpt-3.5, gpt-4)
- ✅ Local Ollama models (CodeLlama, Mistral, Phi)
- ✅ Anthropic (Claude) *(planned)*


## 🧑‍💻 Author
Made with Go, LLMs, and caffeine by [@vinit-chauhan](https://github.com/vinit-chauhan)

## 📄 License
MIT © Vinit Chauhan
