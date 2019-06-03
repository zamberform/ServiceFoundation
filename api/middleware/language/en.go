package language

import "api/pkg/codes"

var MsgEnFlags = map[int]string{
	codes.SUCCESS:        "ok",
	codes.ERROR:          "fail",
	codes.INVALID_PARAMS: "error params",
}
