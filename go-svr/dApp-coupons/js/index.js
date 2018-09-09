
// var web3 = new Web3(new Web3.providers.HttpProvider("http://192.168.0.199:8545"));
var web3 = new Web3(new Web3.providers.HttpProvider("http://192.154.97.75:8545"));

var abi = JSON.parse(
'[ { "constant": true, "inputs": [ { "name": "", "type": "address" } ], "name": "tokensBurnt", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "name", "outputs": [ { "name": "", "type": "string" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "_to", "type": "address" } ], "name": "withdrawStudent", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "_to", "type": "address" }, { "name": "_pin", "type": "uint256" } ], "name": "validLogin", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "totalSupply", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "INITIAL_SUPPLY", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "decimals", "outputs": [ { "name": "", "type": "uint8" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "_to", "type": "address" }, { "name": "_pin", "type": "uint256" } ], "name": "newStudent", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "_value", "type": "uint256" } ], "name": "burn", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "_owner", "type": "address" } ], "name": "balanceOf", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [], "name": "renounceOwnership", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" } ], "name": "isStudent", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "index", "type": "uint256" } ], "name": "getStudentAddr", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "getNStudent", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "owner", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "symbol", "outputs": [ { "name": "", "type": "string" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "_to", "type": "address" }, { "name": "_value", "type": "uint256" } ], "name": "transfer", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "Assignment", "type": "uint256" } ], "name": "redeamToken", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "_to", "type": "address" }, { "name": "_pin", "type": "uint256" } ], "name": "updatePin", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "haveCoupons", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "_newOwner", "type": "address" } ], "name": "transferOwnership", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "inputs": [], "payable": false, "stateMutability": "nonpayable", "type": "constructor" }, { "anonymous": false, "inputs": [ { "indexed": false, "name": "Student", "type": "address" }, { "indexed": false, "name": "Assignment", "type": "uint256" } ], "name": "CouponUsed", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": false, "name": "Student", "type": "address" }, { "indexed": false, "name": "Success", "type": "bool" } ], "name": "PinUpdated", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "previousOwner", "type": "address" } ], "name": "OwnershipRenounced", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "previousOwner", "type": "address" }, { "indexed": true, "name": "newOwner", "type": "address" } ], "name": "OwnershipTransferred", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "burner", "type": "address" }, { "indexed": false, "name": "value", "type": "uint256" } ], "name": "Burn", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "name": "from", "type": "address" }, { "indexed": true, "name": "to", "type": "address" }, { "indexed": false, "name": "value", "type": "uint256" } ], "name": "Transfer", "type": "event" } ]'
);

var TheContract = web3.eth.contract(abi);

var contractInstance = TheContract.at(
	// "0xb6c75ddf7ae39046f60c3ed3b6d17a4682083fd2"	// .199
	"0xec849736cb876eb6050afe2886abb7f8a6582abc"	// new network
);

// var ownerAddr = "0x6ffba2d0f4c8fd7961f516af43c55fe2d56f6044";
var ownerAddr = "0x9a6446d642d76a3ac1baf6c6d8c1e5179c58d87f";

var _origPin = "";			// Xyzzy - pull sha1 from static/srp/Client/lib/sjcl - and use to store/compare.
var _LoginAddress = "";

// -----------------------------------------------------------------------------------------------------------------------------------------
// Test:
//		<a href="#" onclick="doPage('p_RedeamCoupon')" class="btn btn-primary">Redeam Coupon</a> <br>
// 2. Last Student Page:
//		<a href="#" onclick="doPage('p_TransferToken')" class="btn btn-primary">Transfer Token</a> <br>
// 3. Admin Pages:
//		<a href="#" onclick="doPage('a_NewStudent')" class="btn btn-primary my_admin">New Student</a> <br>
//		<a href="#" onclick="doPage('a_WithdrawStudent')" class="btn btn-primary my_admin">Student Withdraw</a> <br>
//		<a href="#" onclick="doPage('a_UpdatePin')" class="btn btn-primary my_admin">Update Pin</a> <br>
//		<a href="#" onclick="doPage('a_ActivityList')" class="btn btn-primary my_admin">Activity List (Redeamed Coupons)</a> <br>
// 4. Report - # of sutdents / pins / etc. -- Owner Only
// 1. change validLogin to return 0, err, 1 user, 2 admin
// 2. Put some margin around the buttons - so not touching.
// -----------------------------------------------------------------------------------------------------------------------------------------



// function redeamToken(uint256 Assignment) public {
// function haveCoupons() public view returns (uint256) {
// function getStudentAddr(uint index) public view onlyOwner returns (address) {

// Notes: https://blog.hellobloom.io/how-to-make-a-user-friendly-ethereum-dapp-5a7e5ea6df22
//        https://medium.com/blockchannel/how-to-save-your-ethereum-dapp-users-from-paying-gas-for-transactions-cfc665891ab4
//		  https://programtheblockchain.com/posts/2018/02/17/signing-and-verifying-messages-in-ethereum/
//		  https://youtu.be/6Gf_kRE4MJU -- MetaMask Video -- From: https://metamask.io/
//		  https://medium.com/@angellopozo/ethereum-signing-and-validating-13a2d7cb0ee3
//		  https://medium.com/blockchannel/life-cycle-of-an-ethereum-transaction-e5c66bae0f6e

function doTransferToken() {
	doPage('p_TransferToken');
}

function doTransferTokenSubmit() {
	var _to =$("#toTT").val();
	var _pw = $("#pwTT").val();
	if ( _to === _LoginAddress ) {
		$("#outErr").html("<br><h2>You can not transfer tokens to yourself.</h2><br>");
		return
	}
	console.log ( "Transfer: to=", _to, "pw=", _pw );
	web3.personal.unlockAccount(_LoginAddress, _pw, 10, function( err, result ) {
		console.log ( "unlockAccount: err=", err, "result=", result );
		contractInstance.transfer(_to, 1, {gas:300000, from: ownerAddr}, function(err, result) {
			console.log ( "err=", err, "result=", result );
			$("#outErr").html("Token Transfered To:"+_to);
			// xyzzy - conform - switch page. - display data on main.
			// xyzzy - call standard function for this.
		});
	});
}

// <a href="#" onclick="do('p_RedeamCoupon')" class="btn btn-primary">Redeam Coupon</a> <br>
// <a href="#" onclick="doRedeamCoupon()" class="btn btn-primary">Redeam Coupon</a> <br>
function doRedeamCoupon() {
	doPage('p_RedeamCoupon');
	// Assignment: <input type="text" id="assignment" name="assignment" class="input_data" /> <br>
	// Confirm Pin: <input type="text" id="pinRA" name="pinRA" class="input_data" /> <br>
	$("#assignment").val("");
	$("#pinRA").val("");
}

var mapAssignmentToNumber = {
	"A-01" 		:	1,
	"A-02" 		:	2,
	"A-03" 		:	3,
	"A-04" 		:	4,
	"A-05" 		:	5,
	"A-06" 		:	6,
	"A-paper" 	:	12,
	"A-07" 		:	7,
	"A-08" 		:	8,
	"A-09" 		:	9,
	"A-10" 		:	10,
	"A-11" 		:	11
};

function doRedeamCouponSubmit() {
	var _assignment = $('input[name=assignment]:checked').val(); 
	var _pw = $("#pwRA").val();
	// function redeamToken(uint256 Assignment) public {
	// TODO: function redeamToken(uint256 Assignment, uint256 Pin) public {
	console.log ( "DO-IT: assignment=", _assignment, "pw=", _pw, "ownerAddr=", ownerAddr);
	_assignment = mapAssignmentToNumber[_assignment];
	console.log ( "After Map: assignment=", _assignment, "pin=", _pin, "ownerAddr=", ownerAddr);

	web3.personal.unlockAccount(_LoginAddress, _pw, 10, function( err, result ) {
		console.log ( "unlockAccount: err=", err, "result=", result );
		contractInstance.redeamToken(_assignment, {gas:300000, from: ownerAddr}, function(err, result) {
			console.log ( "contract.redeamToken err=", err, "result=", result );
			$("#outErr").html("Token Applied to Assignment");
			// xyzzy - conform - switch page. - display data on main.
			// xyzzy - call standard function for this.
		});
	});

}

// function getNStudent() public view onlyOwner returns (uint256) {
function getNStudent() {
	contractInstance.getNStudent({gas:300000, from: ownerAddr}, function(err, result) {
		console.log ( "err=", err, "result=", result );
		$("#outErr").html(result.c[0]);
	});
}

// function newStudent(address _to) public onlyOwner {
function newStudent(_to,_pin) {
	// var to = web3.eth.accounts[0];
	contractInstance.newStudent(_to, _pin, {gas:300000, from: ownerAddr}, function(err, result) {
		console.log ( "err=", err, "result=", result );
		$("#outErr").html("created");
	});
}

var curPage = "p_Login";
var debug00 = true;
var isAdmin = true;	// xyzzy - TODO - get this info back from server.

$(document).ready(function() {
	doPage( curPage );
});

function errorMsg ( msg ) {
	console.log ( msg );
	$("#outErr").html("<br><h4>"+msg+"</h4><br>");
	// display the message for 4 sec, then clean up.
	setTimeout ( function() {
		$("#outErr").html("");
	}, 4000 );
}

// if login is valid, then swithc to main, else show error.
function doLogin() {
	var _to = $("#username").val();
	_LoginAddress = _to;
	if ( ! isValidAddress (_to) ) {
		errorMsg ( "Invalid Address" );
		// xyzzy - return
	}
	var _pin = $("#pin").val();
	_origPin = _pin;
	if ( !debug00 ) {
		if ( ! isValidPin (_pin) ) {
			errorMsg ( "Invalid Pin" );
			// xyzzy - return
		}
	}
	contractInstance.validLogin( _to, _pin, {gas:300000, from: ownerAddr}, function(err, result) {
		console.log ( "err=", err, "result=", result );
		if ( err ) {
			$("#outErr").html(err);
		} else {
			$("#outErr").html(result);
			if ( result ) {
				ownerAddr = _LoginAddress;
				doPage ( "p_Main" );
			} else {
				$("#outErr").html("<br><h4>Invalid Login</h4><br>Please Try Again<br>");
				doPage ( "p_Login" );
			}
		}
	});
}

function doShowBalance() {
	doPage( 'p_ShowBalance' );
	contractInstance.balanceOf(_LoginAddress, {gas:300000, from: ownerAddr}, function(err, result) {
		console.log ( "err=", err, "result=", result );
		$("#outToken").html(result.c[0]);
	});
}

function doShowNoOfStudetns() {
	doPage( 'p_ShowBalance' );
	contractInstance.getNStudent({gas:300000, from: ownerAddr}, function(err, result) {
		console.log ( "err=", err, "result=", result );
		$("#outToken").html(result.c[0]);
	});
}


function doLogout( ) {
	curPage = 'p_Logout';
    $(".my_page").hide();
    $("#"+curPage).show();
	$(".input_data").val("");
	$(".output_data").html("");
	_origPin = "";
	_LoginAddress = "";
	// ownerAddr = "0x6ffba2d0f4c8fd7961f516af43c55fe2d56f6044";
	ownerAddr = "0x9a6446d642d76a3ac1baf6c6d8c1e5179c58d87f";
	$(".id_radio").prop('checked', false);
	setTimeout ( function() {
		doPage ( 'p_Login' );
	}, 2000 );
}

function doPage( pn ) {
	curPage = pn;
    $(".my_page").hide();
    $("#"+curPage).show();
	$("#outErr").html("");		
	if ( !isAdmin ) {
		$(".my_admin").hide();
	}
}

function isValidAddress (address) {
    return /^0x[0-9a-f]{40}$/i.test(address);
}

function isValidPin (pin) {
    return /^[0-9]{8}$/i.test(pin);
}
