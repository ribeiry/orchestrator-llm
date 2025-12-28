package orchestrator

import (
	"orchestrator_llm/domain"
	"orchestrator_llm/llm"
	"orchestrator_llm/prompt"
	"orchestrator_llm/rag"
)

type Service struct {
	RagA *rag.Client
	RagB *rag.Client
	LLM  *llm.Client
}

func (s *Service) Ask(req domain.AskRequest) (string, error) {
	var ctxA, ctxB []string

	if req.Context == domain.ContextA || req.Context == domain.ContextBoth {
		chunks, _ := s.RagA.Search(req.Question)
		for _, c := range chunks {
			ctxA = append(ctxA, c.Content)
		}
	}
	if req.Context == domain.ContextB || req.Context == domain.ContextBoth {
		chunks, _ := s.RagB.Search(req.Question)
		for _, c := range chunks {
			ctxB = append(ctxB, c.Content)
		}
	}
	p := prompt.Build(req.Question, ctxA, ctxB)

	return s.LLM.Generate(p)

}
