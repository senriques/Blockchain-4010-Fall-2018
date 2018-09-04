package mine

import (
	"encoding/hex"
	"fmt"

	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-04/block"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-04/hash"
)

// MineBlock implements a proof of work mining system where the first 4 digits (2 bytes) of the hex value are 0.
// Difficulty can be increaesed by requiring more digits to be 0 or by requring some other pattern to apear in
// the resulting hash.
func MineBlock(bk *block.BlockType, difficulty string) {
	for {
		data := block.SerializeForSeal(bk)
		hh := hash.HashOf(data)
		hhStr := hex.EncodeToString(hh)
		fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\r", hhStr, bk.Nonce)
		// This is where proof-of-stake will be checked
		if len(difficulty) == 0 || hhStr[0:len(difficulty)] == difficulty {
			bk.Seal = hh
			fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\n", hhStr, bk.Nonce)
			return
		}
		bk.Nonce++
	}
}
