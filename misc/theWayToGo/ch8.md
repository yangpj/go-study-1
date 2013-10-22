#  MAP
====
无序键值对
亦称为关联数组或字典。其他语言中的hash dict HashTable ...

### 概念
map 是引用类型 声明形式：  var map1 map[keyType]valueType
map的长度在声明期是未知的 他可以动态增长！未初始化的map值是nil
key 是任何能够用== != 比较的类型：string int float 所以array slice 和struct不能被用于key，但指针和interface类型可以！！！一种用struct 作为key的办法是：Key() Hash()方法这样一个唯一的字符串或数组可以从结构内部的字段算出。
value 类型可以是任何类型。通过使用空接口，我们可以存储任何值，但使用它后我们可能需要是用type断言来做后续工作！

虽然map查询很快但仍比在slice或者数组中用索引访问要慢（100x）
 	v :=map1[key1]   如果key1 并不存在那么返回map值类型对应的零值.
!! 不要使用new  对map类型总是使用make！

###  capacity

	不同于array map无固定大小 但在使用make时可以可选的给一个cap
	
> make(map[keyType]valueType, cap)		

当map超过cap时 size会自动增一 所以对于大的map 他们经常增长很快的话最好给一个初始cap 大概的预估值也行。

### slice 作为map的值
当一对一关系时 key可以是原生类型，如果一个key对应多个值呢？
eg:
>  mp1 := make(map[int][]int)
>  mp2 := make(map[int]*[]int)

## 测试 建或者值存在于map中 删除一个元素
虽然访问一个key1对应的元素并没有加入map中得到其值为 值类型的零值； 但这是有歧义的（加入的值如果本身就是零值呢）
comma ok 模式：  val1, isPresent = map1[key1]

如果仅仅想知道是否存在：
>  _ , ok := map1[key1] // ok == true 如果键存在 不然就是false
>  if _ , ok := map1[key1]; ok{
	// ..
}

### 通过key删除元素：
> 	delete(map1,key1)


### for range 结构：
~~~
	for  key , value = range map1 {
		...
	}
~~~
如果你只重视值 那么用_ 忽略key
如果仅仅需要keys ：
>	for key := range map1{
	fmt.Printf("key is %d\n",key)
}


##  排序map
默认 map是无序的，如果想排序 那么先拷贝keys或者vaules到slice中 然后借用sort包中的方法排序slice 之后遍历被排序的keys或者values。



 


