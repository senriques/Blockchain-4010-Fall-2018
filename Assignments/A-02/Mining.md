What is Mining?
=============================================================

At the root of our blockchain is a special block called a "genesis" block.  
Basically the "genesis" block is the beginning block in our block chain.  It is
special because it will not point back to any previous block.

The code that you are given can write out the genesis block.  And
an index to where to find it.   In most blockchains like Bitcoin 
some sort of a database is used for storing the blocks.

We are not going to do that.  We are going to store all of them
in files in the file system.   This is so you can see the blocks.
Also we are going to store the blocks in JSON as text so you can
read the blocks.  Using a database is faster but has lots of overhead.
It can also be very frustrating when you are attempting to determine
if the block is correct and you can't easily see what is in the 
block.

Our blocks will be written (by default) in the `./data` directory.
The format for the blocks is `hash.json`, where hash is the block
hash.

To get started, first checkout the code for Assignment 2 - this 
can be done by:

```sh
	cd ~/go/src/github.com/
	mkdir Univ-Wyo-Education
	cd Univ-Wyo-Education
	git clone github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018.git
```

If you have already done this you should update your copy of the class
repository with:

```sh
	cd ~/go/src/github.com/Univ-Wyo-Education
	git pull
```

Then change directory into assignment 2.

```sh
	cd Assignments/A-02
```

Our starting code is in this directory.  Specifically we will want to
compile the main program.  It is in ./main.  Cd to that directory.
You shoud end up in:
`~/go/src/github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02`


```sh
	cd ./main
	go build
```

Run main to create the genesis block and the initial index.

```sh
	./main --create-genesis
```

This should create a directory with 2 files in it.  The default 
is in the ./data irectory.  The files are:

```
	136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e.json
	index.json
```

You should look at the contents of the 2 files.  The one with the long name
is our genesis block.  The `index.json` is an index that will allow us to
find data blocks as we are building the this blockchain.

The code is missing the chunk that will do the block mining.  The stubbed
out function is in ./mine/mine.go.  Your assignment is to implement the
body of the function.  You will want to verify that it works by running

```sh
	go test
```

int the `./mine` directory.

Take the time to go and poke through the code.  This code is the basis
for your mid-term project.  You are going to need to be familiar with
all of it.  Run all the tests.   If you have questions about it now is
the time to be asking them.

### Pseudo Code for the mining.

In the file ./mine/mine.go, implement the function MineBlock.

1. Use an infinite loop to:
  1. Serialize the data from the block for hashing, Call block.SearilizeForSeal to do this.
  2. Calculate the hash of the data, Call hash.HashOf to do this. This is the slow part.  What would happen if we
     replaced the software with a hash calculator on a graphics card where you could run 4096 hahes at once?
     What would happen if we replaced the graphics card with an ASIC - so you had dedicated hardware to do
     the hash and you could run 4 billion hashes a second?
  3. Convert the hash (it is []byte) to a hex string.  Use the hex.EncodeToString standard go library function.
  4. `fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\r", theHashAsAString, bk.Nonce)`
  5. See if the first 4 characters of the hash are 0's. - if so we have met the work criteria.
     In go this is `if theHashAsAString[0:4] == "0000" {`.  This is create a slice, 4 long from
     character 0 with length of 4, then compare that to the string `"0000"`.
   - Set the block's "Seal" to the hash
   - `fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\n", theHashAsAString, bk.Nonce)`
   - return
  5. Increment the Nonce in the block, and...
  6. Back to the top of the loop for another try at finding a seal for this block.


In the ./mine directory there is a test.  Complete the MineBlock function.  Run the
test.

Submit your completed ./mine/mine.go file.  

My output when running the test.

```
+> go test
	((Mining)) Hash for Block [0000ae2cab130b4836988969f731c4f884ac4675790e5575a5161e5b96ab13d7] nonce [   54586]
	((Mining)) Hash for Block [0000adc29a80f1f0df08c8687c013d179050f5d1b449599e4d1437e4fad23525] nonce [   46734]
	((Mining)) Hash for Block [000013ce557332aaa68abe3b7bf1be856743a03689a802606a732e81713bb78c] nonce [    4527]
	PASS
	ok  	github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/mine	0.914s
```

The grader has a somewhat more comprehensive automated test to run with this code (There is one more block
to mine).   You get all the points for the assignment when it passes the test.

### Before you submit your code!

Use the go formatter on your code.  Either `goimports -w *.go` or setup your editor to run
goimpors every time you save a go file.  I have this setup in `vim`.  Other editors can
do this also.

Run `go vet` and `golint *.go` on it.  Fix any errors.


