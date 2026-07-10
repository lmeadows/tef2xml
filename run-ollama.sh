#!/bin/bash
# Local environment variables scoped only to this running process
export OLLAMA_NUM_PARALLEL=1
export OLLAMA_KEEP_ALIVE=-1

sudo systemctl stop ollama

# Change the port if you don't want to clash with a global Ollama instance (optional)
# export OLLAMA_HOST="127.0.0.1:11435" 

echo "Starting isolated Ollama instance for this project..."
exec ollama serve
