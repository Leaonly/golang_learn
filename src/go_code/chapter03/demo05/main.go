package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println("i=", i)

	var n2 int64 = 10
	fmt.Printf("n2数据类型是 %T ,n2字节数是 %d\n", n2, unsafe.Sizeof(n2))
	// shift+alt+↓ 复制下一行
	fmt.Printf("n2数据类型是 %T ,n2字节数是 %d\n", n2, unsafe.Sizeof(n2))

	var c1 int = '北'
	fmt.Printf("c1=%c \n", c1)

	var c2 int = 10 + 'a'
	fmt.Printf("c1=%c \n", c2)

	var s1 int = 'h'
	s1 = s1 + 10
	fmt.Println(s1)

	var s2 = "hello"
	fmt.Printf("%T\n",s2)

	// 反引号，可以打印批量字符串
	var str3 = `
	package main

	// import "fmt"
	// import "unsafe"
	
	import (
		"fmt"
		"unsafe"
	)
	`
	fmt.Println(str3)

	// 字符串拼接
	var str4 = "Hello " + "World! "
	str4 += "你好!"
	fmt.Println(str4)
}
