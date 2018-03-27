package classfile

import (
	"fmt"
	"unicode/utf16"
)

type ConstantUtf8Info struct{
	str string
}

//读取CONSTANT_Utf8_info

func (self *ConstantUtf8Info) readInfo(read *ClassReader){
	//读取length(2个节),并转换成unit32
	length :=uint32(reader.readUint16())
	//读取指定长度的字节
	bytes :=reader.readBytes(length)
	//将字节转换成utf-8
	self.str=decondeMUTF8(bytes)
}

//获取string 常量值 

func (self *ConstantUtf8Info) Str() string{
	return self.str
}

//该方法是书的作者参照