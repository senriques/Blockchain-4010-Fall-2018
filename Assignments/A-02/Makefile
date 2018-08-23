

all: Mining.html build

build:
	( cd main ; go build )

Mining.html: Mining.md
	markdown-cli --input=./Mining.md --output=Mining.html

