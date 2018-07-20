package props

type SourceCustom struct {
	props map[string]string
}

func NewSourceCustom() *SourceCustom {
	return &SourceCustom{
		props: make(map[string]string),
	}
}

func (s SourceCustom) Find(key string) (string, bool) {
	val, ok := s.props[key]
	if !ok {
		return "", false
	}

	return val, true
}

func (s *SourceCustom) Set(key string, val string) {
	s.props[key] = val
}
