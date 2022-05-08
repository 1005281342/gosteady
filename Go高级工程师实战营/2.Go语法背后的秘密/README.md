[TOC]

## 编译原理基础

### 编译过程

![](README.assets/截屏2021-12-25 下午5.07.25.png)

相关工具：

语法分析工具：https://astexplorer.net/

中间代码（SSA）生成：通过https://golang.design/gossa?id=6acdea5f-6563-11ec-b3e4-0242ac160010 （这里已经准备了一个例子）或`GOSSAFUNC=funcname go build x.go`命令

机器码生成：https://godbolt.org/z/5Tnenzq7M（这里已经准备了一个例子）

### 链接过程

代码：

```go
package main

func main() {
	println("hi, go")
}

```

#### 编译后产生目标文件：

编译命令：`go tool compile -S x.go`

```text
"".main STEXT size=86 args=0x0 locals=0x18 funcid=0x0
	0x0000 00000 (x.go:3)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (x.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (x.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (x.go:3)	PCDATA	$0, $-2
	0x000d 00013 (x.go:3)	JLS	79
	0x000f 00015 (x.go:3)	PCDATA	$0, $-1
	0x000f 00015 (x.go:3)	SUBQ	$24, SP
	0x0013 00019 (x.go:3)	MOVQ	BP, 16(SP)
	0x0018 00024 (x.go:3)	LEAQ	16(SP), BP
```

编译生成`.o`目标文件，未进行链接，汇编代码中main函数还没有包相关信息，指令从起点0x0000开始偏移。

#### 链接后产生可执行文件：

使用go build x.go完成变异，然后通过`go tool objdump ./x | grep "x.go"`查看编译并链接后的汇编代码

```text
TEXT main.main(SB) /x.go
  x.go:3		0x45ec60		64488b0c25f8ffffff	MOVQ FS:0xfffffff8, CX			
  x.go:3		0x45ec69		483b6110		CMPQ 0x10(CX), SP			
  x.go:3		0x45ec6d		7640			JBE 0x45ecaf				
  x.go:3		0x45ec6f		4883ec18		SUBQ $0x18, SP				
  x.go:3		0x45ec73		48896c2410		MOVQ BP, 0x10(SP)			
  x.go:3		0x45ec78		488d6c2410		LEAQ 0x10(SP), BP			
  x.go:4		0x45ec7d		0f1f00			NOPL 0(AX)				
  x.go:4		0x45ec80		e8db05fdff		CALL runtime.printlock(SB)		
  x.go:4		0x45ec85		488d0574620100		LEAQ 0x16274(IP), AX			
  x.go:4		0x45ec8c		48890424		MOVQ AX, 0(SP)				
  x.go:4		0x45ec90		48c744240807000000	MOVQ $0x7, 0x8(SP)			
  x.go:4		0x45ec99		e8e20ffdff		CALL runtime.printstring(SB)		
  x.go:4		0x45ec9e		6690			NOPW					
  x.go:4		0x45eca0		e83b06fdff		CALL runtime.printunlock(SB)		
  x.go:5		0x45eca5		488b6c2410		MOVQ 0x10(SP), BP			
  x.go:5		0x45ecaa		4883c418		ADDQ $0x18, SP				
  x.go:5		0x45ecae		c3			RET					
  x.go:3		0x45ecaf		e8acafffff		CALL runtime.morestack_noctxt(SB)	
  x.go:3		0x45ecb4		ebaa			JMP main.main(SB)
```

**main函数由包信息，并且指令偏移起点非0x0000**

### go compile和go tool objdump的区别

go compile是完成编译产生`.o`文件

objdump是把可执行文件反汇编得到汇编代码

#### 通过调试工具查看编译链接后的汇编代码

`readelf -h ./x`查看入口地址

`dlv exec ./x`开始调试

`b *0x45cd80`

`disass`

## 语法实现分析

通过工具查看汇编代码结合goland或者dlv进行调试分析



## 函数调用规约



## 参考

[走进Golang之编译器原理](https://xiaomi-info.github.io/2019/11/13/golang-compiler-principle/)

