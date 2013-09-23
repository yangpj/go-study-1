#   Functions

-----------------------
方法时go语言的基本构造块

##  介绍
-----------
每个程序都由一些函数构成：基本的代码构件块

由于go是编译性语言 其出现的顺序并不重要！尽管如此 可读性上最好始于main 并以逻辑顺序来书写function（比如在main中的调用顺序）

函数主要用于任务分解 代码重用！《重构》
DRY： Dont Repeat Yourself 执行特定任务的代码只出现一次！
三种形式的function：
	- 有标示符的函数
	- 匿名函数或lambda
	- methods 方法
任何形式都可以有参数和返回值
方法调用的基本形式：
packX.Function(arg1,arg2,..,argn) 方法总是被另一个方法调用----那个方法被称为 调用者方法 calling function  一个方法可以调用任意多 （任意深）的方法 知道栈耗尽

方法重载在go中是不允许的！会引发编译错误！方法重载强制runtime做额外的类型匹配 会消弱性能的。这样在go中你不得不为你的方法起个合适的名字 

声明一个外部实现的方法 你仅仅只需给出其签名即可 不需要方法体：
	func   flushICache(begin, end uintptr) // implementd externally
funcitons 也可以用在声明形式中 作为一种function类型：
	type binOp func(int,int) int
	// 仍旧没有方法体哦！
function 也可以被赋值给变量  ： add := binOp 
	add获取了引用 知道其指向的方法签名 不可以在被赋给另一种签名了！
function 的值可以被比较 : 如果他们指向同一个方法 或者都是nil 他们就被认为相等. 

function 不允许嵌套声明 但可以用匿名函数来模拟！

go 目前没有generice 范式 ， 可用interface实现 特别是空接口！但相应性能也降低了 所以如果着重性能那么最好使用明确类型来创建函数！

## 参数和返回值
-------------
go	 中可以返回零个或者多个值 这种技术在测试是否方法执行错误时特别方便。
无参的方法 称为niladic function 如：main.main().

### call by value / Call by reference
--------
go 的默认方式是传值调用的。参数被复制后传到function环境下 方法中对参数的修改不会影响原始数据。
如果需要修改原始数据 那么可以传递引用  通过取址操作 & 这样传递一个指针给function 。此时拷贝的就不是变量指向的数据了 而是地址被拷贝传递了。通过指针可以改变原始值。

slice，map interface channel 默认就是以引用传递的。
一些function只是执行一个人任务 ， 并不返回值：被称为 side-effect（边缘效应），如打印到控制台，发送email，logging错误，等..

!! 使用named返回变量：使得代码更清晰，更短并且自文档化。

### blank identifier
----
	_ 可以用来丢弃值。
	
##  6.3传递不定数目的参数	
----
如果方法的最后一个参数是：...type 这就表明方法可以接受不定量的那个类型的参数。甚至是0：称之为variadic funciton

形式： func myFunc(a,b, arg ...int){}

设想如下方法：  
~~~
	func Greeting(prefix string, who ...string)
	// 调用  Greeting("hello :" ,"Joe","Anna","Eilleen")

~~~
在Greeting方法内部 who 拥有的值是： []string{"Joe","Anna","Eilleen"}
!!! 如果参数列表被存为数组arr 那么可以传递 arr...  Greeting("hello",arr...)  
不定长参数:
~~~
	function F1(s … string) {
		F2(s …)
		F3(s)
	}
	func F2(s … string) { }
	func F3(s []string) { }
~~~
可以被传递 到其slice类型的函数中去！

变长参数类型不一样时 有两种解决方案：
1. 使用结构
	type Options struct{
		par1 type1
		par2 type2
			...
	}
传递参数： F1(a,b,Options{})
如果可选参数多于一个值： F1(a,b,Options{par1:val1, par2:val2)

2. 使用空接口 interface{}
如果使用空接口 那么可能需要type-switch 结构去检测：
~~~
 		func typecheck(values ...interface{}){
			for _ , value := range values{
				switch v := value.(type){
					case int: 
						....
					case float: ..
					case string: ...
					case bool: ...
					default : ...						
				}
				
			}
		}
~~~

## 6.4 defer and tracing 
-----
defer 类似某些语言中的finally块 如java C#,经常用来回收分配的资源
如果多次调用了defer 那么其执行顺序是其出现顺序的倒序。
例子：
- defer file.Close() // 关闭文件
- 
	~~~
		mu.Lock()
		defer mu.Unlock() // 解锁
		
	~~~
- 报表中打印footer
	~~~
		printHeader()
		defer printFooter()
	~~~	
- 关闭数据库连接
	~~~
		// OPEN A db CONNECTION
		defer  disconnecitonFromDB()
		
	~~~	
	
用defer来追踪函数执行：
	~~~
		func trace(s string){ fmt.Println("entering",s)}
		func untrace(s string ){ fmt.Println("leaving : ",s)}
		
	~~~

内置函数 ， 不隶属任何包
- close 				用在channel通讯
- len	cap 	
- new 	make
	new 用在值类型和用户自定义类型（struct）		
	make 用于内置引用类型（slice ， maps ， channels）
- copy  append 	
- panic recover 
- print println    更底层的函数 在生产环境中用fmt包中的
- complex real imag 用来操作复杂数字

##  funciton 作为参数传递
----
function 做为参数传递给另一个function 经常被称为callback	

## 6.8 Closure (function literals)
----
有时并不想给函数起名字。
		

								


