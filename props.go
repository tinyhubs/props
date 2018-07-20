package props

import (
	"strconv"
	"sort"
)

type Source interface {
	Find(key string) (string, bool)
}

type Properties interface {
	Add(priority uint8, s Source)
	Find(key string) (string, bool)
	String(key string, def string) string
	Int64(key string, def int64) int64
	Uint64(key string, def uint64) uint64
	Bool(key string, def bool) bool
	Expand(s string) string
}

func NewProperties() Properties {
	return &implProperties{
		items: make([]*implSourceItem, 0, 5),
	}
}

type implSourceItem struct {
	priority int
	source   Source
	next     *implSourceItem
}

type implProperties struct {
	items []*implSourceItem
}

func (p *implProperties) Add(priority uint8, s Source) {
	newItem := &implSourceItem{
		priority: int(priority),
		source:   s,
		next:     nil,
	}

	for i := 0; i < len(p.items); i++ {
		if int(priority) == p.items[i].priority {
			newItem.next = p.items[i]
			p.items[i] = newItem
			return
		}
	}

	p.items = append(p.items, newItem)
	sort.Slice(p.items, func(i, j int) bool {
		return p.items[i].priority < p.items[j].priority
	})
}

func (p implProperties) Find(key string) (string, bool) {
	for i := len(p.items) - 1; i >= 0; i-- {
		val, ok := p.items[i].source.Find(key)
		if ok {
			return val, true
		}
	}

	return "", false
}

func (p implProperties) String(key string, def string) string {
	val, ok := p.Find(key)
	if !ok {
		return def
	}

	return val
}

func (p implProperties) Int64(key string, def int64) int64 {
	val, ok := p.Find(key)
	if !ok {
		return def
	}

	i, err := strconv.ParseInt(val, 0, 64)
	if nil != err {
		return def
	}

	return i
}

func (p implProperties) Uint64(key string, def uint64) uint64 {
	val, ok := p.Find(key)
	if !ok {
		return def
	}

	i, err := strconv.ParseUint(val, 0, 64)
	if nil != err {
		return def
	}

	return i
}

func (p implProperties) Bool(key string, def bool) bool {
	val, ok := p.Find(key)
	if !ok {
		return def
	}

	i, err := strconv.ParseBool(val)
	if nil != err {
		return def
	}

	return i
}

func (p implProperties) Expand(s string) string {
	//TODO 太困了，先睡觉
	return s
}
