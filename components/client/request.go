package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

const (
	qvmRegionIDHeaderKey  = "X-Qiniu-Regionid"
	qvmRIOZoneIDHeaderKey = "X-RIO-Zone-ID"
)

// Request API 请求对象
type Request struct {
	Method  string
	URL     string
	Queries url.Values
	Headers map[string]string
	Body    []byte
}

// NewRequest 构造新的请求
func NewRequest(method string, url string) *Request {
	return &Request{
		Method:  method,
		URL:     url,
		Queries: make(map[string][]string),
		Headers: make(map[string]string),
		Body:    nil,
	}
}

// WithRegionID 设置 RegionIDs 参数
func (r *Request) WithRegionID(regionID *string) *Request {
	if regionID == nil {
		return r
	}

	r.Headers[qvmRegionIDHeaderKey] = *regionID
	return r
}

// WithRegionIDs 设置 RegionIDs 参数
func (r *Request) WithRegionIDs(regionIDs ...string) *Request {
	rids := strings.Join(regionIDs, ",")
	r.Headers[qvmRegionIDHeaderKey] = rids
	return r
}

// WithZoneID 设置 ZoneID 参数
func (r *Request) WithZoneID(zoneID *string) *Request {
	if zoneID == nil {
		return r
	}

	r.Headers[qvmRIOZoneIDHeaderKey] = *zoneID
	return r
}

// WithQueries 设置 URL Query
// 如果非 urls.Values 类型，则会进行 JSON 转换成 map[string]string 后设置
// 注意：不会处理错误，如果转换失败，则不会设置 Query
func (r *Request) WithQueries(rawQueries interface{}) *Request {
	if rawQueries == nil {
		return r
	}

	switch q := rawQueries.(type) {
	case url.Values:
		r.Queries = q
	case map[string][]string:
		r.Queries = url.Values(q)
	default:
		transQ, err := transObject2Queries(rawQueries)
		if err == nil {
			r.Queries = transQ
		}
	}
	return r
}

// WithJSONBody 设置 JSON body 内容
// 1 将参数转换为 JSON
// 2 设置 Header Content-Type 字段
// 注意：不会处理错误，遇到错误则不设置
func (r *Request) WithJSONBody(rawBody interface{}) *Request {
	bts, err := json.Marshal(rawBody)
	if err != nil {
		return r
	}

	r.Body = bts
	r.Headers["Content-Type"] = "application/json"

	return r
}

// transObject2Queries 将对象转成 url.Values
// 注意：仅转换一层结构，key 为 对象属性的 json tag 表示（或默认）
func transObject2Queries(obj interface{}) (q url.Values, err error) {
	bts, err := json.Marshal(obj)
	if err != nil {
		return q, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(bts, &m)
	if err != nil {
		return q, err
	}

	q = make(map[string][]string)

	for k, v := range m {
		if v != nil {
			q[k] = []string{fmt.Sprintf("%v", v)}
		}
	}

	return
}
