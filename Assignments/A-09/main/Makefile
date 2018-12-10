
all:
	go build

.PHONY: test test001 test_sf1 load_c

test: test001 test_sf0 test_sf1 load_c


load_c:
	go build
	./main --load-contract --contract-file ./testdata/sample01.js \
		--key-file ./testdata/UTC--2018-02-15T19-57-35.216297214Z--6ffba2d0f4c8fd7961f516af43c55fe2d56f6044 \
		--key-file-password password

genesis:
	go build
	mkdir -p data
	rm -f ./data/*.json
	./main --create-genesis


#Test Order
#
#	Test 000
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		3. ** Verify index is correct **
#		4. ** Verify Tx/Block is correct **
#					"Data": "{\"test01\":\"-setup-\",\"test02\":\"-now-\"}",
c000: genesis
	./main --load-contract --contract-file ./testdata/sample01.js \
		--key-file ./testdata/UTC--2018-02-15T19-57-35.216297214Z--6ffba2d0f4c8fd7961f516af43c55fe2d56f6044 \
		--key-file-password password
	grep 'test01.*-setup-.*-now-' ./data/1cc65a7db6f16e6d9ff5c5850fc5e337cfb36da600ea2047c628a496d57b137d.json


#		0. Test(go) for these 3 routines (ct.go)
c001a:


#
#
#	Test 001
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		3. Pull back address of contract (CLI) (xyzzy-impleent)
#			+0. /api/getAddressFor - owner/name -> most recent address for a contrat
#			+0. /api/getAddressListFor - owner/name -> all addresss for a contrat
#		4. ** Verify is correct **
#	Test 002
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		2. Load a 2nd contract (Same Name/Dif body)
#		3. Pull back address of contract (CLI)
#		4. ** Verify is correct **
#	Test 003
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		2. Load a 2nd contract (Same Name/Same body)
#		3. Pull back address of contract (CLI)
#		4. ** Verify is correct **
#	Test 004
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		2. Load a 2nd contract (Same Name/Dif body)
#		3. Pull back address [list] for contract (CLI) (xyzzy-impleent)
#		4. ** Verify is correct **
#		5. ** Constructed public/private data **
#	Test 005
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		2. Load a 2nd contract (Same Name/Dif body)
#		3. Call a method in contract (xyzzy-impleent)
#		3. Call a view in contract (xyzzy-impleent)
#		4. ** Verify call results **
#		4. ** Verify output **
#		4. ** Verify trace **
#		4. ** Verify log **
#	Test 006
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		2. Load a 2nd contract (Same Name/Dif body)
#		3. Call a method in contract ( 10 times, inc a counter ) 
#		3. Call a view in contract - to get back value of counter
#		4. ** Verify call results **
#		4. ** Verify output **
#		4. ** Verify trace **
#		4. ** Verify log **
#	Test 007
#		0. Cleanup
#		1. genesis-init system
#		2. Load a contract
#		3. Verify other API calls	
#			http.HandleFunc("/api/get-contract-data", respHandlerPlaceholder)                        // Get the data for a contract - reports address
#			http.HandleFunc("/api/get-contract-trace", respHandlerPlaceholder)                       // Get the input/output trace for calls to contract
#			http.HandleFunc("/api/get-contract-log", respHandlerPlaceholder)                         // Get the event log for contract
#			http.HandleFunc("/api/set-contract-trace", respHandlerPlaceholder)                       // Turn on/off tracing of contract
#			http.HandleFunc("/api/get-contract-output", respHandlerPlaceholder)                      // Return the "output" from the contract (debug info)
#			http.HandleFunc("/api/get-contract-private-data", respHandlerPlaceholder)                // Get the contract private data (encrypted); reports address
#			http.HandleFunc("/api/get-contract-info", respHandlerPlaceholder)                        // Return list of owners and addresses for this named contracts
#		+0. Fetch back the contract
#			http.HandleFunc("/api/get-contract-source", respHandlerPlaceholder)                      // Get source of named contract at an address
#






T1=29e68e530d0718dd01759e8c9a5276d20687bc5ec23e60dce063c2b97ba6b04f.json

start_server:
	go build
	echo Starting server on port 9191
	./main --server 192.168.0.157:9191 &

test001:
	go build
	mkdir -p data
	rm -f ./data/*.json
	rm -f data/${T1}
	rm -f data/index.json
	rm -f data/TestWriteBlock.json
	./main --cfg test-genesis.json --create-genesis
	diff data/index.json ref
	diff data/${T1} ref
	./main --test-read-block
	./main --test-write-block
	diff data/TestWriteBlock.json ref
	@echo 
	@echo PASS

test_sf0:
	go build
	mkdir -p data
	rm -f ./data/*.json
	mkdir -p ./out
	./main --create-genesis

# You can run this test multiple times after 'test_sf0' has setup the ./data directory.
test_sf1:
	go build
	mkdir -p ./out
	./main --show-balance 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 >./out/bal885.bef.out
	./main --show-balance 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 >./out/bal6E0.bef.out
	./main --test-send-funds 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 33 x x x "Tx Memo"
	./main --show-balance 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 >./out/bal885.aft.out
	./main --show-balance 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 >./out/bal6E0.aft.out
	./validateBalance.sh ./out/bal885.bef.out -33 ./out/bal885.aft.out
	./validateBalance.sh ./out/bal6E0.bef.out  33 ./out/bal6E0.aft.out
	@echo 
	@echo PASS

x:
	./validateBalance.sh ./out/bal885.bef.out -33 ./out/bal885.aft.out

y:
	./validateBalance.sh ./out/bal6E0.bef.out  33 ./out/bal6E0.aft.out

test_create_genesis:
	go build
	rm -f ./data/*.json
	./main --create-genesis

test_list_acct:
	go build
	rm -f ./data/*.json
	./main --create-genesis
	./main --list-accounts

test_show_balance:
	go build
	rm -f ./data/*.json
	./main --create-genesis
	./main --show-balance 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 

test_send_funds:
	go build
	rm -f ./data/*.json
	mkdir -p ./out
	./main --create-genesis
	./main --test-send-funds 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 33 x x x "Memo 02"
	./main --show-balance 0xdb180da9a8982c7bb75ca40039f959cb959c62e8 >./out/tsf.out
	./validateBal2.sh ./out/tsf.out 500200
	@echo 
	@echo PASS

w:
	go build
	./main --show-balance 0xdb180da9a8982c7bb75ca40039f959cb959c62e8
