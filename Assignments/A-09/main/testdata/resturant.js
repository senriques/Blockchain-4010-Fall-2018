//$$spec
// ContractName ResturantPayment
// Constructor initResturant
// Upgrader upgradeResturant
// Methods PayToResturant PayBill PaymentFinished
// Views LeftToPay
//$$end-spec


// 3. Implement the multi-party payment with 15% tip.
//	PayToResurant ( Total, NParties, TipPercent ) -> ID
//	CalulatePaymetnsTip ( ID ) -> AmountPerPayment, AmountTip
//	PayBill ( Amt )
//	PayBill ( Amt )
//	PaymentFinished()


function initResturant( msg, data, cnst, plist) {
	data.GenID = 10000;
	data.CurPayments = [ 0 ];
	return data;
}

function upgradeResturant( msg, data, cnst, plist) {
	return data;
}

// PayToResturnat sets up a multip party payment to a
// resturnat and calculates the tip and how much each
// party should pay.
// 
// Inputs:
//		Total			Amount to pay for food.
//		NParties		How mahy to pay.
//		TipPercent		Percentage to pay as a tip.
//		ResturantAcct 	Account of resturant to send funds to.
//		TipAcct 		Account of server(s) to send tip to.
//
// Output: Event "PayToResturantEvent"
//		ID				ID of this tranaction.
//		Tip				What is the tolal amount of tip.
//		EachPay			How much will each person pay.
//	
function PayToResturant ( msg, data, plist ) {	
	var ID = data.GenID;
	data.GenID++;

	var tip = plist["Total"] * ( plist["TipPercent"] / 100.0 );
	var eachPay = ( plist["Total"] + tip ) / plist["NParties"];
	var rPay = plist["Total"] / plist["NParties"];
	var tPay = tip / plist["NParties"];

	data.CurPayments[ID] = {
		"Total": plist["Total"],
		"NParties": plist["NParties"],
		"TipPct": plist["TipPercent"],
		"Tip": tip,
		"EachPay": eachPay,
		"ResturantPay": rPay,
		"TipPay": tPay,
		"PayedBy": [],
		"ResturantAcct": plist["ResturantAcct"],
		"TipAcct": plist["TipAcct"],
		"PaymentRemaining": plist["Total"] + tip
	};

	tranactionEvent ( "PayToResturantEvent", {
		"ID": ID,
		"Tip": tip,
		"EachPay": eachPay
	} );

	return data;
}

// PayBill Make a payment in the multi party payment.  Payment includes
// the tip amount.
function PayBill ( msg, data, plist ) {
	var ID = plist["ID"];
	var amount = msg.amount;
	if ( amount != data.CurPayments[ID].EachPay ) {
		raiseError ( "Payment Incorrect, should be "+data.CurPayments[ID].EachPay );
		return data;
	}
	data.CurPayment[ID].PayedBy.push ( msg.sender )	
	transferTo ( msg.sender, data.CurPayment[ID].ResturantAcct,
		data.CurPayments[ID].ResturantPay );
	transferTo ( msg.sender, data.CurPayment[ID].TipAcct,
		data.CurPayments[ID].TipPay );
	data.CurPayments[ID].PaymentRemaining -= amount;
	return data;
}

// LeftToPay - view - how much is still left to be payed.
function LeftToPay ( msg, data, plist ) {
	var ID = plist["ID"];
	returnData ( data.CurPayments[ID].PaymentRemaining );
}

// PaymehnFinished - will return success and cleanup data if the payment process
// has finished..
function PaymentFinished ( msg, data, plist ) {
	if ( data.CurPayments[ID].PaymentRemaining === 0 ) {
		data.CurPayments[ID] = {};
	} else {
		raiseError ( "Payment not finialized.  Still owe: " + data.CurPayments[ID].PaymentRemaining );
	}
	returnData ( data );
}
