
all:
	go build

.PHONY: test test001 test_sf1

test: test001 test_sf1

test001:
	go build
	mkdir -p data
	rm -f data/136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e.json 
	rm -f data/index.json
	rm -f data/TestWriteBlock.json
	./main --cfg test-genesis.json --create-genesis
	diff data/index.json ref
	diff data/136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e.json ref
	./main --test-read-block
	./main --test-write-block
	diff data/TestWriteBlock.json ref
	@echo 
	@echo PASS

test_sf1:
	go build
	mkdir -p ./out
	./main --show-balance 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 >./out/bal885.bef.out
	./main --show-balance 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 >./out/bal6E0.bef.out
	./main --test-send-funds 0x885765a2fcfB72e68D82D656C6411b7D27BacDD7 0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1 33 x x x 
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

test_list_acct:
	go build
	rm -f ./data/*.json
	./main --create-genesis
	./main --list-accounts

