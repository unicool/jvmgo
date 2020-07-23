package constants

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	// nothing to do
}
