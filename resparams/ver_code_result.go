package resparams

import "HelloApp_Api/constants"

type SmsVerCodeResult struct {
	constants.Result
	Data ResultSmsVerCode `json:"data"`
}

type ResultSmsVerCode struct {
	VerCode string `json:"ver_code"`
}
