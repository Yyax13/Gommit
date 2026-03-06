.PHONY: build install clean uninstall
.SILENT: build install clean uninstall

build:
	go build -ldflags="-s -w" -o build/gommit .
	
install: 
	sudo cp build/gommit /usr/bin/gommit

clean:
	rm -rf build/gommit

uninstall:
	sudo rm -rf /usr/bin/gommit