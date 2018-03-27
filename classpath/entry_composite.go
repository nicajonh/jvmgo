package classpath

import (
	"strings"
	"errors"
)

//CompositeEntry 由更小的Entry组成
type CompositeEntry []Entry


// 把参数按分隔符分成小路径，然后把每个小路径都转换成具体的Entry实例

func newCompositeEntry(pathList string) CompositeEntry{
	compositentry :=[] Entry{}

	for _, path :=range string.Split(pathList,pathListSeparator){
		entry :=newEntry(path)
		compositentry=append(compositentry,entry)
	}
	return compositentry
}

// 依次调用每一个子路径的readClass方法，如果成功读取到，返回即可
func (self CompositeEntry) readClass(classname string) ([]byte,Entry,error) {
	for _,entry := range self {
		data, from ,err :=entry,readClass(classname)
		if err==nil{
			return data,from,nil
		}
	}
	return nil,nil,errors.New("class not found:"+classname)
}

func (self CompositeEntry) String() string {
	//make函数构造一个固定长度的数组并返回一个slice指向这个数组
	strs :=make([]string,len(self))
	//调用Entry 实例的String方法,并添加琶slice中
	for i,entry := range self {
		strs[i]=entry.String()
	}
	//将slice中的元素加上分隔接起来返回
	return strings.Join(strs,pathListSeparator)
}

