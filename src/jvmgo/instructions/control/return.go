package control

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction // return void
}
type ARETURN struct {
	base.NoOperandsInstruction // return reference
}
type DRETURN struct {
	base.NoOperandsInstruction // return double
}
type FRETURN struct {
	base.NoOperandsInstruction // return float
}
type IRETURN struct {
	base.NoOperandsInstruction // return int
}
type LRETURN struct {
	base.NoOperandsInstruction // return long
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFramg := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokerFramg.OperandStack().PushRef(retVal)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFramg := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFramg.OperandStack().PushDouble(retVal)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFramg := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFramg.OperandStack().PushFloat(retVal)
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFramg := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFramg.OperandStack().PushInt(retVal)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFramg := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFramg.OperandStack().PushLong(retVal)
}
