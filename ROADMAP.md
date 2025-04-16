# 🛠️ DevMind Development Roadmap

A modular, LLM-powered CLI agent that helps developers understand and manipulate their code, logs, and environments.


## ✅ Phase 1: Core CLI + LLM Integration

### 🎯 Goal: Standalone CLI that can talk to an LLM with basic prompt

- [ ] Setup `main.go` with CLI entrypoint (use `cobra` or `urfave/cli`)
- [ ] Implement `llm/` package for OpenAI API integration
- [ ] Read prompt from CLI flag or stdin
- [ ] Send prompt to OpenAI and print result
- [ ] Add `.env` config loader (`viper`, `godotenv`)


## 🚧 Phase 2: Context Ingestion Layer

### 🎯 Goal: Add structured local context as input to LLM

- [ ] Parse Go/Python/JS files and extract:
  - [ ] Function names
  - [ ] Comments and signatures
- [ ] Parse Git diff (`git diff HEAD`)
- [ ] Sanitize and include environment variables
- [ ] Chunk large input into ~2K token-safe segments
- [ ] Compose final prompt with system + context + user input


## 🧪 Phase 3: Feature Commands

### 🎯 Goal: Build meaningful commands around DevMind capabilities

- [ ] `devmind explain <file>` → explain a code file
- [ ] `devmind fix <logfile>` → analyze and suggest fix
- [ ] `devmind run '<natural task>'` → generate bash/Dockerfile/etc
- [ ] `devmind summary` → summarize current Git diff + files


## 🌐 Phase 4: Multi-Model Support

### 🎯 Goal: Run with different LLM backends

- [ ] Add Ollama integration via `http://localhost:11434/api/generate`
- [ ] Add backend selector flag: `--provider=openai|ollama`
- [ ] Normalize prompt/response handling across providers


## 💻 Phase 5: Interactive REPL / Chat Mode

### 🎯 Goal: Run DevMind in conversational mode

- [ ] Implement `devmind chat` REPL
- [ ] Persist context across messages (simple in-memory memory)
- [ ] Keyboard shortcuts for history, clearing, saving


## 📊 Phase 6: Advanced UX and Plugins

- [ ] Prompt templates (stored in `.devmind/templates/*.tmpl`)
- [ ] Syntax-highlighted output (color terminal / TUI)
- [ ] Plugin support: user-defined commands via YAML + Go templates
- [ ] LLM-based code linting


## 🐳 Phase 7: Deployment & Packaging

- [ ] Dockerize CLI for isolated use
- [ ] Provide binary releases via `goreleaser`
- [ ] Publish GitHub release + installation instructions


## ✨ Stretch Goals

- [ ] Web UI (Tauri/Webview or simple browser)
- [ ] VSCode Extension
- [ ] Token budget estimator and cost display
- [ ] Local model selector via `gguf`, llama.cpp, LM Studio
- [ ] Context caching for faster re-queries


## 🔄 Milestone Summary

| Milestone        | Description                          | Status  |
|------------------|--------------------------------------|---------|
| ✅ Phase 1        | CLI + LLM Integration                | ⏳       |
| ✅ Phase 2        | Context gathering                    | ⏳       |
| ✅ Phase 3        | Useful commands (explain/fix/run)    | ⏳       |
| ✅ Phase 4        | Multi-model backend support          | ⏳       |
| ✅ Phase 5        | Chat mode                            | ⏳       |
| ✅ Phase 6        | Plugins and UI polish                | ⏳       |
| ✅ Phase 7        | Deployment & packaging               | ⏳       |

