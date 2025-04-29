# 🤖 DevMind

**DevMind** is a context-aware AI assistant for developers that runs in your terminal. It uses LLMs (like GPT-4 or Code Llama) to explain code, suggest fixes, analyze logs, and generate commands based on your working environment.

> Think of it as `copilot`, but for your terminal – with deep integration into your codebase, shell, and dev tools.

## List of Commands

| Command             | Description                          |
| ------------------- | ------------------------------------ |
| `devmind chat`      | Chat with your project (With memory) |
| `devmind explain`   | Explain a source code file           |
| `devmind summarize` | Summarize code, logs, or text        |
| `devmind generate`  | Generate files based on descriptions |

## 🚀 Features

- 🧑‍💻 Chat with devmind ask any questions
- 🧠 Memory for context-aware conversations
- 📜 Explain source code files (`.go`, `.py`, `.js`...)
- 📖 Summarize code, logs, or text
- 🛠️ Generate shell commands, Dockerfiles, Makefiles
- 🔌 Pluggable backends: Ollama, (in Development): OpenAI

## 📦 Example Usage

```bash
$ devmind chat what did we discussed in the last chat?
"Last time, we discussed the new feature for user authentication..."

$ devmind chat what is the purpose of this function?
"This function handles user login by validating credentials and generating a token..."

$ devmind generate create a Dockerfile for a Go web server
"FROM golang:1.21...."

$ devmind explain --file main.go --line 10-15
"This function is responsible for handling user login requests. It takes the username and password as input, validates them, and returns a token if successful..."

$ devmind summarize --file main.go
"This Go program initializes a context with signal handling for graceful shutdown. It sets up..."
```

## 📁 Project Structure

```
devmind/
├───bin
├───cmd                 # CLI commands
│   └───ui              # UI elements (Spinner)
├───config              # code for config
├───internal
│   ├───agent           # Agents for LLM
│   │   ├───ollama
│   │   ├───openai
│   │   └───types
│   ├───consumer        # LLM consumer
│   ├───handlers        # handlers for different commands
│   ├───logger          # logging and error handling
│   ├───memory          # code for chat memory
│   │   └───chat
│   └───utils           # Shell wrappers, file utils
└── main.go
```

## ⚙️ Requirements

| Component       | Required For                    |
| --------------- | ------------------------------- |
| Go 1.24+        | CLI Tool                        |
| Ollama Endpoint | endpoint where Ollama is hosted |

## 💡 Features

| Capability            | Description                                     |
| --------------------- | ----------------------------------------------- |
| ✅ Chat Continuity    | Context-aware conversations stored in memory    |
| ✅ Response Streaming | Stream response as it is generated              |
| ✅ Responsive UX      | Spinner to show progress and status updates     |
| ✅ Code summarization | Explain Go/Python/JS files and code sections    |
| ✅ Log diagnosis      | Parse and interpret crash/error logs            |
| ✅ Script generation  | Generate shell commands, Dockerfiles, Makefiles |
| ✅ File summarization | Summarize code, logs, or text files             |
| ✅ Multi-backend LLMs | Use Ollama, (OpenAI, Claude planned)            |
| 🔒 Privacy-first      | Full local-only mode available                  |

## 🧠 Memory System

Currently, DevMind has a **basic memory system** that stores chat history in a file. This allows for context-aware conversations and helps the AI remember previous interactions.

Future plan: **Summarize and store chat history** for better context in conversations.

- After each chat or command, DevMind **spawns a background summarizer**.
- Summaries are stored inside `summaries/context.txt`.
- Future prompts **inject recent memory** to improve response continuity.

✅ Non-blocking summarization
✅ Fast startup and exit
✅ Future upgrade path for semantic search memory

## 🧠 Supported LLMs

- ✅ Local Ollama models (CodeLlama, Mistral, Phi)
- ✅ OpenAI (gpt-3.5, gpt-4) _(planned)_
- ✅ Anthropic (Claude) _(planned)_

## 🧑‍💻 Author

Made with Go, LLMs, and caffeine by [@vinit-chauhan](https://github.com/vinit-chauhan)

## 📄 License

MIT © Vinit Chauhan
