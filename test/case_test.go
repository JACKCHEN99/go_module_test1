package test

import (
	"fmt"
	"testing"
)

// 定义一个Talker接口
type Talker interface {
	// 定义Talk方法，返回一个字符串
	Talk() string
}

// 定义一个Greeter结构体
type Greeter struct {
	Greeting string
}

// 定义一个Talk方法，返回Greeting属性
func (g Greeter) Talk() string {
	return g.Greeting
}

// 定义一个Dog结构体
type Dog struct{}

// 定义一个Talk方法，返回字符串"Woof!"
func (d Dog) Talk() string {
	return "Woof!"
}

// 定义一个PerformTalk函数，参数为Talker类型
func PerformTalk(t Talker) {
	// 打印出Talk方法返回的字符串
	fmt.Println(t.Talk())
}

// 定义一个TestAny函数，用于测试
func TestAny(t *testing.T) {
	// 定义一个Talker变量
	var a Talker
	// 将Greeter结构体赋值给Talker变量
	a = Greeter{Greeting: "Hello"}
	// 调用PerformTalk函数，传入Talker变量
	PerformTalk(a)

	// 将Dog结构体赋值给Talker变量
	a = Dog{}
	// 调用PerformTalk函数，传入Talker变量
	PerformTalk(a)
}

// 在这个例子中，我们定义了一个 Talker 接口和两个实现了这个接口的类型：Greeter 和 Dog
// 我们还编写了一个 PerformTalk 函数，它接受一个 Talker 类型的参数
// 在 main 函数中，我们可以看到 PerformTalk 能够处理任何实现了 Talker 接口的类型，无论是 Greeter 还是 Dog
// 这就是接口的强大之处：它使得我们的代码更加灵活和通用。
// 因此，虽然在简单的例子中直接调用方法似乎足够，但在更复杂的应用中，接口的使用能显著提高代码的质量和可维护性。
