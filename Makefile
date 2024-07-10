BINARIES := h2d d2h
PLATFORMS := windows darwin linux
ARCHITECTURES := amd64 arm64

.PHONY: all clean

all: $(BINARIES)

$(BINARIES):
  @mkdir -p binary
  @for platform in $(PLATFORMS); do \
    mkdir -p binary/$$platform; \
    for arch in $(ARCHITECTURES); do \
      if [ $$platform = "windows" ]; then \
        GOOS=$$platform GOARCH=$$arch go build -o binary/$$platform/$@_$${arch}.exe ./cmd/$@; \
      else \
        GOOS=$$platform GOARCH=$$arch go build -o binary/$$platform/$@_$${arch} ./cmd/$@; \
      fi; \
    done; \
    if [ $$platform = "darwin" ]; then \
      lipo -create -output binary/$$platform/$@ binary/$$platform/$@_amd64 binary/$$platform/$@_arm64; \
      rm binary/$$platform/$@_amd64 binary/$$platform/$@_arm64; \
    fi; \
  done

clean:
  rm -rf binary
