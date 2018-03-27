package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/

type ExceptionsAttribute struct{
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo (read *ClassReader){
	self.exceptionIndexTable=read.readUint16
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16{
	return self.exceptionIndexTable
}

