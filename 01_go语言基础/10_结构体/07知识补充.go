package main

import "fmt"

//因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。我们来看下面的例子：
type Person struct {
	Name   string
	Age    int8
	Dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.Dreams = dreams
}

//正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。
func (p *Person) SetDreams2(dreams []string) {
	p.Dreams = make([]string, len(dreams))
	copy(p.Dreams, dreams)
}
func main() {
	p1 := Person{Name: "老刘", Age: 12}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams2(data)
	fmt.Println(p1)

	data[1] = "不睡觉"
	fmt.Println(p1.Dreams) // SetDreams共用一个data， 对外面的data修改会影响到结构体内部
	fmt.Println(data)
}
