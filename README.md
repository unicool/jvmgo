# java virtual machine
------

## 字节码文件文件结构核心定义

### class file path
- 启动类路径(bootstrap classpath)：   -> jre/lib
- 扩展类路径(extension classpath)：   -> jre/lib/ext
- 用户类路径(user classpath)：        -> -classpath/-cp


### class file architecture
> 为了描述class文件格式，Java虚拟机规范定义了u1、u2和u4三种数据类型来表示1、2和4字节无符号整数，分别对应Go语言的uint8、uint16和uint32类型

```golang
ClassFile {
    u4                          magic; // 0xCAFEBABE
    u2                          minor_version;
    u2                          major_version;
    u2                          constant_pool_count;
    cp_info                     constant_pool[constant_pool_count-1];
    u2                          access_flags;
    u2                          this_class;
    u2                          super_class;
    u2                          interfaces_count;
    u2                          interfaces[interfaces_count];
    u2                          fields_count;
    field_info                  fields[fields_count];
    u2                          methods_count;
    method_info                 methods[methods_count];
    u2                          attributes_count;
    attribute_info              attributes[attributes_count];
}
```

#### constant_pool
> Java虚拟机规范一共定义了14种常量

```golang
cp_info {
    u1 tag;
    u1 info[];
}
```
```golang
tag_const ( // 区别于运行时常量池
	CONSTANT_Class              = 7     // u2(name_index)
	CONSTANT_Fieldref           = 9     // u2(class_index) + u2(name_and_type_index)
	CONSTANT_Methodref          = 10    // u2(class_index) + u2(name_and_type_index)
	CONSTANT_InterfaceMethodref = 11    // u2(class_index) + u2(name_and_type_index)
	CONSTANT_String             = 8     // u2(Utf8 index)
	CONSTANT_Integer            = 3     // u4
	CONSTANT_Float              = 4     // u4
	CONSTANT_Long               = 5     // u4(high_bytes) + u4(low_bytes)
	CONSTANT_Double             = 6     // u4(high_bytes) + u4(low_bytes)
	CONSTANT_NameAndType        = 12    // u2(name_index) + u2(descriptor_index)
	CONSTANT_Utf8               = 1     // u2(length)+bytes
	CONSTANT_MethodHandle       = 15    // u1(reference_kind) + u2(reference_index)
	CONSTANT_MethodType         = 16    // u2(descriptor_index)
	CONSTANT_InvokeDynamic      = 18    // u1(bootstrap_method_attr_index) + u2(name_and_type_index)
)
```

#### field_info & method_info
> 字段和方法的基本结构大致相同，差别仅在于属性表

```golang
member_info {
    u2                          access_flags;
    u2                          name_index;
    u2                          descriptor_index;
    u2                          attributes_count;
    attribute_info              attributes[attributes_count];
}
```
> eg:
> ![descriptor_index](.\doc\descriptors_eg.jpg)

#### attribute_info
> Java虚拟机规范预定义了23种属性，按照用途可以分为三组。第一组属性是实现Java虚拟机所必需的，共有5种；第二组属性是Java类库所必需的，共有12种；第三组属性主要提供给工具使用，共有6种。第三组属性是可选的，也就是说可以不出现在class文件中。

```golang
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
```

> ![attributes](.\doc\attributes_23.jpg)


#### access_flags
> 字段和方法的访问标志含义有所区别

```golang
const (
	ACC_PUBLIC       = 0x0001 // class field method
	ACC_PRIVATE      = 0x0002 // field method
	ACC_PROTECTED    = 0x0004 // field method
	ACC_STATIC       = 0x0008 // field method
	ACC_FINAL        = 0x0010 // class field method
	ACC_SUPER        = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 // method
	ACC_VOLATILE     = 0x0040 // field
	ACC_BRIDGE       = 0x0040 // method
	ACC_TRANSIENT    = 0x0080 // field
	ACC_VARARGS      = 0x0080 // method
	ACC_NATIVE       = 0x0100 // method
	ACC_INTERFACE    = 0x0200 // class
	ACC_ABSTRACT     = 0x0400 // class method
	ACC_STRICT       = 0x0800 // method
	ACC_SYNTHETIC    = 0x1000 // class field method
	ACC_ANNOTATION   = 0x2000 // class
	ACC_ENUM         = 0x4000 // class field
)
```


------
## JVM 内存结构

### run-time data area
> ![run-time data area](.\doc\run-time data area.jpg)

#### thread
> 线程私有的运行时数据区，用于辅助执行Java字节码
- pc寄存器（Program Counter）
- java虚拟机栈（JVM Stack）
    - Stack Frame
        - Local Variable
        - Operand Stack
        - head.Method
        - nextPC // the next instruction after the call

#### 指令集和解释器
> - 目前定义了205条指令：0x00 - 0x202, 0xFE, 0xFF
> - 指令分类：常量（constants）、加载（loads）、存储（stores）、操作数栈（stack）、数学（math）、转换（conversions）、比较（comparisons）、控制（control）、引用（references）、扩展（extended）、保留（reserved）
- new：创建类实例
    > 通过该uint16索引，从当前类的运行时常量池中找到一个类符号引用，解析这个类符号引用，拿到数据，然后创建对象，并把对象引用推入栈顶
- putstatic、getstatic：存取静态变量
- putfield、getfield：存取实例变量
- instanceof、checkcast：判断对象是否属于某种类型
- ldc：系统指令，把运行时常量池中的常量推到操作数栈顶

#### head
> 多线程共享的运行时数据区
##### method area
- 存放：1.从class文件获取的类信息；2.类变量
- 运行时常量池：字面量(literal)和符号引用(symbolic)，(数值不再使用index引用)
- class：类变量空间大小、静态变量\空间大小、etc
    > jvm并不限制类型不同但名字相同的字段

##### ClassLoader 
> 数组类和普通类有很大的不同，他的数据并不是来自class文件，而是由Java虚拟机在运行期间生成
1. 找到class文件并把数据读取到内存
2. 解析class文件，生成虚拟机可以使用的类数据，并放入方法区
3. 链接：验证()，准备：给类变量分配空间并赋予初始值

#### 方法调用和返回（Java）
> 调用分类1：静态方法/类方法、实例方法
> 实现分类2：抽象方法、Java方法、本地方法
- 指令
    - invokestatic：调用静态方法
    - invokespecial：调用无须动态绑定的实例方法，包括构造方法、私有方法和超类方法
    - invokeinterface：动态绑定，针对接口类型的引用
    - invokevirtual：动态绑定
    - invokedynamic：为了更好地支持动态类型语言
##### 基本过程
> 首先，方法调用指令需要n+1个操作数，其中第1个操作数是uint16索引，在字节码中紧跟在指令操作码的后面。
> 通过这个索引，可以从当前类的运行时常量池中找到一个方法符号引用，解析这个符号引用就可以得到一个方法。
> 注意，这个方法并不一定就是最终要调用的那个方法，所以可能还需要一个查找过程才能找到最终要调用的方法。
> 剩下的n个操作数是要传递给被调用方法的参数，从操作数栈中弹出。
>
> 如果要执行的是Java方法（而非本地方法），下一步是给这个方法创建一个新的帧，并把它推到Java虚拟机栈顶。
> 传递参数之后，新的方法就可以开始执行了。
>
- 方法调用和参数传递
> 在定位到需要调用的方法之后，Jvm要給这个方法创建一个新的帧(frame)并把它推入`jvm stack`顶，然后传递参数

##### 类初始化
> `<clinit>`: 类的初始化方法

#### array & string
- 数组类由JVM在运行时生成
    > superClass: "java/lang/Object",
    > interfaces: {
    > 	"java/lang/Cloneable",
    > 	"java/io/Serializable",
    > },
- 指令
    - newarray：操作数1-指令码后面的atype(8bits)；操作数2-操作数栈的count
    - anewarray：创建引用类型数组
    - multianewarray：创建多维数组
    - arraylength：弹出引用
    - <t>aload：按索引取数组元素值
    - <t>astore：按索引給数组赋值
    
- 在`class`中`string`是以`MUTF8`格式保存的，而在JVM运行时又是以`UTF16`格式保存的
> 






