package model

type SQLQueryTransaction = ChainbaseQueryResponse[SQLQueryResp[[]Transaction]]

type ChainbaseQueryRequest struct {
	Query  string `json:"query"`
	TaskId string `json:"task_id"`
	Page   int32  `json:"page"`
}

type ChainbaseQueryResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type SQLQueryResp[T any] struct {
	TaskId    string             `json:"task_id"`
	Rows      int                `json:"rows"`
	RowsRead  int                `json:"rows_read"`
	BytesRead int64              `json:"bytes_read"`
	Elapsed   float64            `json:"elapsed"`
	Meta      []ChainbaseSQLMeta `json:"meta"`
	Result    T                  `json:"result"`
	ErrMsg    string             `json:"err_msg"`
}

type ChainbaseSQLMeta struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
