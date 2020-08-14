package loads

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type DLOAD struct {
	base.Index8Instruction
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, uint(0))
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, uint(1))
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, uint(2))
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, uint(3))
}
