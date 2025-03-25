build:
	go build 

swap-proxy:build
	./anchor-go -type-id=uint8 -src=./swap_proxy.json -pkg=swap_proxy -dst=./generated/swap_proxy

temp:build
	./anchor-go -type-id=uint8 -src=./temp.json -pkg=temp -dst=./generated/temp

run:
	go run ./ -src=D:\project\anchor-go\idl\temp.json



test:build
	./anchor-go -type-id=uint8 -src=./test/swap_proxy.json -pkg=swap_proxy -dst=./generated/test/swap_proxy
	./anchor-go -type-id=uint8 -src=./test/temp.json -pkg=temp -dst=./generated/test/temp
	./anchor-go -type-id=uint8 -src=./test/cookMeme.json -pkg=cookMeme -dst=./generated/test/cookMeme
	./anchor-go -type-id=uint8 -src=./test/jup_new.json -pkg=jupiter -dst=./generated/test/jupiter
	./anchor-go -type-id=uint8 -src=./test/stabble_new.json -pkg=stabble -dst=./generated/test/stabble
	./anchor-go -type-id=uint8 -src=./test/pump_amm.json -pkg=pump_amm -dst=./generated/test/pump_amm
	./anchor-go -type-id=uint8 -src=./test/lb_clmm.json -pkg=lb_clmm -dst=./generated/test/lb_clmm

upgrade-anchor:
	avm install 0.30.1 && anchor --version 

convert:
	anchor idl convert ./test/old.json >./test/new.json