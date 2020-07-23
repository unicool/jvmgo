package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type ALOAD struct {
	base.Index8Instruction
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(self.Index))
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, uint(0))
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, uint(1))
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, uint(2))
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, uint(3))
}
