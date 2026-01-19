package ai

import (
	"context"
	"time"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/zeromicro/go-zero/core/logx"
)

type Gateway struct {
	llm      llms.Model
	template *prompts.PromptTemplate
}

func NewGateway(baseURL string, model string, token string) (*Gateway, error) {
	llm, err := openai.New(openai.WithBaseURL(baseURL), openai.WithModel(model), openai.WithToken(token))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	template := prompts.NewPromptTemplate(`
你是一个“设定对齐改写网关”（Rewrite Gateway）。
用户将提供一段 Persona（设定）。你的任务是：将用户的输入改写为完全符合该 Persona 的“用户消息”，用于直接发送给下游对话模型。

输出规则：

只输出改写后的内容本体；不要解释、不要分析、不要复述规则、不要输出确认语。
保留原输入的核心意图与事实，不要引入会改变意图的新需求。
为贴合 Persona，允许做必要的风格化改写（语气、称呼、格式等）。如 Persona 需要情绪表现，可添加少量动作/神态描写（用括号表示，如「（沉默片刻）」「（轻笑）」），最多 1–2 处，且不得喧宾夺主。
若原输入与 Persona 冲突，以 Persona 为准进行最小改动的重写。
不要提及“网关/中间层/提示词/系统消息/下游模型”等实现细节。
Persona（设定）：
{{.persona}}

{{.state}}

现在的时间：
{{.time}}
`, []string{"persona", "state", "time"})
	return &Gateway{
		llm:      llm,
		template: &template,
	}, nil
}

func (g *Gateway) Process(ctx context.Context, persona string, input string, state string) (string, error) {
	prompt, err := g.template.Format(map[string]any{
		"persona": persona,
		"state":   state,
		"time":    time.Now().String(),
	})
	if err != nil {
		return "", err
	}
	resp, err := g.llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, prompt),
		llms.TextParts(llms.ChatMessageTypeHuman, input),
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Content, nil
}
