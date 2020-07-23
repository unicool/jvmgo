package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type MarkerAttribute struct {
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
