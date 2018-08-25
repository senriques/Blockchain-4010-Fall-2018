
all: Lect-01.html

Lect-01.html: Lect-01.md
	markdown-cli --input=./Lect-01.md --output=Lect-01.html
	echo cat ../../md.css Lect-01.html >/tmp/Lect-01.html
	cat ./css/pre ./css/markdown.css ./css/post ../../md.css ./css/hpre Lect-01.html ./css/hpost >/tmp/Lect-01.html
	mv /tmp/Lect-01.html ./Lect-01.html

run:
	./run.sh

