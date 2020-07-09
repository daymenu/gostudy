package main

import (
	"fmt"
	"math"
)

func main() {
	// 格式化十进制整数输出 %d 表示10进制整数
	fmt.Printf("整数：|%d|\n", math.MaxUint8)
	// + 表示 显示符号 10 表示 长度为10
	fmt.Printf("整数长度十位： |%+10d|\n", math.MaxUint8)
	// - 表示 字符要靠左展示
	fmt.Printf("整数长度十位： |%-+10d|\n", math.MaxUint8)
	// 格式化整数其他进制输出 %x 表示十六进制 %o表示八进制 %b 表示二进制 [1] 表示取后面第几个参数
	fmt.Printf("十进制: %d 二八十六进制表示 \t十六进制： |%+10[1]x|\t 八进制: |%+10[1]o| \t二进制： |%+10[1]b|\n", math.MaxUint8)
	// # 表示展示 二八十六进制的前缀 0x 0 0b
	fmt.Printf("十进制: %d 二八十六进制带前缀表示 \t十六进制： |%#+10[1]x|\t  八进制: |%#+10[1]o| \t二进制： |%#+10[1]b|\n", math.MaxUint8)
	// %f 表示浮点数
	fmt.Printf("浮点数保留两位小数：|%+10.2f|\t 浮点数：|%+10.2[1]g|\t 浮点数：|%+10.2[1]g|\n", float32(math.MaxUint8))
	// %t 表示 布尔
	fmt.Printf("布尔型： 真：|%t| \t 假: |%10t|\n", true, false)
	// %c 表示一个 rune unicode码点 rune 其实是32位的整
	fmt.Printf("中国的中： %c 十进制表示 %[1]d\n", []rune("中国")[0])
	// %s 表示一个字符串
	fmt.Printf("字符串： %s\n", "我是一个中国man，正在学习google的 go 语言，别看这里很多知识点，其实就是照搬c语言")
	// %q 表示字符串
	fmt.Printf("字符串加双引号： %q rune加单引号: %q \n", "hello 你好， 我是中国人", []rune("中国")[0])
	// %v 会以易读的形式打印出来
	fmt.Printf("字符串加双引号： %v 17: %v \n", "hello 你好， 我是中国人", 17)
	// %T 表示类型
	fmt.Printf("整型的类型： %T\n", math.MaxInt8)
}
