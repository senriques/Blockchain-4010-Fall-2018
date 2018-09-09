require('babel-register');
require('babel-polyfill');

module.exports = {
	networks: {
		development: {
			host: "127.0.0.1",
			port: 9545,
			network_id: "*" // Match any network id
		}
		, dApp: {
			host: "192.154.97.75",
			port: 8545,
			network_id: "*",
		    gas: 4712388,
			from: "0x9a6446d642d76a3ac1baf6c6d8c1e5179c58d87f"
		}
	}
};
