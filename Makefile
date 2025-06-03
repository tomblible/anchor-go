build:
	go build 

pump_curve:build
	./anchor-go  -src=./dex/fyd/pump_curve.json -pkg=pump_curve -dst=./generated/prd/pump_curve

pump_amm:build
	./anchor-go  -src=./dex/fyd/pump_amm.json -pkg=pump_amm -dst=./generated/prd/pump_amm

raydium_amm:build
	anchor idl convert ./dex/fyd/raydium_amm.json >./dex/fyd/raydium_amm_new.json
	./anchor-go -type-id=uint8 -src=./dex/fyd/raydium_amm_new.json -pkg=raydium_amm -dst=./generated/prd/raydium_amm

raydium_clmm:build
	anchor idl convert ./dex/fyd/raydium_clmm.json >./dex/fyd/raydium_clmm_new.json
	./anchor-go -src=./dex/fyd/raydium_clmm_new.json -pkg=raydium_clmm -dst=./generated/prd/raydium_clmm

raydium_cpmm:build
	anchor idl convert ./dex/fyd/raydium_cpmm.json >./dex/fyd/raydium_cpmm_new.json
	./anchor-go -src=./dex/fyd/raydium_cpmm_new.json -pkg=raydium_cpmm -dst=./generated/prd/raydium_cpmm

whirlpool:build
	anchor idl convert ./dex/fyd/whirlpool.json >./dex/fyd/whirlpool_new.json
	./anchor-go -src=./dex/fyd/whirlpool_new.json -pkg=whirlpool -dst=./generated/prd/whirlpool

meteora_dynamic_bonding_curve:build
	./anchor-go -src=./dex/fyd/meteora_dynamic_bonding_curve.json -pkg=meteora_dynamic_bonding_curve -dst=./generated/prd/meteora_dynamic_bonding_curve

meteora_damm_v2:build
	./anchor-go -src=./dex/fyd/meteora_damm_v2.json -pkg=meteora_damm_v2 -dst=./generated/prd/meteora_damm_v2

meteora_dlmm:build
	anchor idl convert ./dex/fyd/meteora_dlmm.json >./dex/fyd/meteora_dlmm_new.json
	./anchor-go -src=./dex/fyd/meteora_dlmm_new.json -pkg=meteora_dlmm -dst=./generated/prd/meteora_dlmm

meteora_pools:build
	anchor idl convert ./dex/fyd/meteora_pools.json >./dex/fyd/meteora_pools_new.json
	./anchor-go -src=./dex/fyd/meteora_pools_new.json -pkg=meteora_pools -dst=./generated/prd/meteora_pools

meteora_vault:build
	anchor idl convert ./dex/fyd/meteora_vault.json >./dex/fyd/meteora_vault_new.json
	./anchor-go -src=./dex/fyd/meteora_vault_new.json -pkg=meteora_vault -dst=./generated/prd/meteora_vault

upgrade-anchor:
	avm install 0.30.1 && anchor --version 
	
all: pump_curve pump_amm raydium_amm raydium_clmm raydium_cpmm whirlpool meteora_dynamic_bonding_curve meteora_damm_v2 meteora_dlmm meteora_pools meteora_vault

convert:
	# anchor idl convert ./test/old.json >./test/new.json
	anchor idl convert ./test/raydium.json >./test/raydium_new.json