package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/

type ConstantMethodHandleInfo struct{
	referenceKind uint8
	referenceIndex uint8
}

//读取CONSTANT_MethodHandle_info

func(self *ConstantMethodHandleInfo) readInfo(reader *ClassReader){
	self.referenceKind=reader.readUint8()
	self.referenceIndex=reader.readerUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

// 读取CONSTANT_MethodType_info
func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}


CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}


// 读取CONSTANT_InvokeDynamic_info
func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}