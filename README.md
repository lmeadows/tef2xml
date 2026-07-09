# tef2xml

A cross-platform CLI tool written in Go that extracts proprietary TablEdit (`.tef`) banjo tablature data and exports it to standard, portable MusicXML format using a local LLM runtime (`ollama`).

## Project Architecture & Roadmap

To ensure high reliability and eliminate LLM syntax drift, this project enforces a strict boundary: the local LLM handles semantic data normalization (converting raw bytes/tokens to musical concepts), while native Go code handles rigid data serialization (converting concepts to structured MusicXML blocks).
