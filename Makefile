all: 
	mkdir out
	go build -o out/ .

clean: 
	rm -rf out/