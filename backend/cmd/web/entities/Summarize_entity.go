package entities

type Message struct {
	Role      string        `json:"role"`
	Content   string        `json:"content"`
	ToolCalls []interface{} `json:"tool_calls"` // Changed to a slice
}

type Choice struct {
	Index        int         `json:"index"`
	Message      Message     `json:"message"`
	Logprobs     interface{} `json:"logprobs"` // Keep this as is to handle null
	FinishReason string      `json:"finish_reason"`
	StopReason   *string     `json:"stop_reason"` // Keep this as is for handling null
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
	CompletionTokens int `json:"completion_tokens"`
}

type ChatResponse struct {
	ID             string      `json:"id"`
	Object         string      `json:"object"`
	Created        int64       `json:"created"`
	Model          string      `json:"model"`
	Choices        []Choice    `json:"choices"`
	Usage          Usage       `json:"usage"`
	PromptLogprobs interface{} `json:"prompt_logprobs"` // Keep this as is to handle null
}
