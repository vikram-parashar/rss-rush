build:
	@go build -o bin/rssrush

run: build
	@./bin/rssrush
