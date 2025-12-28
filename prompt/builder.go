package prompt

import "strings"

func Build(question string, ctxA, ctxB []string) string {

	var sb strings.Builder

	sb.WriteString(
		`Você é um assistente.
		Use SOMENTE os contextos fornecidos.
		Se não souber, diga que não sabe.`)

	if len(ctxA) > 0 {
		sb.WriteString("\n[CONTEXT_APP_A]\n")
		for _, c := range ctxA {
			sb.WriteString("-" + c + "\n")
		}
	}

	if len(ctxB) > 0 {
		sb.WriteString("\n[CONTEXT_APP_A]\n")
		for _, c := range ctxB {
			sb.WriteString("-" + c + "\n")
		}
	}

	sb.WriteString("\nPERGUNTA:\n" + question)
	return sb.String()

}
