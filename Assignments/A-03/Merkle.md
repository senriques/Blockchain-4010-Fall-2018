Merkle Trees
==============================

One of the major security and validity checks that blockchains
do is using Merkle trees.

In this assignment you will implement a Merkle tree hash.

When you pull code from git you should have a `./A-03/hash` directory.
This is a copy of Assignment 2's `./A-02/hash`.  The new code that you will
be working on is in `./merkle`.


### Pseudo Code

1. 

```
func InstructorMerkleHash(data [][]byte) []byte {
	// Build a place to put the hashes for the leaves
	hLeaf := make([][]byte, 0, len(data))
	// Calculate Leaf Hashes
	for ii := range data {
		aHash := hash.HashOf(data[ii])
		hLeaf = append(hLeaf, aHash)
	}
	// fmt.Printf("Leaf Hashes : %s, AT: %s\n", dumpSSB(hLeaf), godebug.LF())
	hMid := make([][]byte, 0, (len(hLeaf)/2)+1)
	ln := len(hLeaf)/2 + 1
	// fmt.Printf("ln=%d AT:%s\n", ln, godebug.LF())
	for ln >= 1 {
		// fmt.Printf("\n%s-------- TOP -------- AT:%s %s\n", MiscLib.ColorGreen, godebug.LF(), MiscLib.ColorReset)
		for ii := 0; ii < len(hLeaf); ii += 2 {
			// fmt.Printf("ii+1 = %d len(hLeaf) = %d AT:%s\n", ii+1, len(hLeaf), godebug.LF())
			if ii+1 < len(hLeaf) {
				hT := hash.Keccak256(hLeaf[ii], hLeaf[ii+1])
				hMid = append(hMid, hT)
				// fmt.Printf("AT:%s\n", godebug.LF())
			} else {
				hT := hash.Keccak256(hLeaf[ii])
				hMid = append(hMid, hT)
				// fmt.Printf("AT:%s\n", godebug.LF())
			}
		}
		hLeaf = hMid
		ln = len(hLeaf) / 2
		hMid = make([][]byte, 0, ln)
		// fmt.Printf("ln = %d Mid Hashes : %s, AT: %s\n", ln, dumpSSB(hLeaf), godebug.LF())
	}
	// fmt.Printf("\nFinal Hash : %s, AT: %s\n", dumpSSB(hLeaf), godebug.LF())
	return hLeaf[0]
```

# References

1. [Wikipedia has a nice discussion](https://en.wikipedia.org/wiki/Merkle_tree)
2. [Another explanation of Merkle Trees - with more details](https://brilliant.org/wiki/merkle-tree/)

