package main

import (
	"log"
	"net/http"
	"orchestrator_llm/api"
	"orchestrator_llm/llm"
	"orchestrator_llm/orchestrator"
	"orchestrator_llm/rag"
)

func main() {

	log.Println("🚀 Starting LLM Anti-Corruption Layer on :9000")

	ragA := &rag.Client{BaseURL: "http://app-a:8080"}
	ragB := &rag.Client{BaseURL: "http://app-b:8080"}

	llmClient := &llm.Client{
		BaseURL: "http://localhost:11434/api/generate",
		Model:   "gemma3",
	}

	service := &orchestrator.Service{
		RagA: ragA,
		RagB: ragB,
		LLM:  llmClient,
	}

	handler := &api.Handler{Service: service}

	http.HandleFunc("/ask", handler.Ask)
	http.HandleFunc("/health", handler.HealthOK)

	log.Println("✅ Server ready, waiting for requests...")
	log.Fatal(http.ListenAndServe(":9000", nil))

}
