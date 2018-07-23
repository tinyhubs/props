package props

import (
	"strconv"
	"sort"
)

//	Source 定义了配置数据来源.
//	一个数据来源必须提供一个 Find 方法用于支持从该来源读取配置数据
type Source interface {
	Find(key string) (string, bool)
}

//	Props 定义了配置数据的操作集合.
//	该接口提供了添加配置数据源，配置项查找，配置项读取，使用配置项扩展变量等多个函数.
type Props interface {
	//	Add 用于将配置数据源 s，以优先级 priority，添加到 Props 对象中去.
	Add(priority uint8, s Source) Source

	//	Find 函数尝试从多个配置数据源查找名字为 key 的配置项.
	//	如果能够找到返回找到的配置项的值和true；否则，第二个返回值为false。
	//	如果多个配置数据源都存在名字为 key 的配置项，那么以优先级最高的配置项为准。
	//	String、Int64、Uint64、Bool 等函数是对 Find 函数的简单封装。
	Find(key string) (string, bool)

	//	String 尝试将找到的配置项以 string 的形式返回，如果找不到就返回 def.
	String(key string, def string) string

	//	Int64 尝试将找到的配置项转换为 int64 的形式，如果找不到或者转换失败，就返回 def.
	Int64(key string, def int64) int64

	//	Uint64 尝试将找到的配置项转换为 uint64 的形式，如果找不到或者转换失败，就返回 def.
	Uint64(key string, def uint64) uint64

	//	Bool 尝试将找到的配置项转换为 bool 的形式，如果找不到或者转换失败，就返回 def.
	//	当前支持的数据映射方式为：
	// 	"1", "t", "T", "true", "TRUE", "True" 会被映射为 true；
	//	"0", "f", "F", "false", "FALSE", "False" 会被映射为 false。
	//	如果配置项的值为其他的值，作为转换失败处理。
	Bool(key string, def bool) bool

	//	将字符串 s 中的所有 ${key} 形式的可扩展变量，替换为本 Props 对象内部的配置项 key 的值，并最终返回替换后的结果.
	//	如果某个 key 的配置项在 Props 对象中不存在，将原样保留 ${key}；
	//	如果某个 key 的配置项的值里面还含有 ${xxx}，那么也会自动展开；
	//	如果变量存在循环引用现象，返回失败；
	Expand(s string) (string, error)
}

//	NewProps 用于创建一个新的 Props 对象.
func NewProps() Props {
	return &implProps{
		items: make([]*implSourceItem, 0, 5),
	}
}

type implSourceItem struct {
	priority int
	source   Source
	next     *implSourceItem
}

type implProps struct {
	items []*implSourceItem
}

func (p *implProps) Add(priority uint8, s Source) Source {
	newItem := &implSourceItem{
		priority: int(priority),
		source:   s,
		next:     nil,
	}

	for i := 0; i < len(p.items); i++ {
		if int(priority) == p.items[i].priority {
			newItem.next = p.items[i]
			p.items[i] = newItem
			return s
		}
	}

	p.items = append(p.items, newItem)
	sort.Slice(p.items, func(i, j int) bool {
		return p.items[i].priority < p.items[j].priority
	})

	return s
}

func (p implProps) Find(key string) (string, bool) {
	for i := len(p.items) - 1; i >= 0; i-- {
		val, ok := p.items[i].source.Find(key)
		if ok {
			return val, true
		}
	}

	return "", false
}

func (p implProps) String(key string, def string) string {
	val, ok := p.Find(key)
	if !ok {
		return def
	}

	return val
}

func (p implProps) Int64(key string, def int64) int64 {
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

func (p implProps) Uint64(key string, def uint64) uint64 {
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

func (p implProps) Bool(key string, def bool) bool {
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

func (p implProps) Expand(s string) (string, error) {
	return s, nil
}
