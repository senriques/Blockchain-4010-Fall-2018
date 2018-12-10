//$$spec
// ContractName Test01Contract
// Constructor initTest01Contract
// Upgrader upgradeTest01Contract
// Methods setTest01 setTest02 getTest01 getTest02
// Views nTest
//$$end-spec

function initTest01Contract( msg, data, cnst, plist ) {
	data.test01 = "-setup-";
	data.test02 = "-now-";
	cnst.nTest = 2;
	return data;
}

function upgradeTest01Contract( msg, data, cnst, plist ) {
	data.test01 = "-updated-";
	data.test02 = "-now-now-";
	cnst.nTest = 2;
	return data;
}

// plist error - is array of values
function setTest01 ( msg, data, nv ) {	
	data.test01 = nv;
	return data;
}

// plist error - is array of values
function setTest02 ( msg, data, nv ) {
	data.test02 = nv;
	return data;
}

// plist error - is array of values
function getTest01 ( msg, data ) {
	returnData ( data.test01 );
	return data;
}

// plist error - is array of values
function getTest02 ( msg, data ) {
	returnData ( data.test02 );
	return data;
}

// plist error - is array of values
function nTest ( msg, data ) {
	returnData ( data.nTest );
}

