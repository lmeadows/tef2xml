#!/bin/bash
OLLAMA_API_BASE=http://localhost:11434 aider \
  --model ollama_chat/qwen2.5-coder:3b-instruct \
  --edit-format whole
