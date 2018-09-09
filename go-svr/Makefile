all: go-svr \
	./contracts/UW4010_sol_UW4010.abi 

linux: 
	GOOS=linux go build -o go-svr.linux .

go-svr: go-server.go lib.go
	go build 

install:
	cp go-svr /home/pschlump/bin


dApp_migrate:
	./unlock-9a6-1.sh
	truffle migrate --reset --network dApp | tee m.out
	upd-gcall <m.out
	save-old.sh

lint:
	solium --file contracts/UW4010.sol 

./contracts/UW4010_sol_UW4010.abi: ./contracts/UW4010.sol
	( cd contracts ; solc openzeppelin-solidity/=../openzeppelin-solidity/ --allow-paths . --overwrite --abi -o . UW4010.sol ; mv UW4010.abi UW4010_sol_UW4010.abi )
	( cd contracts ; /Users/corwin/bin/check-json-syntax -p <UW4010_sol_UW4010.abi >,tmp ; mv ,tmp UW4010_sol_UW4010.abi )
	( cd contracts ; solc openzeppelin-solidity/=../openzeppelin-solidity/ --allow-paths . --overwrite --bin -o . UW4010.sol ; mv UW4010.bin UW4010_sol_UW4010.bin )
	( cd contracts ; solc openzeppelin-solidity/=../openzeppelin-solidity/ --allow-paths . --overwrite --bin-runtime -o . UW4010.sol )
	( cd contracts ; abigen --abi UW4010_sol_UW4010.abi --pkg UW4010 --out UW4010.go )
	mkdir -p ./lib/UW4010
	cp ./contracts/UW4010.go ./lib/UW4010
	cp ./contracts/UW4010_sol_UW4010.abi ./abi
	( cd ./lib/UW4010 ; go build )

run_test:
	date
	/Users/corwin/Projects/eth158/unlock-c2a.sh
	truffle test --network local

