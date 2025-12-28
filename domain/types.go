package domain

type ContextType string

const (
	ContextA ContextType = "A"
	ContextB ContextType = "B"

	ContextBoth ContextType = "BOTH"
)

type AskRequest struct {
	Question string      `json:"question"`
	Context  ContextType `json:"context"`
}

type AskResponse struct {
	Answer string `json:"answer"`
}
