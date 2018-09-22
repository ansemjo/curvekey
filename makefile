# targets that are not actual files
.PHONY : build clean mkrelease-prepare mkrelease mkrelease-finish release

GOFILES := $(shell find -type f -name '*.go') go.mod go.sum
GOFLAGS = -ldflags="-s -w"

# build static binary w/o debugging symbols
build : $(GOFILES)
	go build $(GOFLAGS)
	command -v upx >/dev/null && upx curvekey

# clean anything not tracked by git
clean :
	git clean -dfx

# use host os/arch and current directory by default
OS    := $(shell go env GOHOSTOS)
ARCH  := $(shell go env GOHOSTARCH)

RELEASEDIR := $(PWD)

# makerelease targets for reproducible builds, ansemjo/makerelease
mkrelease-prepare:
	go mod download

EXT := $(if $(findstring windows,$(OS)),.exe)
mkrelease: $(GOFILES)
	GOOS=$(OS) GOARCH=$(ARCH) go build $(GOFLAGS) -o $(RELEASEDIR)/curvekey-$(OS)-$(ARCH)$(EXT)

mkrelease-finish:
	upx $(RELEASEDIR)/* || true
	printf "# built with %s in %s\n" "$$MKR_VERSION" "$$MKR_IMAGE" > $(RELEASEDIR)/SHA256SUMS
	cd $(RELEASEDIR) && sha256sum *-*-* | tee -a SHA256SUMS

release:
	git archive --prefix=./ HEAD | mkr rl
