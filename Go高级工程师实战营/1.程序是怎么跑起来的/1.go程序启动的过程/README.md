## 环境准备

1. cat dockerfile

   ```dockerfile
   FROM centos
   
   RUN yum install golang -y \ 
   
   && yum install dlv -y \
   
   && yum install binutils -y \ 
   
   && yum install vim -y \ 
   
   && yum install gdb -y
   ```

2. docker build -t test .

3. docker run -it --rm test bash



## 编译

1. 准备文本 `cat main.go`

   ```go
   package main
   
   import "fmt"
   
   func main() {
       fmt.Println("hello")
   }
   ```

2. 进行编译`go build main.go`

3. 查看编译产物`ls`

   ```shell
   [root@0e5e5af7ef31 demo]# ls
   main main.go
   ```

4. 清理编译文件和缓存`rm -f main`，`rm -rf /root/.cache/go-build/`

5. 使用`go build -x main.go`查看编译过程

   ```shell
   [root@0e5e5af7ef31 demo]# go build -x main.go 
   WORK=/tmp/go-build2845090676
   mkdir -p $WORK/b001/
   cat >$WORK/b001/_gomod_.go << 'EOF' # internal
   package main
   import _ "unsafe"
   //go:linkname __debug_modinfo__ runtime.modinfo
   var __debug_modinfo__ = "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nmod\tcommand-line-arguments\t(devel)\t\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
   EOF
   cat >$WORK/b001/importcfg << 'EOF' # internal
   # import config
   packagefile fmt=/usr/lib/golang/pkg/linux_amd64/fmt.a
   packagefile runtime=/usr/lib/golang/pkg/linux_amd64/runtime.a
   EOF
   cd /root/go/src/github.com/1005281342/demo
   /usr/lib/golang/pkg/tool/linux_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid 0q_2jJiX3iThnZL4awvF/0q_2jJiX3iThnZL4awvF -goversion go1.16.12 -D _/root/go/src/github.com/1005281342/demo -importcfg $WORK/b001/importcfg -pack -c=4 ./main.go $WORK/b001/_gomod_.go
   /usr/lib/golang/pkg/tool/linux_amd64/buildid -w $WORK/b001/_pkg_.a # internal
   cp $WORK/b001/_pkg_.a /root/.cache/go-build/c6/c68e2d874e04c58ecf34c0b11a2652d7c64e2534dc4ce07395a51d56b70e8e6e-d # internal
   cat >$WORK/b001/importcfg.link << 'EOF' # internal
   packagefile command-line-arguments=$WORK/b001/_pkg_.a
   packagefile fmt=/usr/lib/golang/pkg/linux_amd64/fmt.a
   packagefile runtime=/usr/lib/golang/pkg/linux_amd64/runtime.a
   packagefile errors=/usr/lib/golang/pkg/linux_amd64/errors.a
   packagefile internal/fmtsort=/usr/lib/golang/pkg/linux_amd64/internal/fmtsort.a
   packagefile io=/usr/lib/golang/pkg/linux_amd64/io.a
   packagefile math=/usr/lib/golang/pkg/linux_amd64/math.a
   packagefile os=/usr/lib/golang/pkg/linux_amd64/os.a
   packagefile reflect=/usr/lib/golang/pkg/linux_amd64/reflect.a
   packagefile strconv=/usr/lib/golang/pkg/linux_amd64/strconv.a
   packagefile sync=/usr/lib/golang/pkg/linux_amd64/sync.a
   packagefile unicode/utf8=/usr/lib/golang/pkg/linux_amd64/unicode/utf8.a
   packagefile internal/bytealg=/usr/lib/golang/pkg/linux_amd64/internal/bytealg.a
   packagefile internal/cpu=/usr/lib/golang/pkg/linux_amd64/internal/cpu.a
   packagefile runtime/internal/atomic=/usr/lib/golang/pkg/linux_amd64/runtime/internal/atomic.a
   packagefile runtime/internal/math=/usr/lib/golang/pkg/linux_amd64/runtime/internal/math.a
   packagefile runtime/internal/sys=/usr/lib/golang/pkg/linux_amd64/runtime/internal/sys.a
   packagefile internal/reflectlite=/usr/lib/golang/pkg/linux_amd64/internal/reflectlite.a
   packagefile sort=/usr/lib/golang/pkg/linux_amd64/sort.a
   packagefile math/bits=/usr/lib/golang/pkg/linux_amd64/math/bits.a
   packagefile internal/oserror=/usr/lib/golang/pkg/linux_amd64/internal/oserror.a
   packagefile internal/poll=/usr/lib/golang/pkg/linux_amd64/internal/poll.a
   packagefile internal/syscall/execenv=/usr/lib/golang/pkg/linux_amd64/internal/syscall/execenv.a
   packagefile internal/syscall/unix=/usr/lib/golang/pkg/linux_amd64/internal/syscall/unix.a
   packagefile internal/testlog=/usr/lib/golang/pkg/linux_amd64/internal/testlog.a
   packagefile io/fs=/usr/lib/golang/pkg/linux_amd64/io/fs.a
   packagefile sync/atomic=/usr/lib/golang/pkg/linux_amd64/sync/atomic.a
   packagefile syscall=/usr/lib/golang/pkg/linux_amd64/syscall.a
   packagefile time=/usr/lib/golang/pkg/linux_amd64/time.a
   packagefile internal/unsafeheader=/usr/lib/golang/pkg/linux_amd64/internal/unsafeheader.a
   packagefile unicode=/usr/lib/golang/pkg/linux_amd64/unicode.a
   packagefile internal/race=/usr/lib/golang/pkg/linux_amd64/internal/race.a
   packagefile path=/usr/lib/golang/pkg/linux_amd64/path.a
   EOF
   mkdir -p $WORK/b001/exe/
   cd .
   /usr/lib/golang/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=cINSmju7QDsk0H7CWqKP/0q_2jJiX3iThnZL4awvF/K76lx1xOGuydgqcIkUd4/cINSmju7QDsk0H7CWqKP -extld=gcc $WORK/b001/_pkg_.a
   /usr/lib/golang/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal
   mv $WORK/b001/exe/a.out main
   rm -r $WORK/b001/
   ```

   编译代码产出.a文件

   ```shell
   /usr/lib/golang/pkg/tool/linux_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid 0q_2jJiX3iThnZL4awvF/0q_2jJiX3iThnZL4awvF -goversion go1.16.12 -D _/root/go/src/github.com/1005281342/demo -importcfg $WORK/b001/importcfg -pack -c=4 ./main.go $WORK/b001/_gomod_.go
   /usr/lib/golang/pkg/tool/linux_amd64/buildid -w $WORK/b001/_pkg_.a # internal
   ```

   链接完成.out文件

   ```shell
   /usr/lib/golang/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=cINSmju7QDsk0H7CWqKP/0q_2jJiX3iThnZL4awvF/K76lx1xOGuydgqcIkUd4/cINSmju7QDsk0H7CWqKP -extld=gcc $WORK/b001/_pkg_.a
   /usr/lib/golang/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal
   ```

   重命名可执行文件`mv $WORK/b001/exe/a.out main`




## 执行

### 可执行文件

Linux的可执行文件为`ELF(Executable and Linkable Format)`

由`ELF header`、`Section header`和`Sections`三部分构成

[elf101.pdf](https://github.com/corkami/pics/blob/28cb0226093ed57b348723bc473cea0162dad366/binary/elf101/elf101.pdf)

### 执行步骤

1. 解析`ELF header`
2. 加载文件内容至内存
3. 从`entry point`开始执行代码

通过`readelf`查看`entry point`：

```shell
[root@bdb892eb3ea4 /]# readelf -h main
ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00 
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x465760
  Start of program headers:          64 (bytes into file)
  Start of section headers:          456 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         7
  Size of section headers:           64 (bytes)
  Number of section headers:         23
  Section header string table index: 3
```

通过`dlv`找到go进程启动位置：

```shell
[root@bdb892eb3ea4 /]# dlv exec ./main
Type 'help' for list of commands.
(dlv) b *0x465760
Breakpoint 1 set at 0x465760 for _rt0_amd64_linux() .usr/lib/golang/src/runtime/rt0_linux_amd64.s:8
```

**可以看见是从`runtime`中启动的**



## runtime

而在程序运行时自动 加载/运行的一些模块。诸如python、java等语言没有runtime。

runtime主要包括：scheduler、netpoll、memory management、garbage collector等四个模块，其中scheduler负责串联整个runtime。



## 协程

当使用go func() { // ...}时，实际上是提交了一个任务给runtime scheduler。

[将在这里持续整理](https://1005281342.gitbook.io/gofun/go-xie-cheng/goroutine)

