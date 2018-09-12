pragma solidity ^0.4.24;

// import "../token/ERC20/StandardToken.sol";
import "openzeppelin-solidity/contracts/token/ERC20/BurnableToken.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";

/**
 * @title SimpleToken
 * @dev Very simple ERC20 Token example, where all tokens are pre-assigned to the creator.
 * Note they can later distribute these tokens as they wish using `transfer` and other
 * `StandardToken` functions.
 */
contract UW4010 is BurnableToken, Ownable {

	string public constant name = "UWBlockChain4010"; 
	string public constant symbol = "UW4010"; 
	uint8 public constant decimals = 0; 

	uint256 public constant INITIAL_SUPPLY = 10000000000;

	mapping(address => uint256) public isStudent;
	mapping(address => uint256) public tokensBurnt;
	mapping(address => uint256) public balance1;
	address[] private listOfStudentAddrs;
	uint256[] private listOfStudentPin;

	event CouponUsed(address Student, uint256 Assignment);
	event PinUpdated(address Student, bool Success);

	/**
	 * @dev Constructor that gives msg.sender all of existing tokens.
	 */
	constructor() public {
		totalSupply_ = INITIAL_SUPPLY;
		balances[msg.sender] = INITIAL_SUPPLY;
		balance1[msg.sender] = INITIAL_SUPPLY;
		emit Transfer(0x0, msg.sender, INITIAL_SUPPLY);
	}

	/**
	 * @dev Use a token to turn in a homework late.
	 */
	// TODO: function redeamToken(uint256 Assignment, uint256 Pin) public {
	function redeamToken(uint256 Assignment) public {
		tokensBurnt[msg.sender]++;
		burn(1);
		emit CouponUsed(msg.sender, Assignment);
	}

	/**
	 * @dev Create a new student with 3 tokens.
	 */
	function newStudent(address _to, uint256 _pin) public onlyOwner {
		// require(isStudent[_to] <= 0, "Already a student.");		// fix this!!!! FIXME
		tokensBurnt[_to] = 0;
		listOfStudentAddrs.push(_to);
		listOfStudentPin.push(_pin);
		isStudent[_to] = listOfStudentAddrs.length;
		// transfer(_to, 3);
    // require(_to != address(0));
    // require(_value <= balances[msg.sender]);
    balances[msg.sender] = balances[msg.sender] - 3;
    balance1[msg.sender] = balance1[msg.sender] - 3;
    balances[_to] = balances[_to] + 3;
    balance1[_to] = 3;
    emit Transfer(msg.sender, _to, 3);
	}

	function convToAddress(address _x) public view returns(address){
		return(_x);
	}

	function newStudentCoupons(address _to, uint256 _pin, uint256 _value) public onlyOwner {
		// require(isStudent[_to] <= 0, "Already a student.");		// fix this!!!! FIXME
		tokensBurnt[_to] = 0;
		listOfStudentAddrs.push(_to);
		listOfStudentPin.push(_pin);
		isStudent[_to] = listOfStudentAddrs.length;
		// transfer(_to, _coupons);
    require(_to != address(0));
    require(_value <= balances[msg.sender]);

    balances[msg.sender] = balances[msg.sender].sub(_value);
    balances[_to] = balances[_to].add(_value);
    emit Transfer(msg.sender, _to, _value);
	}

	function haveCouponsStudent(address _to) public view returns(uint256) {
    	return balances[_to];
	}

	/**
	 * @dev TODO
	 */
	function withdrawStudent(address _to) public onlyOwner {
		isStudent[_to] = 0;
		tokensBurnt[_to] = 0;
		uint256 _value;
		_value = balanceOf(_to);
		balances[_to] = 0;
		totalSupply_ = totalSupply_.sub(_value);
	}

	/**
	 * @dev TODO
	 */
	function updatePin(address _to, uint256 _pin) public onlyOwner returns(bool) {
		uint256 pos = isStudent[_to]-1;
		if ( pos >= 0 && pos < listOfStudentPin.length ) {
			listOfStudentPin[pos] = _pin;
			emit PinUpdated(_to,true);
			return(true);
		}
		emit PinUpdated(_to,false);
		return(false);
	}

	/**
	 * @dev Return true if the login pin matches.
	 */
	function validLogin(address _to, uint256 _pin) public view returns(bool) {
		uint256 pos = isStudent[_to]-1;
		if ( pos >= 0 && pos < listOfStudentPin.length ) {
			if ( listOfStudentPin[pos] == _pin ) {
				return(true);
			}
		}
		return(false);
	}

	/**
	 * @dev Report on the number of tokens that a student has left.
	 */
	function haveCoupons() public view returns (uint256) {
		return balanceOf(msg.sender);
	}

	/**
	 * @dev getStudentAddr will return the address of a single student in the list.
	 * @param index Index to return.
	 */
	function getStudentAddr(uint index) public view onlyOwner returns (address) {
		require(index >= 0 && index <= listOfStudentAddrs.length, "Index out of range.");
		return listOfStudentAddrs[index];
	}

	/**
	 * @dev Get number of students in list.
	 */
	function getNStudent() public view onlyOwner returns (uint256) {
		return listOfStudentAddrs.length;
	}

}
