## 目录结构
##### jvmgo jvm的go语言实现
##### cmdline 这个包用来解析java命令行参数
##### classfile 这个包把[]byte转化成ClassFile结构
##### classpath 这个包实现了class文件查找
##### native rt.jar里的本地方法实现
##### jvm jvm核心功能
##### instructions 指令集
##### rtda 运行时数据区（Runtime Data Area），线程、堆栈、栈帧等都在这个包里
##### class 类和对象（以及数组）等在这个包里

      核心JVM
            实现核心JVM其实是相对比较简单的，对照着JVM规范，把Thread、Frame、Operand Stack、Local Vars、Class、Object、Array、指令集等一一实现就可以了。下面以FrameStack和OperandStack为例，简单介绍一下：
```golang
     type Stack struct {
       maxSize uint
       size    uint
       _top    *Frame // stack is implemented as linked list
     }
     type OperandStack struct {
       size  uint
       slots []Any
     }
     //FrameStack是用链表（LinkedList）实现的，OperandStack内部其实用了Slice。
     指令集
            为了更好的代码可读性，每一个指令都实现成一了个struct，下面是iinc指令的完整代码：

     package instructions

     import "jvmgo/jvm/rtda"

     // Increment local variable by constant
     type iinc struct {
       index  uint
       _const int32
     }

     func (self *iinc) fetchOperands(decoder *InstructionDecoder) {
       self.index = uint(decoder.readUint8())
       self._const = int32(decoder.readInt8())
     }

     func (self *iinc) Execute(frame *rtda.Frame) {
       localVars := frame.LocalVars()
       val := localVars.GetInt(self.index)
       val += self._const
       localVars.SetInt(self.index, val)
     }
     //因为大部分指令都是需要操作OperandStack和／或LocalVars的，所以指令的Execute方法参数设计接收为*Frame类型的参数：
     type Instruction interface {
       fetchOperands(decoder *InstructionDecoder)
       Execute(frame *rtda.Frame)
     }
```
      类库和本地方法
            最初开始写jvm.go的时候，用的是OpenJDK的rt.jar。但是因为要经常查看rt.jar的Java代码，用IDE可以直接跳进Oracle JDK的rt.jar代码里。所以为了方便，后来就改为针对Oracle的rt.jar进行开发。rt.jar里有几千个本地方法，目前为止，只实现了不到100个。下面是本地方法的类型定义：

            type NativeMethod func(frame *rtda.Frame)
            垃圾回收
            Go本身就是垃圾回收语言，所以jvm.go没有单独实现垃圾回收机制。
      线程
            jvm.go把每个Java线程都映射为一个goroutine，下面是Thread.start0()本地方法的实现代码：
```golang
     // private native void start0();
     // ()V
     func start0(frame *rtda.Frame) {
       vars := frame.LocalVars()
       this := vars.GetThis()

       newThread := rtda.NewThread(this)
       runMethod := this.Class().GetInstanceMethod("run", "()V")
       newFrame := newThread.NewFrame(runMethod)
       newFrame.LocalVars().SetRef(0, this)
       newThread.PushFrame(newFrame)

       this.LockState()
       this.SetExtra(newThread)
       this.UnlockState()
       go interpreter.Loop(newThread)
     }
```

