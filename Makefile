

all: README.html sillabus.html Handout01.html attestation.html class-summary.html Security.html

FR=./Lectures/Lect-01

README.html: README.md
	markdown-cli --input=./README.md --output=README.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ./md.css ${FR}/css/hpre README.html ${FR}/css/hpost >/tmp/README.html
	mv /tmp/README.html ./README.html

class-summary.html: class-summary.md
	markdown-cli --input=./class-summary.md --output=class-summary.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ./md.css ${FR}/css/hpre class-summary.html ${FR}/css/hpost >/tmp/class-summary.html
	mv /tmp/class-summary.html ./class-summary.html


sillabus.html: sillabus.md
	markdown-cli --input=./sillabus.md --output=sillabus.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ./md.css ${FR}/css/hpre sillabus.html ${FR}/css/hpost >/tmp/sillabus.html
	mv /tmp/sillabus.html ./sillabus.html



Handout01.html: Handout01.md
	markdown-cli --input=./Handout01.md --output=Handout01.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ./md.css ${FR}/css/hpre Handout01.html ${FR}/css/hpost >/tmp/Handout01.html
	mv /tmp/Handout01.html ./Handout01.html


attestation.html: attestation.md
	markdown-cli --input=./attestation.md --output=attestation.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ./md.css ${FR}/css/hpre attestation.html ${FR}/css/hpost >/tmp/attestation.html
	mv /tmp/attestation.html ./attestation.html

Security.html: Security.md
	markdown-cli --input=./Security.md --output=Security.html
	cat ${FR}/css/pre ${FR}/css/markdown.css ${FR}/css/post ./md.css ${FR}/css/hpre Security.html ${FR}/css/hpost >/tmp/Security.html
	mv /tmp/Security.html ./Security.html

