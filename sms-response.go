package aliyunsmsclient

// The response code which stands for a sms is sent successfully.
const ResponseCodeOk = "OK"

// @see https://help.aliyun.com/document_detail/55284.html#出参列表
// The Response of sending sms API.
type Response struct {
	// The raw response from server.
	RawResponse []byte `json:"-"`
	/* Response body */
	RequestId string `json:"RequestId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	BizId     string `json:"BizId"`
}

func (m *Response) IsSuccessful() bool {
	return m.Code == ResponseCodeOk
}
