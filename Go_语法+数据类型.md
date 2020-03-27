- 一行代表一个语句结束，多个语句写在一行用;分开
- 标识符：用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。如 a_123、1a
- 用`+`连接字符串
# 关键字
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var
# 预定义标识符
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr
# Go 语言中变量的声明必须使用空格隔开 `var age int`
# Go数据类型
1.布尔型（bool） var b bool = true
2.整数型 (int float32  float64) var f float32 = 3.21
3.字符串类型(string) var str = "hello"
4.派生类型
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型
5.数字类型
    1	uint8
    无符号 8 位整型 (0 到 255)
    2	uint16
    无符号 16 位整型 (0 到 65535)
    3	uint32
    无符号 32 位整型 (0 到 4294967295)
    4	uint64
    无符号 64 位整型 (0 到 18446744073709551615)
    5	int8
    有符号 8 位整型 (-128 到 127)
    6	int16
    有符号 16 位整型 (-32768 到 32767)
    7	int32
    有符号 32 位整型 (-2147483648 到 2147483647)
    8	int64
    有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
6.浮点型
    1	float32
    IEEE-754 32位浮点型数
    2	float64
    IEEE-754 64位浮点型数
    3	complex64
    32 位实数和虚数
    4	complex128
    64 位实数和虚数
7.其他数字型
    1	byte
    类似 uint8
    2	rune
    类似 int32
    3	uint
    32 或 64 位
    4	int
    与 uint 一样大小
    5	uintptr
    无符号整型，用于存放一个指针

#  :=和var 为声明语句,声明过的变量不能使用:= 可以将 var f string = "Runoob" 简写为 f := "Runoob"

# 常量是一个简单值的标识符，在程序运行时，不会被修改的量。常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。`const identifier [type] = value
iota
iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

iota 可以被用作枚举值：

//select基本用法
select {
case <- chan1:
// 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
default:
// 如果上面都没有成功，则进入default处理流程
