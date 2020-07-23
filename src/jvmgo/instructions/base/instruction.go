package base

import "jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// 没有操作数的指令
type NoOperandsInstruction struct {
	// empty
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 跳转指令
type BrandInstruction struct {
	Offset int // 跳转偏移量
}

func (self *BrandInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

// 存储和加载类指令需要根据索引存取局部变量表
type Index8Instruction struct {
	Index uint // 局部变量表索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// 用于访问常量池
type Index16Instruction struct {
	Index uint // 常量池索引
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
