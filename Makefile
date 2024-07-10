BINARIES := h2d d2h
PLATFORMS := darwin linux windows
ARCHS := amd64 arm64

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