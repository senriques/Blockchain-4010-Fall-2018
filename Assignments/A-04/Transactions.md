Transactions
-------------------------------

In bitcoin and Ethereum you can send funds from one account to another.
This activity is captured in a transaction.

In this assignment we are going to add transactions to our blockchain.

You will need to implement code in ./cli/cli.go starting at line 183.
The function is `func (cc *CLI) SendFundsTransaction(`

Side Note: notice how line continuation works in go with the
declaration of the function.

The function calls
`cc.InstructorSendFundsTransaction` remove that. That is the 
instructors version of the code (The answer that I implemented).

Work through the pseudo code and implement the transaction.

Basically a transaction is finding all the outputs for an account
that do not have any corresponding  input.  These are the unused
outputs that represent the value of the account.   Fortunately
we have an index that tells us where to find these.  Verify that
there is sufficient funds in the account then create inputs for
our new block/new transaction that collects all of the funds.
Given that we have the `sum` of the funds, now create 1 or 2
output transactions.  First a transaction output to the destination
`to` account.  If the amount of the funds is larger than the
transferred amount then some "change" is owed back to the "from"
account.  If "change" is needed then create a transaction output
with the "change".

The pseudo code is in a comment in the file `./cli/cli.go` and it
is reproduced below.

```
	//
	// Pseudo Code:
	// 1. Calcualte the total value of the account 'from'.  Call this 'tot'.
	//    You can do this by calling `cc.GetTotalValueForAccount(from)`.
	// 2. If the total, `tot` is less than the amount that is to be transfered,
	//	  `amount` then fail.  Return an error "Insufficient funds".  The person
	//    is trying to bounce a check.
	// 3. Get the list of output tranactions ( ../transactions/tx.go TxOutputType ).
	//    Call this 'oldOutputs'.
	// 4. Find the set of (may be empty - check for that) values that are pointed
	//    to in the index - from the 'from' account.  Delete this from the
	//    index.
	// 5. Create a new empty transaction.  Call `transctions.NewEmptyTx` to create.
	//	  Pass in the 'memo' and the 'from' for this tranaction.
	// 6. Convert the 'oldOutputs' into a set of new inputs.  The type is
	//    ../transctions/tx.go TxInputType.  Call `transactions.CreateTxInputsFromOldOutputs`
	//	  to do this.
	// 7. Save the new inputs in the tx.Input.
	// 8. Create the new output for the 'to' address.  Call `transactions.CreateTxOutputWithFunds`.
	//    Call this `txOut`.    Take `txOut` and append it to the tranaction by calling
	//    `transactions.AppendTxOutputToTx`.
	// 9. Calcualte the amount of "change" - if it is larger than 0 then we owe 'from'
	//    change.  Create a 2nd tranaction with the change.  Append to the tranaction the
	//    TxOutputType.
	// 10. Return
	//
```
