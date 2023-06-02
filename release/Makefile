.PHONY: windows linux macos copy gpg clean

all: windows linux macos copy

linux:
	make -C linux all

macos:
	make -C macos all

windows:
	make -C windows all

copy:
	rm -rf installers
	mkdir installers
	cp macos/build/detahard-bridge*.pkg* installers
	cp windows/build/detahard-bridge*.exe* installers
	cp linux/build/*.deb linux/build/*.rpm installers
	rm -rf binaries
	mkdir binaries
	cp macos/build/detahardd binaries/detahardd-darwin-universal
	cp windows/build/detahardd-32b.exe binaries/detahardd-windows-386.exe
	cp windows/build/detahardd-64b.exe binaries/detahardd-windows-amd64.exe
	cp linux/build/detahardd-linux-* binaries/

clean:
	rm -rf installers binaries macos/build linux/build windows/build
