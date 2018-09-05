package main


func main() {
}

// golang set集合去重以及交叉并集计算
func main1() {
	// 我这边有个场景是针对数据去重, 但又不是简单的去重, 是有时间区间范围内的交集、并集计算!!!
	// 废话不多说, 我估计有不少人记不清集合的并集、差集、交集的概念了!
	
	/*
		集合的分类:
		
		并集:		以属于A或属于B的元素为元素的集合成为A与B的并(集)
		交集:		以属于A且属于B的元素为元素的集合成为A与B的交(集)
		差集:		以属于A而不属于B的元素为元素的集合成为A与B的差(集)
			
	*/
}

// 如果只是去重的化, 完全可以使用golang map实现, 比如下面的例子就可以实现去重!
func main2() {
	m := map[string]int{
		"baidu.com":  0,
		"apple.com":  0,
		"google.com": 0,
		"google.com": 0,
	}
}

// 自定义集合Set
// 在Go语言中有作为Hash Table实现的字典(Map)类型, 但标准数据类型中并没有集合(Set)这种数据类型
// 比较Set和Map的主要特性, 有类似特性如下:
/*
		1.它们中的元素都是不可重复的;
		2.它们都只能用迭代的方式取出其中的所有元素;
		3.对它们中的元素进行迭代的顺序都是与元素插入顺序无关的, 同时也不保证任何有序性
*/
//	但是, 它们之间也有一些区别, 如下:
/*
		Set的元素是一个单一的值, 而Map的元素则是一个键值对;
		Set的元素不可重复指的是不能存在任意两个单一值相等的情况, Map的元素不可重复指的是任意两个键值对中的键的值不能相等。
*/
// 从上面的特性可知, 可以把集合类型(Set)作为字典类型(Map)的一个简化版本;
// 也就是说, 可以用Map来编写一个Set类型的实现;
// 实际上, 在Java语言中, java.util.HashSet类就是用java.util.HashMap类作为底层支持的!
// 所以这里就从HashSet出发, 逐步抽象出集合Set!

/*
	1.定义HashSet
			首先, 在工作区的src目录的代码包basic/set(可以自行定义, 但后面要保持一致)中, 创建一个名为hash_set.go的源码文件;
			根据代码包 basic/set 可知，源码文件 hash_set.go 的包声明语句（关于这个一些规则可以看前面的系列博文）如下：
			上面提到可以将集合类型作为字典类型的一个简化版本。
			现在我们的 HashSet 就以字典类型作为其底层的实现。
			HashSet声明如下:
			
			type HashSet struct {
				m map[interface{}]bool
			}
			
			如上声明HashSet类型中的唯一的字段的类型是map[interface{}]bool。
			选择这样一个字典类型是因为通过将字典m的键类型设置为interface{}, 让HashSet的元素可以是任何类型的, 
					因为这里需要使用m的值中的键来存储HashSet类型的元素值。
			那使用bool类型作为m的值的元素类型的好处如下:
					从值的存储形式的角度看, bool类型值只占用一个字节!!!
					从值的表示形式的角度看, bool类型的值只有两个—true和false, 并且, 这两个值度都是预定义常量!

			把bool类型作为值类型更有利于判断字典类型值中是否存在某个键。
			例如:
					如果在向m的值添加键值对的时候总是以true作为其中的元素的值, 则索引表达式 m[“a”]的结果值总能体现出在m的值中是否包含键为“a”的键值对。
			对于 map[interface{}]bool 类型的值来说, 如下:
					if m["a"] {// 判断是否m中包含键为“a”的键值对
						//省略其他语句
					}
						
	2.如上HashSet类型的基本结构已确定了, 现在考虑如何初始化HashSet类型值;
			由于字典类型的零值为nil, 而用new函数来创建一个HashSet类型值, 也就是new(HashSet).m的求值结果将会是一个nil(关于new函数可以查阅本人另一篇博文Go语言学习笔记5)。
			因此, 这里需要编写一个专门用于创建和初始化HashSet类型值的函数, 该函数声明如下:
						func NewHashSet() *HashSet {
							return &HashSet{m: make(map[interface{}]bool)}
						}
			如上可以看到, 使用make函数对字段m进行了初始化。
			同时注意观察函数NewHashSet的结果声明的类型是*HashSet而不是HashSet, 目的是让这个结果值的方法集合中
					包含调用接收者类型为HashSet或*HashSet的所有方法。
			这样做的好处将在后面编写Set接口类型的时候再予以说明!

	3.实现HashSet的基本功能
			依据其他编程语言中的HashSet类型可知, 它们大部分应该提供的基本功能如下:
				添加元素值。
				删除元素值。
				清除所有元素值。
				判断是否包含某个元素值。
				获取元素值的数量。
				判断与其他HashSet类型值是否相同。
				获取所有元素值, 即生成可迭代的快照。
				获取自身的字符串表示形式。
	
	4.
	5.
	6.
	7.
*/

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

// 1.添加元素值
//		方法Add会返回一个bool类型的结果值, 以表示添加元素值的操作是否成功。
//		方法Add的声明中的接收者类型是*HashSet。
/*
		这里使用 *HashSet 而不是 HashSet, 主要是从节约内存空间的角度出发, 分析如下:
			(1).当Add方法的接收者类型为HashSet的时候, 对它的每一次调用都需要对当前HashSet类型值进行一次复制。
						虽然在HashSet类型中只有一个引用类型的字段, 但是这也是一种开销。
						而且这里还没有考虑HashSet类型中的字段可能会变得更多的情况。
			(2).当Add方法的接收者类型为*HashSet的时候, 对它进行调用时复制的当前*HashSet的类型值只是一个指针值。
					    在大多数情况下, 一个指针值占用的内存空间总会被它指向的那个其他类型的值所占用的内存空间小。
						无论一个指针值指向的那个其他类型值所需的内存空间有多么大, 它所占用的内存空间总是不变的。
*/
func (set *HashSet) Add(e interface{}) bool {
    if !set.m[e] {
		// 当前的m的值中还未包含以e的值为键的键值对
		set.m[e] = true				// 将键为e(代表的值)、元素为true的键值对添加到m的值当中
		return true					// 添加成功
    }
	
	return false	// 添加失败
}

// 2.删除元素值
// 调用delete内建函数删除HashSet内部支持的字典值
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)	// 第一个参数为目标字典类型, 第二个参数为要删除的那个键值对的键
}

// 3.清除所有元素
// 为HashSet中的字段m重新赋值
/*
	如果接收者类型是HashSet, 该方法中的赋值语句的作用只是为当前值的某个复制品中的字段m赋值而已, 而当前值中的字段m则不会被重新赋值。
	方法Clear中的这条赋值语句被执行之后, 当前的HashSet类型值中的元素就相当于被清空了。
	已经与字段m解除绑定的那个旧的字典值由于不再与任何程序实体存在绑定关系而成为了无用的数据。
		它会在之后的某一时刻被Go语言的垃圾回收器发现并回收。
*/
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

// 4.判断是否包含某个元素值
//		方法Contains用于判断其值是否包含某个元素值。
//		这里判断结果得益于元素类型为bool的字段m
/*
		当把一个interface{}类型值作为键添加到一个字典值的时候, Go语言会先获取这个interface{}类型值的实际类型(即动态类型), 
				然后再使用与之对应的hash函数对该值进行hash运算, 
				也就是说, interface{}类型值总是能够被正确地计算出hash值!
		但是字典类型的键不能是函数类型、字典类型或切片类型, 否则会引发一个运行时恐慌, 并提示如下:
					panic: runtime error: hash of unhashable type <某个函数类型、字典类型或切片类型的名称>
*/
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

// 5.获取元素值的数量
//		方法Len用于获取HashSet元素值数量
func (set *HashSet) Len() int {
	return len(set.m)
}

// 6.判断与其他HashSet类型值是否相同
// 方法Same用来判断两个HashSet类型值是否相同
/*
	两个HashSet类型值相同的必要条件是, 它们包含的元素应该是完全相同的。
	由于HashSet类型值中的元素的迭代顺序总是不确定的, 所以也就不用在意两个值在这方面是否一致。
	如果要判断两个HashSet类型值是否是同一个值, 就需要利用指针运算进行内存地址的比较!!!
*/
func (set *HashSet) Same(other *HashSet) bool {
    if other == nil {
        return false
    }
    if set.Len() != other.Len() {
        return false
    }
    for key := range set.m {
        if !other.Contains(key) {
            return false
        }
    }
    return true
}

// (7).获取所有元素值，即生成可迭代的快照。
/*
		所谓快照, 就是目标值在某一个时刻的映像。
		对于一个HashSet类型值来说, 它的快照中的元素迭代顺序总是可以确定的, 快照只反映了该HashSet类型值在某一个时刻的状态。
		另外, 还需要从元素可迭代且顺序可确定的数据类型中选取一个作为快照的类型。
		这个类型必须是以单值作为元素的, 所以字典类型最先别排除。
		又由于HashSet类型值中的元素数量总是不固定的, 所以无法用一个数组类型的值来表示它的快照。
		如上分析可知, Go语言中可以使用的快照的类型应该是一个切片类型或者通道类型。
*/

// 方法Elements用于生成快照
// 注意:
//			在Elements方法中针对并发访问和修改m的值的情况采取了一些措施。
//			但是由于m的值本身并不是并发安全的, 所以并不能保证Elements方法的执行总会准确无误。
//			要做到真正的并发安全, 还需要一些辅助的手段, 比如读写互斥量。
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)		// 获取HashSet中字段m的长度，即m中包含元素的数量
	
	// 初始化一个[]interface{}类型的变量snapshot来存储m的值中的元素值
    snapshot	:= make([]interface{}, initialLen)
	actualLen	:= 0
	
	// 按照既定顺序将迭代值设置到快照值(变量snapshot的值)的指定元素位置上, 这一过程并不会创建任何新值。
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
        } else {
			// m的值中的元素数量有所增加, 使得实际迭代的次数大于先前初始化的快照值的长度
			snapshot = append(snapshot, key)	// 使用append函数向快照值追加元素值。
        }
        actualLen++	//实际迭代的次数
    }
	
	// 对于已被初始化的[]interface{}类型的切片值来说, 未被显示初始化的元素位置上的值均为nil。
	// m的值中的元素数量有所减少, 使得实际迭代的次数小于先前初始化的快照值的长度。
	// 这样快照值的尾部存在若干个没有任何意义的值为nil的元素.
	// 可以通过snapshot = snapshot[:actualLen]将无用的元素值从快照值中去掉。
    if actualLen < initialLen {
        snapshot = snapshot[:actualLen]
	}
	
	return snapshot
}

// (8).获取自身的字符串表示形式。
//		这个String方法的签名算是一个惯用法。
//		代码包fmt中的打印函数总会使用参数值附带的具有如此签名的String方法的结果值作为该参数值的字符串表示形式。
func (set *HashSet) String() string {
	var buf bytes.Buffer			// 作为结果值的缓冲区
    buf.WriteString("HashSet{")
    first := true
	
    for key := range set.m {
		if first {
            first = false
        } else {
            buf.WriteString(",")
        }
        buf.WriteString(fmt.Sprintf("%v", key))
    }
    //n := 1
    //for key := range set.m {
    //  buf.WriteString(fmt.Sprintf("%v", key))
    //  if n == len(set.m) {//最后一个元素的后面不添加逗号
    //      break;
    //  } else {
    //      buf.WriteString(",")
    //  }
    //  n++;
    //}
    buf.WriteString("}")
    return buf.String()  
}

// 如上已经完整地编写了一个具备常用功能的Set的实现类型, 后面将讲解更多的高级功能来完善它!!!

/*
			高级功能
*/

// 集合Set的真包含的判断功能,  根据集合代数中的描述, 如果集合A真包含了集合B, 那么就可以说集合A是集合B的一个超集!
//	判断集合set是否是集合other的超集
func (set *HashSet) IsSuperset(other *HashSet) bool {
    if other == nil {
		// 如果other为nil, 则other不是set的子集
        return false
    }
	
	setLen		:= set.Len()	// 获取set的元素值数量
	otherLen	:= other.Len()	// 获取other的元素值数量
	
	if setLen == 0 || setLen == otherLen {
		// set的元素值数量等于0或者等于other的元素数量
        return false
    }
	
	if setLen > 0 && otherLen == 0 {
		// other为元素数量为0，set元素数量大于0，则set也是other的超集
        return true
    }
	
    for _, v := range other.Elements() {
        if !set.Contains(v) {
			// 只要set中有一个包含other中的数据，就返回false
            return false
        }
    }
	
	return true
}

/*
		集合的运算包括并集、交集、差集和对称差集;
		并集运算是指把两个集合中的所有元素都合并起来并组合成一个集合;
		交集运算是指找到两个集合中共有的元素并把它们组成一个集合。 
		集合 A 对集合 B 进行差集运算的含义是找到只存在于集合 A 中但不存在于集合 B 中的元素并把它们组成一个集合。 
		对称差集运算与差集运算类似但有所区别。
		      对称差集运算是指找到只存在于集合 A 中但不存在于集合 B 中的元素，再找到只存在于集合 B 中但不存在于集合 A 中的元素，最后把它们合并起来并组成一个集合。
*/

// 实现并集运算
//		生成集合set和集合other的并集
func (set *HashSet) Union(other *HashSet) *HashSet {
	if set == nil || other == nil {	// set和other都为nil, 则它们的并集为nil
		return nil
    }
	
	unionedSet := NewHashSet()			// 新创建一个HashSet类型值, 它的长度为0, 即元素数量为0
	for _, v := range set.Elements() {	// 将set中的元素添加到unionedSet中
		unionedSet.Add(v)
    }
	
	if other.Len() == 0 {
		return unionedSet
	}
	
	for _, v := range other.Elements() {	// 将other中的元素添加到unionedSet中, 如果遇到相同, 则不添加(在Add方法逻辑中体现)
		unionedSet.Add(v)
	}
	
	return unionedSet
}

// 实现交集运算
//		生成集合set和集合other的交集
func (set *HashSet) Intersect(other *HashSet) *HashSet {
	if set == nil || other == nil {	
		// set和other都为nil, 则它们的交集为nil
		return nil
    }
	
	intersectedSet := NewHashSet()	// 新创建一个HashSet类型值, 它的长度为0, 即元素数量为0
		
	if other.Len() == 0 {			
		// other的元素数量为0, 直接返回intersectedSet
		return intersectedSet
	}
	
	if set.Len() < other.Len() {
		// set的元素数量少于other的元素数量
		for _, v := range set.Elements() {
			// 遍历set
			if other.Contains(v) {
				// 只要将set和other共有的添加到intersectedSet
				intersectedSet.Add(v)
			}
		}
	} else {
		// set的元素数量多于other的元素数量
		for _, v := range other.Elements() {
			// 遍历other
            if set.Contains(v) {
				// 只要将set和other共有的添加到intersectedSet
				intersectedSet.Add(v)
			}
		}
	}
	
	return intersectedSet
}

// 

func main3() {
	
}



