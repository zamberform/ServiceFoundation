package language

const (
	EN = iota
	ZH
	JP
)

type Language int

var langMap map[int]string

var defaultStr = "not define"

func init() {
	langMap = MsgLang[EN]
}

func Config(lng Language) {
	langMap = MsgLang[int(lng)]
	switch lng {
	case EN:
		defaultStr = "not define"
	case ZH:
		defaultStr = "未定义"
	case JP:
		defaultStr = "定義していません"
	}
}

func GetMsg(code int) string {
	msg, ok := langMap[code]
	if ok {
		return msg
	}

	return defaultStr
}
