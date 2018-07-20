package props

import (
	"os"
	"strings"
)

type SourceCLI struct {
	SourceCustom
}

func NewSourceCLI() *SourceCLI {
	return &SourceCLI{
		SourceCustom: SourceCustom{
			props: nil,
		},
	}
}

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
