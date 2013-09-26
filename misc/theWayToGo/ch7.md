# array and slices
=====

[]表示数组类型，跟所有语言一样。为了理解slice 首先需要理解数组。数组在go中不常见(不太灵活) 但slice是随处可见（更强大,方便）!

## 声明与初始化
----

###  概念：
----
同类型 不同数据容量的数组是不同的类型！！数组的类型有其元素类型+长度决定！

如果想在数组中放不同的类型那么可以使用空接口，但后续操作需要做类型断言！

数组中的元素可以用索引访问。数组是0基的，最大长度是2Gb。
定义形式： 	var identifier [len]type
当定义一个数组时，每个元素都会被初始化为其类型的0值。
修改一个数组：  arr[i] = value
数组的最后一个元素 ： arr[len(arr)-1]
数组越界访问 有可能在编译期发现 或者在运行期触发panic。

使用for 结构来遍历数组。
惯用法:

~~~
	for i:=0 ; i <len(arr); i++{
		 arr[i] = ..
	}
// 惯用法二
	for i:= range arr{
		...
	}	
~~~


Go中的数组是值类型的（不想c/C++中的指向第一个元素） 所以可以用new() 创建：
var arr1 = new([5]int)

数组变量赋值 导致数组被复制。

所有复杂类型都有字面创建的形式：{} 在大括号里面初始化每个元素！

### 多维数组
例： [3][5]int
	[2][2][2]float64
	内部的数组总是相同长度。go的多维数组是矩形式的（唯有slice构成的数组列外）
### 传递数组给方法
----
数组传递可能消耗大量的内存。 有两种解决方案：
	1- 传递指向数组的指针
	2- 使用数组的slice

## slice 
----------
### 概念
指向连续数组段的引用。这个段可以是整个数组，或者是一个窗口段。slice 提供了一个动态窗口来查看底层的数组：
	array:	|----------------------------|			
	slice:       |<------->|  (像游标样的)
	
slice 是变长数组！	
	0  < len(s) <= cap(s)
声明格式： 		 var identifier []type // 不需要长度哦！
未被初始化的slice被设置为nil 0 长度！

初始化一个slice： var slice1 []type = arr1[start:end] 
start:end 被称为slice表达式

arr1[:]等价arr1[0:len(arr1)]
arr1[2:]等价于 arr1[2:len(arr1)]
arr1[:3]等价于 arr1[0:3]

slices 在内存中有3个字段：执行底层数组的指针；slice的长度；capacity slice的容积。

==================
### !!! 不要使用slice 的指针 slice本身就是指针类型了

### 传递slice 到function 
如果用数组做参数会导致数据复制的  slice本身是指针类型 所以传递slice到函数开销很小。example: funcX(arrX[:]) 

### 使用make创建slice
底层数组经常还没有定义，此时可以：var slice1 []type = make([]type , len)
字符串在某种意义是不可变的字节数组。所以 他们也可以被slice引用！

### new与make的区别
---------
他们都从堆上分配空间，但他们确实有区别，使用在不同的类型之上	。
- new(T): 为type T分配零基的存储并返回其地址，*T：返回一个指针 指向一个具有零值的类型T。用于值类型如array、struct。等价于&T{}
- make(T): 返回一个已经被初始化的类型T；仅仅只能被用在三个内建的类型上：slice，maps，channels；
简言之：new 分配； make 初始化。

### bytes 包
-----
bytes类型的slice很常用 所以出现了bytes包专门用来处理这种情况；很像strings包那样！
通过使用buffer 来链接字符串。这中用法很像java中的stringBuilder类
创建一个buffer 用buffer.WriteString(s)添加字符串,	 在结尾处用buffer.String()获取过程操作的最终结果。
~~~
var buffer bytes.Buffer
for {
	if s, ok := getNextString(); ok { //method getNextString() not 	shown here
	buffer.WriteString(s)
} else {
	break
}
}
fmt.Print(buffer.String(), “\n”)****
~~~

## range 结构
----
此结构可被应用在数组和slice上
~~~
	for ix,value := range slice1{
		
	}	
~~~
ix 是数组的索引value是该索引位置上的元素 他们是局部变量 只在该结构内部可见 是数组、slice在该索引位置上的值拷贝 只读的！

ix 可以是_  用来抛弃这个不用的索引 （如果你并不关心它）

如果你仅仅关心索引  那么可以漏掉第二个变量
~~~
	for idx := range slice1{
		
	}
	 
~~~
slices[idx] = {newValue}  可以用来修改正在被遍历的数组！

## reslice 
----
slice 初始化时经常会比底层的数组小：
>	slice1 := make([]type , start_lenght, capacity)
改变	slice 的长度被称为 reslice 如：
>	slice1 = slice1[0:len(slice1) +1] // 扩展下其长度
slice 可以改变大小 直至其占满底层数组！

##	拷贝 增长slice
----
为了扩容slice 我们必须创建一个新的 并拷贝原始slice中的内容到这个新的slice中
>  func append(s[]T , x ...T)[]T
如果增加的x 数量超过底层slice 的容积 那么append会分配一个新的 足够大的slice来盛放原有的元素和新增加的元素。这样返回的slice 可能就指向了一个不同的底层数组！！！ 增加操作总是会成功的 除非计算机内存用尽！

### 7.6.1 从字符串创建bytes slice ！！ to be continue



	