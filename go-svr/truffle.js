require('babel-register');
require('babel-polyfill');

module.exports = {
	networks: {
		development: {
			host: "127.0.0.1",
			port: 9545,
			network_id: "*" // Match any network id
		}
		, test3: {
			host: "192.168.0.199",
			port: 8545,
			network_id: "*",
		    gas: 4712388,
			from: "0x6ffba2d0f4c8fd7961f516af43c55fe2d56f6044"
		}
		, dApp: {
			host: "192.154.97.75",
			port: 28545,
			network_id: "*",
		    gas: 4712388,
			from: "0x9a6446d642d76a3ac1baf6c6d8c1e5179c58d87f"
		}
	}
};
