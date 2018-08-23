package block

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pschlump/godebug"
)

// ReadBlock reads in from a file a block and returns it.
func ReadBlock(fn string) (bk *BlockType, err error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Printf("Unable to read genesis block %s [at %s], %s\n", fn, godebug.LF(), err)
		return nil, err
	}
	bk = &BlockType{}
	err = json.Unmarshal(data, bk)
	if !IsGenesisBlock(bk) {
		fmt.Printf("Unable to read genesis block %s [at %s], %s\n", fn, godebug.LF(), err)
		return nil, err
	}
	return bk, nil
}
