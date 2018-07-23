package props

import (
	"os"
	"strings"
)

//	SourceCLI 用于将命令行输入参数作为配置数据源.
type SourceCLI struct {
	SourceCustom
}

//	NewSourceCLI 创建一个 SourceCLI 对象.
func NewSourceCLI() *SourceCLI {
	return &SourceCLI{
		SourceCustom: SourceCustom{
			props: nil,
		},
	}
}

//	Accept 将逐个扫描所有的命令行的输入参数，并将以 `prefix` 为前缀的参数作为配置数据源的配置项.
//	如果有多个前缀，那么可以调用 `Accept函数多次`。
func (s *SourceCLI) Accept(prefix string) {
	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], prefix) {
			key := ""
			val := ""
			str := os.Args[i][len(prefix):]
			pos := strings.IndexByte(str, '=')
			if pos > 0 {
				key = strings.TrimSpace(str[0:pos])
				val = strings.TrimSpace(str[pos+1:])
			} else if pos < 0 {
				key = strings.TrimSpace(str)
				val = ""
			}

			if "" == key {
				continue
			}

			s.props[key] = val
		}
	}
}
