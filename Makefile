SOURCES:=$(shell find . -type f -name '*.go')

go-prompt: $(SOURCES)
	go build -o $@ .
	strip $@

install: go-prompt
	install -Dm755 $< ~/.local/bin/$<
