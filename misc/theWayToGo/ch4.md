
文件命名
-------------
+  小写 与下划线   xx_jj.go        无空格和其他特殊字符

命名
--------------
+  _ 下划线是个特殊的标识符  ：blank identifier  可被赋值（左值） 但内容会被丢弃（不能当右值）

+  匿名类型和函数 是允许存在的 他们被称为 anonymous

+ 包是用来组织代码的方式、简称pkg 每个go文件隶属（一个且一个）包（类似其他语言的库或者名空间概念）多个不同的.go 文件可以隶属于一个包 第一行就是用来声明文件隶属关系的。可执行入口文件隶属于main包 每个go应用包含一个包叫main

+ 一个应用包含多个不同的包，即便你的应用只有一个main包 你也不必把所有代码都塞在一个单独的文件中。
+ 别名 import  your_alias  "pkg"

go项目的一般结构
------------
+ 在import 之后：声明变量，常量，和类型s。
+ 之后是init() 函数（如果有的话--0个或多个哦！， 此函数是个特殊函数 每个包中的这个函数s会被先执行）
+　在之后就是main() 入口函数了（只能声明其包为main才可以）
+ 在之后才是其他函数 ，类型的方法先出现，或者以main中调用的顺序来声明函数，或者用某种逻辑顺序来排列函	数、方法的顺序
例子：
~~~
package main
import (
“fmt”
)
const c = “C”
var v int = 5
type T struct{}
func init() { // initialization of package
}
func main() {
var a int
Func1()
// ...
fmt.Println(a)
}
func (t T) Method1() {
//...
}
func Func1() { // exported function Func1
//...
}
~~~

程序执行顺序
-----------
+ 在main包中的所有包 会以其出现的顺序被导入。
在每个包中：
+ 如果它也导入了其他包，那么导入规则仍旧同上。但是已经被导入的包只会被导入一次。
+　之后对于每一个包（逆序）所以的常量，变量被计算执行。之后init 方法被执行（如果有的话）。
+ 最后在main包中 以上两条过程才发生 之后main方法开始执行。

惯例
------------
go的类型转换只能是显式的（不能隐式automatic转换），语法类似方法调用：
~~~
	valueOfTypeB = typeB(valueOfTypeA)
	// 如：
	a := 5.0
	b := int(a)
~~~
转型只能发生在特定的两种类型之间。比如窄类型到宽类型（int16==》int32） 当从宽类型到窄类型转换时
可能存在损失精度的危险。还有那些不能被转换的类型如果编译器侦测到会报编译错误的 不然就会报运行期错误。
具有共同底层类型的变量是可被互相装换的（类型别名那种情况！）。

命名
------
最好短，精确，有意义。最好不要再带包前缀了。返回对象的方法应该用名词而不是如其他语言中的getter（用jquery中的命名方式就好了 比如 user() 《==》getUser() ）,改变一个对象可以用setXxx() （即setter）
如果必要 最好使用骆驼命名法（首字大写代表包可见性哦！） 而不是带下划线那种。

常量
------------
不可变的数据   只能是boolean number string 类型的
格式如：
~~~
 const identifier [type] = value 
 // 例如：
	const Pi = 3.141592
~~~
常量是在编译期被计算的。
常量可被用于枚举：
~~~
const (
	Unknown = 0
	Female = 1
	Male = 2
)
~~~

-----
##var##
- var 语法主要用于global 包级别可见性， 在方法中经常用短声明语法 :=

-	go的指针是引用类型 slices,maps channels 都是引用类型，被引用的变量被存储在堆上，它是可被GC管理的并且比栈内存空间要大些。

-  Printf 方法时fmt包中导出的
~~~
	// 方法签名：
	func Printf(format string , list of variables to be printed)
~~~
	%s 表示string值
	%v 是普适的默认格式符
- 短形式的赋值操作符  := 只能用在方法内部 而不是在包作用域。很高效地创建一个新的变量，亦被称作初始化声明。
- 只定义局部变量	并不使用它会导致编译错误 如下：
~~~
	func main(){
		var	a string = "abc"
		fmt.Println("hello, world")
		 
	}
~~~
- 变量交换： a, b = b ,a 

init 函数
---------
除了全局定义并实例化 变量也可以在init方法中实例化。init方法不可以被调用 但是会被自动执行的（先于main方法）每个源文件只能含一个init函数？ 初始化是单线程的包依赖保证其正确的执行顺序。
- 经常被用在程序启动时的后台goroute守护协程
- %t 在格式化字符时 可用在bool类型的变量上 
- 对bool类型的变量命名  劲量以 is Is 开头
- 对浮点类型 精确的表示是不可能的（有些浮点数二进制表示不出来的！） 所以在用== 或者!= 对比时要格外注意。
- 尽量多用float64 math包中普遍都用这个类型进行运算。（不然转型问题需要考虑哦）
- %d 用来格式化整数 %x 和%X 用来格式化十六进制 %g浮点
	%0nd 整数有n个数字。
	
- a23bitInt = int32(a32Float)	的转型操作会截断 所以有时必要检验转型的安全性：
~~~
func Uint8FromInt(n int) (uint8, error) {
if 0 <= n && n <= math.MaxUint8 { // conversion is safe
return uint8(n), nil
}
return 0, fmt.Errorf(“%d is out of the uint8 range”, n)
}
	func IntFromFloat64(x float64) int {
if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
whole, fraction := math.Modf(x)
if fraction >= 0.5 {
whole++
}
return int(whole)
}
panic(fmt.Sprintf(“%g is out of the int32 range”, x))
}
~~~


