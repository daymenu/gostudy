
#include "textflag.h"

GLOBL ·BoolValue(SB), NOPTR ,$1


GLOBL ·TrueValue(SB), NOPTR ,$1
DATA ·TrueValue(SB)/1,$1

GLOBL ·FalseValue(SB), NOPTR ,$1
DATA ·FalseValue(SB)/1,$0

GLOBL ·Float32Value(SB), NOPTR ,$1
DATA ·Float32Value(SB)/4,$1.5

GLOBL ·Float64Value(SB), NOPTR ,$1
DATA ·Float64Value(SB)/4,$0x01020304

GLOBL ·Helloworld(SB), NOPTR ,$16
GLOBL text<>(SB),NOPTR,$16
DATA text<>+0(SB)/8,$"Hello Wo"
DATA text<>+8(SB)/8,$"rld!"

DATA ·Helloworld+0(SB)/8,$text<>(SB)//StringHeader.Data
DATA ·Helloworld+8(SB)/8,$12//StringHeader.Len


GLOBL ·Hello(SB), NOPTR ,$24
DATA ·Hello+0(SB)/8,$text<>(SB)//StringHeader.Data
DATA ·Hello+8(SB)/8,$12//StringHeader.Len
DATA ·Hello+16(SB)/8,$16//StringHeader.Len

GLOBL ·Count(SB), RODATA ,$4
DATA ·Count+0(SB)/1,$1
DATA ·Count+1(SB)/1,$2
DATA ·Count+2(SB)/1,$3
DATA ·Count+3(SB)/1,$4
