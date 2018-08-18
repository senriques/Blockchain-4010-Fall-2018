
all: Intro-To-Go.html

Intro-To-Go.html: Intro-To-Go.md
	markdown-cli --input=./Intro-To-Go.md --output=Intro-To-Go.html

test: test01 test02
.PHONEY: test test01 test02

test01:
	( cd hello-world ; go build ; ./hello-world ) >,a
	diff ,a ref/hw.out

test02:
	( cd echo1 ; go build ; ./echo1 a b c X ) >,b
	diff ,b ref/echo1.out

clean:
	rm ,a ,b echo1/echo1 hello-world/hello-world


