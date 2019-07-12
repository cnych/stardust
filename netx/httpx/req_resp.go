package httpx

import (
	"encoding/json"
	"encoding/xml"
	"github.com/cnych/stardust/encodingx/base64x"
	"github.com/cnych/stardust/encodingx/jsonx/freejson"
	"net/http"
)

type Request struct {
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Header     map[string]string `json:"header"`
	Base64Body string            `json:"body,omitempty"`
	Retries    int               `json:"retries,omitempty"`
	TimeoutMS  int64             `json:"timeout_ms,omitempty"`
}

type Response struct {
	StatusCode int                  `json:"status_code"`
	URL        string               `json:"url"`
	Header     map[string]string    `json:"header"`
	Base64Body string               `json:"body,omitempty"`
	Ex         map[string]*Response `json:"ex,omitempty"`
}

type ResponseWithError struct {
	*Response
	Err error
}

func NewGetReq(url string, retries int, timeoutMS int64) *Request {
	return &Request{
		Method:    http.MethodGet,
		URL:       url,
		Retries:   retries,
		TimeoutMS: timeoutMS,
	}
}

func NewPostWithHeader(url string, body []byte, header map[string]string, retries int, timeoutMS int64) *Request {
	req := &Request{
		Method:    http.MethodPost,
		URL:       url,
		Header:    header,
		Retries:   retries,
		TimeoutMS: timeoutMS,
	}
	req.SetBody(body)
	return req
}

func NewGetReqWithHeader(url string, header map[string]string, retries int) *Request {
	return &Request{
		Method:  http.MethodGet,
		URL:     url,
		Header:  header,
		Retries: retries,
	}
}

func NewGetReqs(urls []string, retries int, timeoutMS int64) []*Request {
	nUrl := len(urls)
	if nUrl == 0 {
		return nil
	}
	reqs := make([]*Request, 0, nUrl)
	for _, url := range urls {
		reqs = append(reqs, NewGetReq(url, retries, timeoutMS))
	}
	return reqs
}

func (req *Request) SetBody(data []byte) {
	if len(data) == 0 {
		return
	}
	req.Base64Body = base64x.EncodeString(data)
}

func (req *Request) BodyAsBytes() ([]byte, error) {
	if req.Base64Body == "" {
		return nil, nil
	}
	return base64x.DecodeString(req.Base64Body)
}

func (resp *Response) SetBody(data []byte) {
	if len(data) == 0 {
		return
	}
	resp.Base64Body = base64x.EncodeString(data)
}

func (resp *Response) AsBytes() ([]byte, error) {
	if resp.Base64Body == "" {
		return nil, nil
	}
	return base64x.DecodeString(resp.Base64Body)
}

func (resp *Response) ParseJSON(v interface{}) error {
	bodyBytes, err := resp.AsBytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, v)
}

func (resp *Response) ParseXML(v interface{}) error {
	bodyBytes, err := resp.AsBytes()
	if err != nil {
		return err
	}
	return xml.Unmarshal(bodyBytes, v)
}

func (resp *Response) ParseFreeJSON(valDec freejson.ValueDecoder) (interface{}, error) {
	bodyBytes, err := resp.AsBytes()
	if err != nil {
		return nil, err
	}
	return freejson.Unmarshal(bodyBytes, valDec)
}
