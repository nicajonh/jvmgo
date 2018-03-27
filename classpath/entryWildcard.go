package classpath

import (
	"os"
	"path/filepath"
	"string"
)


/*
	WildcardEntry实际上也是CompositeEntry
	所以不再定义新的类型了
	也不用定义readClass和String，newWildcardEntry返回的就是CompositeEntry
 */


 func newWildcardEntry(path string) CompositeEntry{
	 baseDir :=path[:len(path)-1]
	 compositeEntry :=[]Entry{}

	/*
		在walkFn中，根据后缀名选出JAR文件，并返回SkipDir跳过子目录
		通配符类路径不能递归匹配子目录下的JAR文件
	 */

	 walkFn :=func(path string,info os.FileInfo,err error) error{
		 if err!=nil{
			 return err
		 }
		 if info.IsDir() && path !=baseDir{
			 return filepath.SkipDir
		 }
		 if string.HasSuffix(path,".jar") || string.HasSuffix(path,".JAR"){
			 jarEntry :=newZipEntry(path)
			 compositeEntry=append(compositeEntry,jarEntry)
		 }
		 return nil
	 }
	 //调用filepath包的Walk函数，遍历baseDir创建ZipEntry
	 filepath.Walk(baseDir,walkFn)
	 return compositeEntry
 }
