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


Steps - copy last 2 assignmetns
1. Implement Tx
2. Add Tx to Block - test
3. Add mining-account to config - test
4. Searalize TxInputType, TxOutputType, TxType - test
5. Read / Write - test
7. Add in merkel hash - test
8. Perform a "send-funds" operation


---------------------------------------------------------------------

Goal: Be able to do a transaction - send funds from A to B - and create a new block with it in it.
	Be able to add a tranaction with funds for an address

ReadIndex

If ValidSignature () {
	tot := TotalAvailFunds ( FromAddr )
	if tot >= Amount {
		SendFunds ( FromAddr, ToAddr, Amount ) Using: Index, Block		-- AddTxToBlock
		if tot > Amount {
			change = tot-Amount
			SendFunds ( FromAddr, FromAddr, change )					-- AddTxToBlock
		}
	}
}
Mine Block
MineReward ( MinderSelfAddr, AmountOfReward )						
	Index
	Block

Update index to have this block in it.
WriteBlock

WriteIndex
