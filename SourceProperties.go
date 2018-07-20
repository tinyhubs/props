package props

import (
	"io"
	"bufio"
	"bytes"
	"unicode"
)

type SourceProperties struct {
	cache map[string]string
}

func NewSourceProperties() *SourceProperties {
	return &SourceProperties{
		cache: make(map[string]string),
	}
}

func (s *SourceProperties) Load(reader io.Reader) error {
	//  创建一个扫描器
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		//  逐行读取
		line := scanner.Bytes()

		//  遇到空行
		if 0 == len(line) {
			continue
		}

		//  找到第一个非空白字符
		pos := bytes.IndexFunc(line, func(r rune) bool {
			return !unicode.IsSpace(r)
		})

		//  遇到空白行
		if -1 == pos {
			continue
		}

		//  遇到注释行
		if '#' == line[pos] {
			continue
		}

		if '!' == line[pos] {
			continue
		}

		//  找到第一个等号的位置
		end := bytes.IndexFunc(line[pos+1:], func(r rune) bool {
			return ('=' == r) || (':' == r)
		})

		//  没有=，说明该配置项只有key
		key := ""
		value := ""
		if -1 == end {
			key = string(bytes.TrimRightFunc(line[pos:], func(r rune) bool {
				return unicode.IsSpace(r)
			}))
		} else {
			key = string(bytes.TrimRightFunc(line[pos:pos+1+end], func(r rune) bool {
				return unicode.IsSpace(r)
			}))

			value = string(bytes.TrimSpace(line[pos+1+end+1:]))
		}

		s.cache[key] = value
	}

	if err := scanner.Err(); nil != err {
		return err
	}

	return nil
}

func (s SourceProperties) Find(key string) (string, bool) {
	val, ok := s.cache[key]
	if !ok {
		return "", false
	}

	return val, true
}
