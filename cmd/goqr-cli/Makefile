
all:
	go build

test1:
	go build
	mkdir -p ./out
	./goqr-cli -url "A Test Case" >out/,out
	diff out/,out ref/test1.out
	@echo PASS

