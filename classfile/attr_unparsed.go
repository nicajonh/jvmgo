package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

type UnparseAttribute struct{
	name string
	length uint32
	info []byte
}

func (self *UnparseAttribute) readInfo(read *ClassReader){
	self.info=reader.readBytes(self.lenggh)
}

func (self *UnparseAttribute) Info() []byte{
	return self.info
}