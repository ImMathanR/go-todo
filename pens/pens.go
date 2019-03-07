package pens

import (
	"github.com/sankalpjonn/pen"
)

type Pens struct {
	m map[string]*pen.Pen
}

func New() *Pens {
	return &Pens{
		m: map[string]*pen.Pen{},
	}
}

func (self *Pens) addPen(name string, path string) {
	self.m[name] = pen.New(name, path)
}

func (self *Pens) Close() {
	for _, p := range self.m {
		p.Lid()
	}
}

func (self *Pens) Write(name string, data string) {
	if self.m[name] == nil {
		self.addPen(name, "/tmp/"+name+".log.%Y%m%d%H")
	}
	self.m[name].Write(data)
}
