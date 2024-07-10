BINARIES := h2d d2h
PLATFORMS := darwin linux windows
ARCHS := amd64 arm64


# Default install locations
INSTALL_DIR_LINUX := /usr/local/bin
INSTALL_DIR_MACOS := /usr/local/bin
INSTALL_DIR_WINDOWS := C:\Windows\System32

.PHONY: all clean test

all: $(BINARIES)

$(BINARIES):
	@echo "Building $@..."
	@for platform in $(PLATFORMS); do \
		for arch in $(ARCHS); do \
			output_dir="binary/$$platform"; \
			mkdir -p $$output_dir; \
			if [ $$platform = "windows" ]; then \
				GOOS=$$platform GOARCH=$$arch go build -o $$output_dir/$@_$${arch}.exe ./cmd/$@; \
			else \
				GOOS=$$platform GOARCH=$$arch go build -o $$output_dir/$@_$${arch} ./cmd/$@; \
			fi; \
		done; \
		if [ $$platform = "darwin" ]; then \
			lipo -create -output $$output_dir/$@ $$output_dir/$@_amd64 $$output_dir/$@_arm64; \
			rm $$output_dir/$@_amd64 $$output_dir/$@_arm64; \
		fi; \
	done

test:
	go test ./pkg/converter -v
	go test ./test -v

clean:
	rm -rf binary


info:
	@echo "Hexadecimal-Decimal Converter"
	@echo "Usage:"
	@echo "  h2d <hexadecimal number>  - Convert hex to decimal"
	@echo "  d2h <decimal number>      - Convert decimal to hex"
	@echo ""
	@echo "Examples:"
	@echo "  ./h2d FF  -> 255"
	@echo "  ./d2h 255 -> FF"
	@echo ""
	@echo "Build: make"
	@echo "Test:  make test"
	@echo "Clean: make clean"


install:
	@echo "Installing binaries..."
	@case "$$(uname -s)" in \
		Linux*) \
			sudo cp binary/linux/h2d_amd64 $(INSTALL_DIR_LINUX)/h2d && \
			sudo cp binary/linux/d2h_amd64 $(INSTALL_DIR_LINUX)/d2h && \
			echo "Installed to $(INSTALL_DIR_LINUX)" ;; \
		Darwin*) \
			sudo cp binary/darwin/h2d $(INSTALL_DIR_MACOS)/h2d && \
			sudo cp binary/darwin/d2h $(INSTALL_DIR_MACOS)/d2h && \
			echo "Installed to $(INSTALL_DIR_MACOS)" ;; \
		MINGW*|MSYS*|CYGWIN*) \
			cp binary/windows/h2d_amd64.exe $(INSTALL_DIR_WINDOWS)/h2d.exe && \
			cp binary/windows/d2h_amd64.exe $(INSTALL_DIR_WINDOWS)/d2h.exe && \
			echo "Installed to $(INSTALL_DIR_WINDOWS)" ;; \
		*) \
			echo "Unsupported operating system, copy executables manually or add executables to PATH" ;; \
	esac
