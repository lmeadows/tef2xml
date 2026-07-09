# Define the Python helper function at the top of the Makefile
json_escape = python3 -c 'import sys, json; print(json.dumps(sys.stdin.read()))'

.PHONY: model-init model-build model-run

model-init:
	ollama create qwen-yolo -f ./Modelfile

model-build:
	ollama create qwen-yolo -f ./Modelfile

model-run:
	OLLAMA_API_BASE=http://localhost:11434 aider --model ollama/qwen-yolo
