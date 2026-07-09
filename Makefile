# Define the Python helper function at the top of the Makefile
json_escape = python3 -c 'import sys, json; print(json.dumps(sys.stdin.read()))'

.PHONY: model-init model-build model-run

model-init:
	ollama create qwen-yolo -f ./Modelfile

model-build:
	ollama create qwen-yolo -f ./Modelfile

# The escape helper runs safely by constructing the JSON payload via a shell variable
model-run:
	@PROMPT_DATA=$$( $(json_escape) < mock_tef_data.txt ); \
	curl -s http://localhost:11434/api/generate -d "{\
		\"model\": \"qwen-yolo\",\
		\"prompt\": $$PROMPT_DATA,\
		\"stream\": false\
	}" | jq -r '.response'
