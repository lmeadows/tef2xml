.PHONY: model-init model-build model-run

model-init:
	ollama create qwen2.5:3b-instruct-tef2xml -f ./Modelfile

model-build:
	ollama create qwen2.5:3b-instruct-tef2xml -f ./Modelfile


model-run:
	OLLAMA_API_BASE=http://localhost:11434 aider --model ollama_chat/qwen2.5:3b-instruct-tef2xml
