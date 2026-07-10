#!/bin/bash
# First, pull the code-specialized 3B variant in your terminal if you don't have it:
# ollama pull qwen2.5-coder:3b-instruct
OLLAMA_API_BASE=http://localhost:11434 aider --model ollama_chat/qwen2.5-coder:3b-instruct
