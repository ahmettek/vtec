package models

type SetRequestModel struct {
	Key   string
	Value string
}

type SetResponseModel struct {
	Success bool
}

type GetResponseModel struct {
	Value string
}

type FlushResponseModel struct {
	Success bool
}
