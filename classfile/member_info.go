package classfile

type MemberInfo struct{

	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

//读取字段表或方法表
func readMembers(reader *ClassReader,cp ConstantPool) []*MemberInfo{
	//获取fields_count 或者methods_count
	memberCount :=reader.readUint16()
	member :=make([]*MemberInfo,memberCount)

	for i:=range members{
		members[i]=readMember(reader,cp)
	}
}