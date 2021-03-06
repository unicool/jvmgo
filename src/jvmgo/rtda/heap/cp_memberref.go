package heap

import "jvmgo/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) copyMemberrefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
