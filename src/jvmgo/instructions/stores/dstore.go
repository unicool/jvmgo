package stores

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type DSTORE struct {
	base.Index8Instruction
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (self *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(self.Index))
}

func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(0))
}

func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(1))
}

func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(2))
}

func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(3))
}
