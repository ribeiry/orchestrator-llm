# 🧠 LLM Anti-Corruption Layer (RAG Orchestrator)

Este projeto implementa uma **camada de Anti-Corruption para LLMs**, responsável por orquestrar múltiplos **RAGs (Retrieval-Augmented Generation)** distintos sobre **um único modelo LLM** (Ollama + Gemma).

O objetivo é permitir que **diferentes aplicações**, cada uma com seu **domínio e contexto próprios**, compartilhem o mesmo LLM **sem acoplamento direto**, mantendo controle de escopo, qualidade e governança.

---

## 🎯 Objetivo

* Centralizar o acesso a LLMs
* Suportar **múltiplos contextos (RAGs)** para o mesmo modelo
* Evitar alucinação e respostas fora de domínio
* Proteger aplicações contra mudanças de modelo ou provider

---

## 🏗️ Visão Geral da Arquitetura

```
┌────────────┐       ┌────────────┐
│   App A    │       │   App B    │
│  (RAG A)   │       │  (RAG B)   │
└─────┬──────┘       └─────┬──────┘
      │                    │
      └──────┬─────────────┘
             │
     ┌───────▼──────────────────┐
     │  LLM Anti-Corruption ACL  │
     │  (Orchestrator + Policy)  │
     └──────────┬────────────────┘
                │
         ┌──────▼───────┐
         │   Ollama     │
         │  Gemma LLM   │
         └──────────────┘
```

### Responsabilidades da ACL

* Selecionar o **RAG correto** (`A` ou `B`)
* Aplicar **system prompts específicos por domínio**
* Validar qualidade do contexto antes do LLM
* Evitar alucinação
* Isolar as aplicações do modelo LLM

---

## 📁 Estrutura do Projeto

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/            # Handlers HTTP
│   ├── orchestrator/   # Core: seleção de RAG + prompt + LLM
│   ├── rag/            # Clients para RAG A e RAG B
│   ├── llm/            # Client do Ollama
│   └── domain/         # DTOs de request/response
└── README.md
```

---

## 🚀 Como Executar

### Pré-requisitos

* Go **1.22+**
* Ollama rodando localmente
* Modelo carregado no Ollama (ex: `gemma-backend`)
* Serviços de RAG A e RAG B disponíveis (ou mocks)

### Subir o Ollama

```bash
ollama run gemma-backend
```

### Rodar a aplicação

```bash
go run cmd/server/main.go
```

Servidor disponível em:

```
http://localhost:9000
```

---

## 🩺 Health Check

```bash
curl http://localhost:9000/health
```

Resposta esperada:

```
OK
```

---

## ❓ Endpoint `/ask`

### POST `/ask`

### Request

```json
{
  "question": "Como funciona o cache distribuído?",
  "context": "A",
  "options": {
    "temperature": 0.3,
    "max_tokens": 400
  }
}
```

### Campos

| Campo      | Descrição                             |
| ---------- | ------------------------------------- |
| `question` | Pergunta do usuário                   |
| `context`  | Contexto/RAG a ser usado (`A` ou `B`) |
| `options`  | Configurações opcionais do LLM        |

---

### Exemplo via `curl`

```bash
curl -X POST http://localhost:9000/ask \
  -H "Content-Type: application/json" \
  -d '{
    "question": "Explique circuit breaker em microsserviços",
    "context": "B"
  }'
```

---

### Response

```json
{
  "answer": "O circuit breaker é um padrão de resiliência...",
  "model": "gemma-backend",
  "context_used": "B",
  "latency_ms": 812
}
```

---

## 🧠 Funcionamento do Orchestrator

1. Recebe a requisição `/ask`
2. Identifica o contexto solicitado (`A` ou `B`)
3. Consulta o RAG correspondente
4. Valida:

   * existência de chunks
   * score mínimo de similaridade
5. Aplica o **system prompt do domínio**
6. Envia prompt final ao LLM
7. Retorna resposta ou recusa explícita

---

## 🛡️ Controle de Alucinação

O LLM é explicitamente instruído a:

* Responder **apenas** com base no contexto fornecido
* Recusar perguntas fora do domínio
* Não inferir ou inventar informações

Exemplo de resposta de recusa:

```
"Não tenho informações suficientes no contexto para responder essa pergunta."
```

---

## 🔐 Por que Anti-Corruption Layer?

Sem essa camada:

* Aplicações ficam acopladas ao LLM
* Troca de modelo impacta todo o sistema
* Não há controle de domínio

Com essa camada:

* LLM vira infraestrutura
* Domínios permanecem isolados
* Evolução segura para outros providers

---

## 🧪 Status

* ✔️ Prova de conceito
* ✔️ Multi-RAG
* ✔️ Single LLM
* ✔️ Arquitetura limpa
* ✔️ Pronto para evolução

---

## 🔮 Próximos Passos

* Similarity threshold configurável
* Streaming de resposta
* Observabilidade (latência, tracing)
* `/health/ready` com checagem real
* Cache de respostas
* Autenticação e rate limit

---

## 🧩 Resumo

> **LLMs são poderosos, mas sem governança viram caos.**
> **Esta camada existe para impor ordem.**
