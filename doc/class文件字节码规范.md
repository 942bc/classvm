**1.****总体格式**

| **Class File format** |                                        |                    |
| --------------------- | -------------------------------------- | ------------------ |
| **type**              | **descriptor**                         | **remark**         |
| u4                    | magic                                  | 0xCAFEBABE         |
| u2                    | minor_version                          |                    |
| u2                    | major_version                          |                    |
| u2                    | constant_pool_count                    |                    |
| cp_info               | constant_pool[cosntant_pool_count – 1] | index 0 is invalid |
| u2                    | access_flags                           |                    |
| u2                    | this_class                             |                    |
| u2                    | super_class                            |                    |
| u2                    | interfaces_count                       |                    |
| u2                    | interfaces[interfaces_count]           |                    |
| u2                    | fields_count                           |                    |
| field_info            | fields[fields_count]                   |                    |
| u2                    | methods_count                          |                    |
| method_info           | methods[methods_count]                 |                    |
| u2                    | attributes_count                       |                    |
| attribute_info        | attributes[attributes_count]           |                    |

 

**2.**   **格式详解**

**2.1** **magic**

magic被称为“魔数”，用来标识.class文件的开头。所有合法的.class字节码都应该是该数开头，占4个字节。

**2.2** **major_version.minor_version**

major_version.minor_version合在一起形成当前.class文件的版本号，该版本号一般由编译器产生，并且由sun定义。如59.0。它们一起占4个字节。

**2.3** **constant_pool**

在Java字节码中，有一个常量池，**用来存放不同类型的常量**。由于Java设计的目的之一就是字节码需要经网络传输的，因而字节码需要比较紧凑，以减少网络传输的流量和时间。常量池的存在则可以让一些相同类型的值通过索引的方式从常量池中找到，而不是在不同地方有不同拷贝，缩减了字节码的大小。

每个常量池中的项是通过**cp_info**的类型来表示的，它的格式如下：

| **cp_info format** |                |            |
| ------------------ | -------------- | ---------- |
| **type**           | **descriptor** | **remark** |
| u1                 | tag            |            |
| u1                 | info[]         |            |

这里tag用来表示当前常量池不同类型的项。info中存放常量池项中存放的数据。

tag中表示的数据类型：

CONSTANT_Class_info                 （7）、

CONSTANT_Integer_info               （3）、

CONSTANT_Long_info                  （5）、

CONSTANT_Float_info                 （4）、

CONSTANT_Double_info               （6）、

CONSTANT_String_info                 （8）、

CONSTANT_Fieldref_info               （9）、

CONSTANT_Methodref_info            （10）、

CONSTANT_InterfaceMethodref_info   （11）、

CONSTANT_NameAndType_info        （12）、

CONSTANT_Utf8_info                  （1）、

注：在Java字节码中，所有boolean、byte、char、short类型都是用int类型存放，因而在常量池中没有和它们对应的项。

**2.3.1**  **CONSTANT_Class_info**

用于记录类或接口名（used to represent a class or an interface）

| **CONSTANT_Class_info format** |                |                                                              |
| ------------------------------ | -------------- | ------------------------------------------------------------ |
| **type**                       | **descriptor** | **remark**                                                   |
| u1                             | tag            | CONSTANT_Class (7)                                           |
| u2                             | name_index     | constant_pool中的索引，CONSTANT_Utf8_info类型。表示类或接口名。 |

注：在Java字节码中，类和接口名不同于源码中的名字，详见**附件****A.**

 

**2.3.2**  **CONSTANT_Integer_info**

用于记录int类型的常量值（represent 4-byte numeric (int) constants:）

| **CONSTANT_Integer_info** |                |                      |
| ------------------------- | -------------- | -------------------- |
| **type**                  | **descriptor** | **remark**           |
| u1                        | tag            | CONSTANT_Integer (3) |
| u4                        | bytes          | 整型常量值           |

 

**2.3.3**  **CONSTANT_Long_info**

用于记录long类型的常量值（represent 8-byte numeric (long) constants:）

| **CONSTANT_Long_info** |                |                   |
| ---------------------- | -------------- | ----------------- |
| **type**               | **descriptor** | **remark**        |
| u1                     | tag            | CONSTANT_Long (5) |
| u4                     | high_bytes     | 长整型的高四位值  |
| u4                     | low_bytes      | 长整型的低四位值  |

 

**2.3.4**  **CONSTANT_Float_info**

用于记录float类型的常量值（represent 4-byte numeric (float) constants:）

| **CONSTANT_Float_info** |                |                    |
| ----------------------- | -------------- | ------------------ |
| **type**                | **descriptor** | **remark**         |
| u1                      | tag            | CONSTANT_Float(4)  |
| u4                      | bytes          | 单精度浮点型常量值 |

几个特殊值：0x7f800000 => Float.POSITIVE_INFINITY、0xff800000 => Float.NEGATIVE_INFINITY、

​         0x7f800001 to 0x7fffffff => Float.NaN、0xff800001 to 0xffffffff => Float.NaN

 

**2.3.5**  **CONSTANT_Double_info**

用于记录double类型的常量值（represent 8-byte numeric (double) constants:）

| **CONSTANT_Double_info** |                |                      |
| ------------------------ | -------------- | -------------------- |
| **type**                 | **descriptor** | **remark**           |
| u1                       | tag            | CONSTANT_Double(6)   |
| u4                       | high_bytes     | 双精度浮点的高四位值 |
| u4                       | low_bytes      | 双精度浮点的低四位值 |

几个特殊值：0x7ff0000000000000L => Double.POSITIVE_INFINITY、

0xfff0000000000000L => Double.NEGATIVE_INFINITY

0x7ff0000000000001L to 0x7fffffffffffffffL => Double.NaN 、

0xfff0000000000001L to 0xffffffffffffffffL => Double.NaN

 

**2.3.6**  **CONSTANT_String_info**

用于记录常量字符串的值（represent constant objects of the type String:）

| **CONSTANT_String_info** |                |                                                              |
| ------------------------ | -------------- | ------------------------------------------------------------ |
| **type**                 | **descriptor** | **remark**                                                   |
| u1                       | tag            | CONSTANT_String(8)                                           |
| u2                       | string_index   | constant_pool中的索引，CONSTANT_Utf8_info类型。表示String类型值。 |

 

**2.3.7**  **CONSTANT_Fieldref_info**

用于记录字段信息（包括类或接口中定义的字段以及代码中使用到的字段）。

| **CONSTANT_Fieldref_info** |                     |                                                              |
| -------------------------- | ------------------- | ------------------------------------------------------------ |
| **type**                   | **descriptor**      | **remark**                                                   |
| u1                         | tag                 | CONSTANT_Fieldref(9)                                         |
| u2                         | class_index         | constant_pool中的索引，CONSTANT_Class_info类型。记录定义该字段的类或接口。 |
| u2                         | name_and_type_index | constant_pool中的索引，CONSTANT_NameAndType_info类型。指定类或接口中的字段名（name）和字段描述符（descriptor）。 |

 

**2.3.8**  **CONSTANT_Methodref_info**

用于记录方法信息（包括类中定义的方法以及代码中使用到的方法）。

| **CONSTANT_Methodref_info** |                     |                                                              |
| --------------------------- | ------------------- | ------------------------------------------------------------ |
| **type**                    | **descriptor**      | **remark**                                                   |
| u1                          | tag                 | CONSTANT_Methodref(10)                                       |
| u2                          | class_index         | constant_pool中的索引，CONSTANT_Class_info类型。记录定义该方法的类。 |
| u2                          | name_and_type_index | constant_pool中的索引，CONSTANT_NameAndType_info类型。指定类中扽方法名（name）和方法描述符（descriptor）。 |

 

**2.3.9**  **CONSTANT_InterfaceMethodref_info**

用于记录接口中的方法信息（包括接口中定义的方法以及代码中使用到的方法）。

| **CONSTANT_InterfaceMethodref_info** |                     |                                                              |
| ------------------------------------ | ------------------- | ------------------------------------------------------------ |
| **type**                             | **descriptor**      | **remark**                                                   |
| u1                                   | tag                 | CONSTANT_InterfaceMethodref(11)                              |
| u2                                   | class_index         | constant_pool中的索引，CONSTANT_Class_info类型。记录定义该方法的接口。 |
| u2                                   | name_and_type_index | constant_pool中的索引，CONSTANT_NameAndType_info类型。指定接口中的方法名（name）和方法描述符（descriptor）。 |

 

**2.3.10**  **CONSTANT_NameAndType_info**

记录方法或字段的名称(name)和描述符(descriptor)（represent a field or method, without indicating which class or interface type it belongs to:）。

| **CONSTANT_NameAndType_info** |                  |                                                              |
| ----------------------------- | ---------------- | ------------------------------------------------------------ |
| **type**                      | **descriptor**   | **remark**                                                   |
| u1                            | tag              | CONSTANT_NameAndType (12)                                    |
| u2                            | name_index       | constant_pool中的索引，CONSTANT_Utf8_info类型。指定字段或方法的名称。 |
| u2                            | descriptor_index | constant_pool中的索引，CONSTANT_utf8_info类型。指定字段或方法的描述符（**见附录****C**） |

 

**2.3.11**  **CONSTANT_Utf8_info**

记录字符串的值（represent constant string values. String content is encoded in *modified UTF-8*.）

modifie

d UTF-8 refer to :

[http://download.ora](http://download.oracle.com/javase/1.4.2/docs/api/java/io/DataInputStream.html)

[cle.com/javase/1.4.2/docs/api/java/io/DataInputStream.html](http://download.oracle.com/javase/1.4.2/docs/api/java/io/DataInputStream.html)

| **CONSTANT_Utf8_info** |                |                                                              |
| ---------------------- | -------------- | ------------------------------------------------------------ |
| **type**               | **descriptor** | **remark**                                                   |
| u1                     | tag            | CONSTANT_Utf8 (1)                                            |
| u2                     | length         | bytes所代表的字符串的长度                                    |
| u1                     | bytes[length]  | 字符串的byte数据，可以通过DataInputStream中的readUtf()方法（实例方法或静态方法读取该二进制的字符串的值。） |

 

**2.4** **access_flags**

指定类或接口的访问权限。

| **类或接口的访问权限** |           |                                                              |
| ---------------------- | --------- | ------------------------------------------------------------ |
| **Flag Name**          | **Value** | **Remarks**                                                  |
| ACC_PUBLIC             | 0x0001    | pubilc，包外可访问。                                         |
| ACC_FINAL              | 0x0010    | final，不能有子类。                                          |
| ACC_SUPER              | 0x0020    | 用于兼容早期编译器，新编译器都设置该标记，以在使用 *invokespecial*指令时对子类方法做特定处理。 |
| ACC_INTERFACE          | 0x0200    | 接口，同时需要设置：ACC_ABSTRACT。不可同时设置：ACC_FINAL、ACC_SUPER、ACC_ENUM |
| ACC_ABSTRACT           | 0x0400    | 抽象类，无法实例化。不可和ACC_FINAL同时设置。                |
| ACC_SYNTHETIC          | 0x1000    | synthetic，由编译器产生，不存在于源代码中。                  |
| ACC_ANNOTATION         | 0x2000    | 注解类型（annotation），需同时设置：ACC_INTERFACE、ACC_ABSTRACT |
| ACC_ENUM               | 0x4000    | 枚举类型                                                     |

 

**2.5** **this_class**

this_class是指向constant pool的索引值，该值必须是CONSTANT_Class_info类型，指定当前字节码定义的类或接口。

 

**2.6** **super_class**

super_class是指向constant pool的索引值，该值必须是CONSTANT_Class_info类型，指定当前字节码定义的类或接口的直接父类。只有Object类才没有直接父类，此时该索引值为0。并且父类不能是final类型。接口的父类都是Object类。

 

**2.7** **interfaces**

interfaces数组记录所有当前类或接口直接实现的接口。interfaces数组中的每项值都是一个指向constant pool的索引值，这些值必须是CONSTANT_Class_info类型。数组中接口的顺序和源代码中接口定义的顺序相同。

 

**2.8** **fields**

fields数组记录了类或接口中的所有字段，包括实例字段和静态字段，但不包含父类或父接口中定义的字段。fields数组中每项都是field_info类型值，它描述了字段的详细信息，如名称、描述符、字段中的attribute等。

| **field_info** |                              |                                                              |
| -------------- | ---------------------------- | ------------------------------------------------------------ |
| **type**       | **descriptor**               | **remark**                                                   |
| u2             | access_flags                 | 记录字段的访问权限。见2.8.1                                  |
| u2             | name_index                   | constant_pool中的索引，CONSTANT_Utf8_info类型。指定字段的名称。 |
| u2             | descriptor_index             | constant_pool中的索引，CONSTANT_Utf8_info类型，指定字段的描述符（见附录C）。 |
| u2             | attributes_count             | attributes包含的项目数。                                     |
| attribute_info | attributes[attributes_count] | 字段中包含的Attribute集合。见2.8.2-2.8.7                     |

**注：**fields中的项目和CONSTANT_Fieldref_info中的项目部分信息是相同的，他们主要的区别是CONSTANT_Fieldref_info中的项目不仅包含了类或接口中定义的字段，还包括在字节码中使用到的字段信息。不过这里很奇怪，为什么field_info结构中不把name_index和descriptor_index合并成fieldref_index，这样的class文件不是更加紧凑吗？？不知道这是sun因为某些原因故意这样设计还是这是他们的失误？？

 

**2.8.1**  **字段访问权限**

| **字段的访问权限** |           |                                                              |
| ------------------ | --------- | ------------------------------------------------------------ |
| **Flag Name**      | **Value** | **Remarks**                                                  |
| ACC_PUBLIC         | 0x0001    | pubilc，包外可访问。                                         |
| ACC_PRIVATE        | 0x0002    | private，只可在类内访问。                                    |
| ACC_PROTECTED      | 0x0004    | protected，类内和子类中可访问。                              |
| ACC_STATIC         | 0x0008    | static，静态。                                               |
| ACC_FINAL          | 0x0010    | final，常量。                                                |
| ACC_VOILATIE       | 0x0040    | volatile，直接读写内存，不可被缓存。不可和ACC_FINAL一起使用。 |
| ACC_TRANSIENT      | 0x0080    | transient，在序列化中被忽略的字段。                          |
| ACC_SYNTHETIC      | 0x1000    | synthetic，由编译器产生，不存在于源代码中。                  |
| ACC_ENUM           | 0x4000    | enum，枚举类型字段                                           |

**注：**接口中的字段必须同时设置：ACC_PUBLIC、ACC_STATIC、ACC_FINAL

 

**2.8.2**  **ConstantValue Attribute** **（****JVM****识别）**

| **ConstantValue Attribute** |                      |                                                              |
| --------------------------- | -------------------- | ------------------------------------------------------------ |
| **type**                    | **descriptor**       | **remark**                                                   |
| u2                          | attribute_name_index | constant_pool中的索引，CONSTANT_Utf8_info类型。指定Attribute的名称(“ConstantValue”)。 |
| u4                          | attribute_length     | 该Attribute内容的字节长度（固定值：2）                       |
| u2                          | constant_value_index | constant_pool中的索引，CONSTANT_Integer_info（int，boolean，char、short、byte）、CONSTANT_Float_info（float）、Constant_Double_info（double）、CONSTANT_Long_info（long）CONSTANT_String_info（String）类型 |

每个常量字段（final，静态常量或实例常量）都包含有且仅有一个ConstantValue Attribute。ConstantValue Attribute结构用于存储一个字段的常量值。

 

对一个**静态常量字段**，该常量值会在类或接口被初始化之前，由JVM负责赋给他们，即它在任何静态字段之前被赋值。

 

对一个**非静态常量字段**，该值会被虚拟机忽略，它的赋值由生成的实例初始化函数（<init>）实现。如类：

**class** A {

  **public** **static** **final** **int** *fa* = 10;

  **public** **final** **int** fa2 = 30;

  **private** **static** **int** *sa* = 20;

  **static** {

​    *sa* = 30;

  }

}

生成的字节码如下：

// Compiled from Test.java (version 1.6 : 50.0, super bit)

class org.levin.insidejvm.miscs.staticinit.A {

 ***public static final int fa = 10;\***

 **public final int fa2 = 30;**

 private static int sa;

 **static {};**

   0 bipush 20

   2 putstatic org.levin.insidejvm.miscs.staticinit.A.sa : int [16]

   5 bipush 30

   7 putstatic org.levin.insidejvm.miscs.staticinit.A.sa : int [16]

  10 return

 **public A();**

   0 aload_0 [this]

   1 invokespecial java.lang.Object() [21]

   4 aload_0 [this]

   **5 bipush 30**

   **7 putfield org.levin.insidejvm.miscs.staticinit.A.fa2 : int [23]**

10 return

 

**2.8.3**  **Synthetic Attribute**

**参考2.11.1**

**2.8.4**  **Signature Attribute**

**参考2.11.2**

**2.8.5**  **Deprecated Attribute**

**参考2.11.3**

**2.8.6**  **RuntimeVisibleAnnotations Attribute**

**参考2.11.4**

**2.8.7**  **RuntimeInvisibleAnnotations Attribute**

**参考2.11.5**