build:
	go build && rm -rf ./generated
	./anchor-go -type-id=uint8 -src=./swap_proxy.json -pkg=swap_proxy -dst=./generated/swap_proxy