build:
	go build -C cmd/rss-aggregator/ -o ../../bin/rss-aggregator

run: build
	./bin/rss-aggregator