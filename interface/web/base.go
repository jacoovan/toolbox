package web

type commonFormPage struct {
	Page     uint `form:"page" json:"page"`
	PageSize uint `form:"pageSize" json:"pageSize"`
}

type commonResp struct {
	Code    uint        `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newCommonResp(code uint, success bool, message string, data interface{}) commonResp {
	resp := commonResp{
		Code:    code,
		Success: success,
		Message: message,
		Data:    make(map[string]interface{}),
	}
	if data != nil {
		resp.Data = data
	}
	return resp
}
