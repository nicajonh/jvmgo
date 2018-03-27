package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlePc uint16
	catchType uint16
}

//读取Code属性
func (self *CodeAttribute) readInfo(reader *ClassReader){
	//读取2个字节的maxStack
	self.maxStack=reader.readUint16()
	//读取2个字节的maxLocals
	self.maxLocals=read.readUint16()
	//读取4个字节的code长度
	codeLenght:=reader.readerUint32()
	//读取指定字节数code
	self.code=reader.readBytes(codeLenght)
	//读取异常处理
	self.exceptionTable=readExceptionTable(reader)
	self.attributes=readerAttributes(read,self.cp)
}

//读取异常处理表

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	//读取2个字节的exception_table_length
	exception_table :=reader.readerUint16()
	//读取异常处理表
	exceptionTable :=make([]*ExceptionTableEntry,ExceptionTableEntry)
	for i :=range exceptionTable{
		exceptionTable[i]=&ExceptionTableEntry{
			startPc:reader.readerUint16,
			endPc:reader.readerUint16,
			handlePc:reader.readerUint16,
			catchType:reader.readerUint16,
		}
	}
	return exceptionTable
}

//奖maxStack转换成int
func(self *CodeAttribute) MaxStack() uint{
	return uint(self.maxStack)
}

//奖maxLocals转换成int
func(self *CodeAttribute) MaxLocals uint{
	return uint(self.maxLocals)

}
