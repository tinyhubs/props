package props

//	SourceCustom 表示用户自定义的配置数据源.
type SourceCustom struct {
	props map[string]string
}

//	NewSourceCustom 创建一个 SourceCustom 的数据源对象.
func NewSourceCustom() *SourceCustom {
	return &SourceCustom{
		props: make(map[string]string),
	}
}

//	Find 为 Source 的接口实现，具体功能参见 `Source`.
func (s SourceCustom) Find(key string) (string, bool) {
	val, ok := s.props[key]
	if !ok {
		return "", false
	}

	return val, true
}

//	Set 用于将 key 和 value 放入 SourceCustom 对象中，方便后续通过 Props 的接口获取.
func (s *SourceCustom) Set(key string, val string) {
	s.props[key] = val
}
