package main

import "fmt"

func main() {
	//byte和rune类型
	//组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
	var a = '中'
	var b = 'x'
	fmt.Println(a, b)
	/*
		Go 语言的字符有以下两种：
			uint8类型，或者叫 byte 型，代表一个ASCII码字符。
			rune类型，代表一个 UTF-8字符。

		当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
	*/
	//Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
	// 遍历字符串
	//traversalString()
	//修改字符串
	changeString()
}

// 遍历字符串
func traversalString() {
	s := "hello 够浪"
	//ChanceStr := "积土成山，风雨兴焉；积水成渊，蛟龙生焉；积善成德，而神明自得，圣心备焉。故不积跬步，无以至千里；不积小流，无以成江海。骐骥一跃，不能十步；驽马十驾，功在不舍。锲而舍之，朽木不折；锲而不舍，金石可镂。蚓无爪牙之利，筋骨之强，上食埃土，下饮黄泉，用心一也。蟹六跪而二螯，非蛇鳝之穴无可寄托者，用心躁也。"
	//EnglishStr := "seveal months befors our trip,wang wei and i went to the library. We found a large atlas with good maps that showed details of world geography. "
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

}

// 输出
//104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 22815(够) 28010(浪)
//104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 229(å) 164(¤) 159(

/*
因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
*/

//修改字符串
//要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))

	s := "hello 够浪"
	//强制类型转换
	runneClic := []rune(s)
	runneClic[7] = '狼'
	s = string(runneClic)
	fmt.Println(s)
}
