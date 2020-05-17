package request

type ReqPath struct {
	Path string `json:"path"`
}

type ReqClearNode struct {
	NodeId string `json:"nodeId"`
	TestPaperID uint `json:"testPaperID"`
	Type string `json:"type"`
	Path string `json:"path"`
}


