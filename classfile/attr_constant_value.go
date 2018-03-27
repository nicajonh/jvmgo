package classfile


type ConstantValueAttribute struct{
	constantValueIndex uint16
}

//读取ConstantValue_attribute
func (self *ConstantValueAttribute) readInfo(reader *ClassReader){
	//读取2个字节的constantValueIndex
	self.constantValueIndex=reader.readUint16
}

//读取constantValueIndex
func (self *ConstantValueAttribute) ConstantValueIndex() uint16{
	return self.constantValueIndex
}
