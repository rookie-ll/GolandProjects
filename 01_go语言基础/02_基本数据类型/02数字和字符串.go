package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		 浮点型
			Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准：
			float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。
			float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。
	*/
	//打印浮点数时，可以使用fmt包配合动词%f，代码如下：
	//fmt.Printf("%f\n", math.Pi)
	//fmt.Printf("%.2f\n", math.Pi)

	/*
		 复数
			complex64和complex128
			复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
	*/
	//var c1 complex64
	//c1 = 1 + 2i
	//var c2 complex128
	//c2 = 2 + 3i
	//fmt.Println(c1)
	//fmt.Println(c2)

	/*
		 布尔值
			Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

			注意：
				布尔类型变量的默认值为false。
				Go 语言中不允许将整型强制转换为布尔型.
				布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/
	/*
		 字符串
			Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
			Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：
	*/
	//s1 := "hello"
	//s2 := "你好"
	//fmt.Println(s1, s2)

	//字符串转义符
	/*
		转义符	含义
		\r	回车符（返回行首）
		\n	换行符（直接跳到下一行的同列位置）
		\t	制表符
		\'	单引号
		\"	双引号
		\\	反斜杠
	*/

	//多行字符串
	//s3 := `第一行
	//		第二行
	//		第三行
	//`
	//fmt.Println(s3)

	// 字符串的常用操作

	/*
		方法	介绍
		len(str)	求长度
		+或fmt.Sprintf	拼接字符串
		strings.Split	分割
		strings.contains	判断是否包含
		strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
		strings.Index(),strings.LastIndex()	子串出现的位置
		strings.Join(a[]string, sep string)	join操作
	*/
	//chanceStr_ := "青，取之于蓝，而青于蓝；冰，水为之，而寒于水。"
	ChanceStr := "积土成山，风雨兴焉；积水成渊，蛟龙生焉；积善成德，而神明自得，圣心备焉。故不积跬步，无以至千里；不积小流，无以成江海。骐骥一跃，不能十步；驽马十驾，功在不舍。锲而舍之，朽木不折；锲而不舍，金石可镂。蚓无爪牙之利，筋骨之强，上食埃土，下饮黄泉，用心一也。蟹六跪而二螯，非蛇鳝之穴无可寄托者，用心躁也。"
	//EnglishStr := "seveal months befors our trip,wang wei and i went to the library. We found a large atlas with good maps that showed details of world geography. "
	ECStr := "hello 够浪"
	ECStr2 := "hello 老李"
	// len(str)
	fmt.Println("len", len(ChanceStr))

	// + 活fmt.Sprintf
	newStr := ECStr2 + " " + ECStr
	fmt.Println(newStr)

	newStr2 := fmt.Sprintf("%v %v", ECStr, ECStr2)
	fmt.Println(newStr2)

	//strings.Split
	newStr3 := strings.Split(ChanceStr, "，")
	fmt.Println(newStr3)

	//strings.contains
	isContain := strings.Contains(ChanceStr, "积土成山")
	fmt.Println(isContain) // true

	//strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	hasPrefix := strings.HasPrefix(ChanceStr, "积")
	fmt.Println(hasPrefix) // true

	//strings.Index(),strings.LastIndex()	子串出现的位置
	indexInt := strings.Index(ChanceStr, "故不积跬步")
	fmt.Println(indexInt) // 108
	indexInt2 := strings.LastIndex(ChanceStr, "故不积跬步")
	fmt.Println(indexInt2) // 108

	//strings.Join(a[]string, sep string)	join操作
	newStr4 := strings.Join(newStr3, "-")
	fmt.Println(newStr4)
}
