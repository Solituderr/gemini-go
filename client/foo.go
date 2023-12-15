package client

type Text struct {
	Text string `json:"text"`
}

type Content struct {
	Role  string `json:"role"`
	Parts []Text `json:"parts"`
}

type Messages struct {
	Contents []Content `json:"contents"`
}

// 返回结构体
type Resp struct {
	Candidates []Content `json:"candidates"`
}
