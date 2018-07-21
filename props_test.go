package props

import (
	"testing"
	"github.com/tinyhubs/et/expect"
)

func Test_Default(t *testing.T) {
	p := New()
	src0 := p.Add(0, NewSourceCustom()).(*SourceCustom)
	src0.Set("a", "1")
	src0.Set("b", "2")

	src1 := p.Add(1, NewSourceCustom()).(*SourceCustom)
	src1.Set("b", "sss")
	src1.Set("c", "4")

	expect.Equali(t, "没有更高优先级的配置", "1", p.String("a", ""))
	expect.Equali(t, "存在更高优先级的配置", "ss", p.String("b", ""))

}
