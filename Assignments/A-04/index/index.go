package index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-04/addr"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-04/block"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-04/lib"
)

/*

Example Data (sort of with stuff ...ed out)
	{
		"Index": [ "126c..", .... ]
		"TxHashIndex": {
			"1212121...": {
				"Index": 0,
				"BlockHash": "1212222",
			},
			"2212121...": {
				"Index": 1,
				"BlockHash": "1212222",
			}
		}
	}

*/

type TxHashIndexType struct {
	Index     int
	BlockHash string
}

type AddrHashIndexType struct {
	Addr addr.AddressType
	Data string
}

type AddressIndex struct {
	AddrIndex map[string]AddrHashIndexType
}

type BlockIndex struct {
	Index       []string
	TxHashIndex map[string]TxHashIndexType
	AddrData    AddressIndex
}

func ReadIndex(fn string) (idx *BlockIndex, err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	idx = &BlockIndex{}
	err = json.Unmarshal(buf, idx)
	if err != nil {
		err = fmt.Errorf("invalid initialization - Unable to parse JSON file, %s", err)
		return nil, err
	}
	return
}

func WriteIndex(fn string, bkslice []*block.BlockType) {
	indexForBlocks := BuildIndex(bkslice)
	ioutil.WriteFile(fn, []byte(lib.SVarI(indexForBlocks)), 0644)
}

func BuildIndex(bkslice []*block.BlockType) (idx BlockIndex) {
	for _, bk := range bkslice {
		idx.Index = append(idx.Index, fmt.Sprintf("%x", bk.ThisBlockHash))
	}
	idx.TxHashIndex = make(map[string]TxHashIndexType)
	for ii, bk := range bkslice {
		idx.TxHashIndex[fmt.Sprintf("%x", bk.ThisBlockHash)] = TxHashIndexType{
			Index:     ii,
			BlockHash: fmt.Sprintf("%x", bk.ThisBlockHash),
		}
	}
	return
}
