
all: Mining.html build

build:
	( cd main ; go build )

# You can get markdown-cli and compile it yourself.
#
# $ cd ~/go/src/github.com/
# $ mkdir pschlump
# $ cd pschlump
# $ git pull https://github.com/pschlump/markdown-cli.git
# $ go get
# $ cd markdown-cli
# $ go build
#
Mining.html: Mining.md
	markdown-cli --input=./Mining.md --output=Mining.html

