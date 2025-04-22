package llm

type LLMClient interface {
    Explain(input string) (string, error)
}
