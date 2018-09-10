
# A Go Core/Panic 

First the Code

```Go
	package main

	import "fmt"

	var mm map[string]int

	func main() {
		fmt.Println("vim-go")
		mm["bob"] = 3
	}
```

Then the bad output.

```
panic: assignment to entry in nil map

goroutine 1 [running]:
panic(0x10a5540, 0x10d03a0)
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/panic.go:551 +0x3c1 fp=0xc42005fec8 sp=0xc42005fe28 pc=0x1025bb1
runtime.mapassign_faststr(0x10a4e80, 0x0, 0x10be68a, 0x3, 0x0)
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/hashmap_fast.go:696 +0x407 fp=0xc42005ff38 sp=0xc42005fec8 pc=0x1009de7
main.main()
	/Users/corwin/go/src/github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Lectures/Lect-04/samp.go:9 +0x92 fp=0xc42005ff88 sp=0xc42005ff38 pc=0x108d6d2
runtime.main()
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/proc.go:198 +0x212 fp=0xc42005ffe0 sp=0xc42005ff88 pc=0x10278c2
runtime.goexit()
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc42005ffe8 sp=0xc42005ffe0 pc=0x104d981

goroutine 2 [force gc (idle)]:
runtime.gopark(0x10c5580, 0x11397a0, 0x10bfafe, 0xf, 0x10c5414, 0x1)
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/proc.go:291 +0x11a fp=0xc42004a768 sp=0xc42004a748 pc=0x1027d0a
runtime.goparkunlock(0x11397a0, 0x10bfafe, 0xf, 0x14, 0x1)
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/proc.go:297 +0x5e fp=0xc42004a7a8 sp=0xc42004a768 pc=0x1027dbe
runtime.forcegchelper()
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/proc.go:248 +0xcc fp=0xc42004a7e0 sp=0xc42004a7a8 pc=0x1027b4c
runtime.goexit()
	/usr/local/Cellar/go/1.10.3/libexec/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc42004a7e8 sp=0xc42004a7e0 pc=0x104d981
```

"assignment to entry in nil map"

..."/Lect-04/samp.go:9"...

