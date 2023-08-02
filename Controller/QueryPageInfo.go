package Controller

import (
	"GINWeb_project/Service"
	"strconv"
)

type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func QueryPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	pageInfo, err := Service.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}
}

