
all: Lect-04.html panic.html

FR=../Lect-01/

Lect-04.html: Lect-04.md
	markdown-cli --input=./Lect-04.md --output=Lect-04.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre Lect-04.html ${FR}/css/hpost >/tmp/Lect-04.html
	mv /tmp/Lect-04.html ./Lect-04.html

panic.html: panic.md
	markdown-cli --input=./panic.md --output=panic.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ../../md.css ${FR}/css/hpre panic.html ${FR}/css/hpost >/tmp/panic.html
	mv /tmp/panic.html ./panic.html

run:
	./run.sh

