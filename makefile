build_x:
	go build -o build/ngo .

build_win:
	rm -rf build/ngo.exe
	GOOS=windows go build -o build/ngo.exe .