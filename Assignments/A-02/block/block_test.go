package block

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_GenesisBlock(t *testing.T) {
	gb := InitGenesisBlock()
	if !IsGenesisBlock(gb) {
		t.Errorf("Should be genesis block")
	}
}

func Test_InitBlock(t *testing.T) {
	bk := InitBlock(12, "Hello World", []byte{1, 2, 3})
	if IsGenesisBlock(bk) {
		t.Errorf("Should not be genesis block")
	}
	if bk.Index != 12 {
		t.Errorf("Expected index of 12")
	}
	exp := "592fa743889fc7f92ac2a37bb1f5ba1daf2a5c84741ca0e0061d243a2e6707ba"
	got := fmt.Sprintf("%x", bk.ThisBlockHash)
	if exp != got {
		t.Errorf("Block hash incorrect, expected ->%s<- got ->%s<-", exp, got)
	}
}

func SVarI(v interface{}) string {
	// s, err := json.Marshal ( v )
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	}
	return string(s)
}

/*
// xyzzy - todo - test of
func SearalizeBlock(bk *BlockType) []byte {
func SearalizeForSeal(bk *BlockType) []byte {
*/
