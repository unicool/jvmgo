package extended

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type IFNULL struct {
	base.BrandInstruction
}

type IFNONNULL struct {
	base.BrandInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
