package classfile

import "fmt"

type ConstantPool []ConstantInfo

//读取常量池
func readConstantPool(reader *ClassReader) ConstantPool{
	//常量池数量(2个字节)
	cpCount:=int(reader.readUint16())
	//生成长度为cpCount,类型是Constant的slice就是常量池
	cp := make([]ConstantInfo,cpCount)
	//注意:常量池的数据从1开始
	for i:=1;i<cpCount;i++ {
		cp[i]=readConstantInfo(reader,cp)
		switch cp[i].(type){
			case 	*ConstantLongInfo, *ConstantDoubleInfo:
				i++
		}
	}
	return cp
}

// 通过索引读取常量结构体
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
		/*
		为什么不是index - 1？
		因为常量池的索引是从1开始的！
	 */
	 if cpInfo:=self[index];cpInfo !=nil{
		 return cpInfo
	 }

	 panic(fmt.Errorf("Invalid constant pool index: %v!", index))