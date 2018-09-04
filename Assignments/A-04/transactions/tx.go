package transactions

import (
	"bytes"
	"encoding/binary"

	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-04/addr"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-04/lib"
)

/*
Transactions
-------------------------------

In bitcoin and Ethereum you can send funds from one account to another.
This activity is captured in a transaction.

In this assignment we are going to add transactions to our blockchain.


yyyy yyyy yyyy yyyy

1. Find all the inputs - this is the input set.
	1. Use the index to find the transaction
2. Add up the value - that is how much you can spend
3. If the value is <= the spend then
	0. Check to see that the sender has signed for this spend - if so then...
	1. Split into the spend to the new account
	2. If what remains is > 0 then give change
	3. Add in a new input/output with to the "mining" account for doing the transaction

Add these new transactions to the current block.


Steps - copy last 2 assignmetns
1. Implement Tx
2. Add Tx to Block - test
3. Add mining-account to config - test
4. Searalize TxInputType, TxOutputType, TxType - test
5. Read / Write - test
7. Add in merkel hash - test
8. Perform a "send-funds" operation

*/

type TransactionType struct {
	TxOffset       int               // The position of this in the block.
	Input          []TxInputType     // Set of inputs to a transaction
	Output         []TxOutputType    // Set of outputs to a tranaction
	SCOwnerAccount addr.AddressType  // Used in HW 6 - Smart Contracts - Owner of Contract
	SCAddress      addr.AddressType  // Used in HW 6 - Smart Contracts - Address of Contract
	SCOutputData   string            // Used in HW 6 - Smart Contracts - Output data from contract
	Account        addr.AddressType  //
	Signature      lib.SignatureType //	Signature that this is a valid tranaction for this account
	Message        string            //	Message (hash of original) that signature signs for
	Comment        string            //
}

type TxInputType struct {
	BlockNo     int // Which block is this from
	TxOffset    int // The transaction in the block. In the block[BlockHash].Tx[TxOffset]
	TxOutputPos int // Position of the output in the transaction. In the  block[BlockHash].Tx[TxOffset].Output[TxOutptuPos]
	Amount      int // Value
}

type TxOutputType struct {
	BlockNo     int              // Which block is this in
	TxOffset    int              // Which transaction in this block.  block[this].Tx[TxOffset]
	TxOutputPos int              // Position of the output in this block. In the  block[this].Tx[TxOffset].Output[TxOutptuPos]
	Account     addr.AddressType // Acctount funds go to (If this is ""chagne"" then this is the same as TransactionType.Account
	Amount      int              // Amoutn to go to accoutn
}

func NewEmptyTx(memo string, passedAcct addr.AddressType) *TransactionType {
	return &TransactionType{
		Input:          make([]TxInputType, 0, 1),  // Set of inputs to a transaction
		Output:         make([]TxOutputType, 0, 1), // Set of outputs to a tranaction
		SCOwnerAccount: []byte{},
		SCAddress:      []byte{},
		SCOutputData:   "",
		Account:        passedAcct,
		Signature:      lib.SignatureType([]byte{}),
		Message:        "",
		Comment:        memo,
	}
}

// SearalizeTransaction searializes into bytes the fields that will be hashed the hash of the block.
// This is the hash that the next block will use to point to this block and the hash that
// this block will be saved as.
func SerializeTransaction(tx *TransactionType) []byte {
	var buf bytes.Buffer

	for _, inp := range tx.Input {
		SerializeTransactionInput(&buf, &inp)
	}
	for _, out := range tx.Output {
		SerializeTransactionOutput(&buf, &out)
	}
	buf.Write([]byte(tx.SCOwnerAccount))
	buf.Write([]byte(tx.SCAddress))
	buf.Write([]byte(tx.SCOutputData))
	buf.Write([]byte(tx.Account))
	buf.Write([]byte(tx.Signature))
	buf.Write([]byte(tx.Message))
	buf.Write([]byte(tx.Comment))

	return buf.Bytes()
}

// SearalizeForSeal searializes into bytes the fields that will be hashed for the mining seal.
func SerializeForSeal(tx *TransactionType) []byte {
	return SerializeTransaction(tx)
}

func SerializeTransactionInput(buf *bytes.Buffer, inp *TxInputType) {
	// binary.Write(buf, binary.BigEndian, inp.BlockNo)
	// buf.Write([]byte(inp.BlockHash))
	binary.Write(buf, binary.BigEndian, inp.BlockNo)
	binary.Write(buf, binary.BigEndian, inp.TxOffset)
	binary.Write(buf, binary.BigEndian, inp.TxOutputPos)
	binary.Write(buf, binary.BigEndian, inp.Amount)
}

func SerializeTransactionOutput(buf *bytes.Buffer, out *TxOutputType) {
	binary.Write(buf, binary.BigEndian, out.BlockNo)
	binary.Write(buf, binary.BigEndian, out.TxOffset)
	binary.Write(buf, binary.BigEndian, out.TxOutputPos)
	buf.Write([]byte(out.Account))
	binary.Write(buf, binary.BigEndian, out.Amount)
}

func CreateTxOutputWithFunds(to addr.AddressType, amount int) (txo *TxOutputType, err error) {
	txo = &TxOutputType{
		Account: to,
		Amount:  amount,
	}
	return
}

func AppendTxOutputToTx(tx *TransactionType, txo *TxOutputType) {
	txo.TxOutputPos = len(tx.Output)
	tx.Output = append(tx.Output, *txo)
}

func CreateTxInputsFromOldOutputs(oldOutput []*TxOutputType) (txIn []TxInputType, err error) {
	for _, vv := range oldOutput {
		txIn = append(txIn, TxInputType{
			BlockNo:     vv.BlockNo,
			TxOffset:    vv.TxOffset,
			TxOutputPos: vv.TxOutputPos,
			Amount:      vv.Amount,
		})
	}
	return
}
