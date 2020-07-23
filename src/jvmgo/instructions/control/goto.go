package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type GOTO struct {
	base.BrandInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
