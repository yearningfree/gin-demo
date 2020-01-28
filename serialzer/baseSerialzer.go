package serialzer

// Response 序列化基础响应
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// DataList
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 序列化数据列表
func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Data: DataList{
			Items: item,
			Total: total,
		},
	}
}
