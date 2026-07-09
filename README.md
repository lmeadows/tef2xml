# tef2xml

A cross-platform CLI tool written in Go that extracts proprietary TablEdit (`.tef`) banjo tablature data and exports it to standard, portable MusicXML format using a local LLM runtime (`ollama`).

## Project Architecture & Roadmap

To ensure high reliability and eliminate LLM syntax drift, this project enforces a strict boundary: the local LLM handles semantic data normalization (converting raw bytes/tokens to musical concepts), while native Go code handles rigid data serialization (converting concepts to structured MusicXML blocks).

### Phase 1: The Golden Master & Structural Normalization (Current Goal)
* **Objective:** Define a rigid intermediate JSON schema ("Golden Master") representing a simple 4-measure Scruggs-style banjo roll.
* **Goal:** Successfully prompt `qwen3.5:9b` via Ollama to take an unparsed, raw structural text dump from a `.tef` file and normalize it into this exact target JSON shape with 100% syntactic accuracy.

### Phase 2: Native Go XML Serialization Engine
* **Objective:** Map the validated intermediate JSON schema into native Go structures using `encoding/xml`.
* **Goal:** Write a deterministic Go compiler that reads the target JSON schema and spits out a schema-compliant, perfectly formatted `.musicxml` file ready to load into MuseScore.

### Phase 3: The Automated Verification Loop
* **Objective:** Build a local test harness that passes a test suite of varied `.tef` data dumps through the pipeline.
* **Goal:** Automatically validate the generated MusicXML outputs against time-signature math and the official W3C MusicXML schema definition to capture edge cases.
