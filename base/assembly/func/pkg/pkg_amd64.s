// TEXT symbol(SB), [flags,] $framesize[-argsize]
// TEXT 用于定义函数符号，函数名中当前包的路径可以省略
// symbol(SB) 函数的名字后面是(SB)，表示是函数名符号相对于SB伪寄存器的偏移量，二者组合在一起最终是绝对地址
// [flags,] 标志部分用于指示函数的一些特殊行为，标志在textlags.h文件中定义，常见的NOSPLIT主要用于指示叶子函数不进行栈分裂
// framesize部分表示函数的局部变量需要多少栈空间，其中包含调用其它函数时准备调用参数的隐式栈空间
// [-argsize] 最后是可以省略的参数大小，之所以可以省略是因为编译器可以从Go语言的函数声明中推导出函数参数的大小
#include "textflag.h"

// func Swap(a, b int) (int, int)
TEXT ·Swap(SB), NOSPLIT, $0-32
    MOVQ a+0(FP),AX // AX = a
    MOVQ b+8(FP),BX // BX = b
    MOVQ BX,ret0+16(FP) // ret0 = BX
    MOVQ AX,ret1+24(FP) // ret1 = AX
    RET

// func Swap(a, b int) (int, int)
// TEXT ·Swap(SB), NOSPLIT, $0

