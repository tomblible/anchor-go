build:
	go build 

pump_curve:build
	./anchor-go  -src=./idl/dex/pump_curve.json -pkg=pump_curve -dst=./generated/prd/pump_curve

pump_amm:build
	./anchor-go  -src=./idl/dex/pump_amm.json -pkg=pump_amm -dst=./generated/prd/pump_amm

raydium_amm:build
	anchor idl convert ./idl/dex/raydium_amm.json >./idl/dex/raydium_amm_new.json
	./anchor-go -type-id=uint8 -src=./idl/dex/raydium_amm_new.json -pkg=raydium_amm -dst=./generated/prd/raydium_amm

raydium_clmm:build
	anchor idl convert ./idl/dex/raydium_clmm.json >./idl/dex/raydium_clmm_new.json
	./anchor-go -src=./idl/dex/raydium_clmm_new.json -pkg=raydium_clmm -dst=./generated/prd/raydium_clmm

raydium_cpmm:build
	anchor idl convert ./idl/dex/raydium_cpmm.json >./idl/dex/raydium_cpmm_new.json
	./anchor-go -src=./idl/dex/raydium_cpmm_new.json -pkg=raydium_cpmm -dst=./generated/prd/raydium_cpmm

whirlpool:build
	# anchor idl convert ./idl/dex/whirlpool.json >./idl/dex/whirlpool_new.json
	./anchor-go -src=./idl/dex/whirlpool.json -pkg=orca_whirlpool -dst=./generated/prd/whirlpool

meteora_dynamic_bonding_curve:build
	./anchor-go -src=./idl/dex/meteora_dynamic_bonding_curve.json -pkg=meteora_curve -dst=./generated/prd/meteora_dynamic_bonding_curve

meteora_dynamic_bonding_curve_v2:build
	./anchor-go -src=./idl/dex/meteora_dynamic_bonding_curve_v2.json -pkg=meteora_curve -dst=./generated/prd/meteora_curve


meteora_damm:build
	./anchor-go -src=./idl/dex/meteora_damm_v2.json -pkg=meteora_damm -dst=./generated/prd/meteora_damm

meteora_dlmm:build
	anchor idl convert ./idl/dex/meteora_dlmm.json >./idl/dex/meteora_dlmm_new.json
	./anchor-go -src=./idl/dex/meteora_dlmm_new.json -pkg=meteora_dlmm -dst=./generated/prd/meteora_dlmm

meteora_pools:build
	anchor idl convert ./idl/dex/meteora_pools.json >./idl/dex/meteora_pools_new.json
	./anchor-go -src=./idl/dex/meteora_pools_new.json -pkg=meteora_pools -dst=./generated/prd/meteora_pools

meteora_vault:build
	anchor idl convert ./idl/dex/meteora_vault.json >./idl/dex/meteora_vault_new.json
	./anchor-go -src=./idl/dex/meteora_vault_new.json -pkg=meteora_vault -dst=./generated/prd/meteora_vault

raydium_launchpad:build
	./anchor-go -src=./idl/dex/raydium_launchpad.json -pkg=raydium_launchpad -dst=./generated/prd/raydium_launchpad

moonit:build
	anchor idl convert ./idl/dex/moonit.json >./idl/dex/moonit_new.json
	./anchor-go -src=./idl/dex/moonit_new.json -pkg=moonit -dst=./generated/prd/moonit

ego_one:build
	./anchor-go -src=./idl/dex/ego_one.json -pkg=ego_one -dst=./generated/prd/ego_one

vr:build
	./anchor-go -src=./idl/dex/vr.json -pkg=vr -dst=./generated/fyd/vr
	
token2022:build
	./anchor-go -type-id=uint8 -src=./idl/solana/spl/token2022.json -pkg=token2022 -dst=./generated/prd/token2022

hpnf:build
	anchor idl convert ./idl/dex/hpnf.json >./idl/dex/hpnf_new.json
	./anchor-go -src=./idl/dex/hpnf.json -pkg=hpnf -dst=./generated/prd/hpnf

saros_dlmm:build
	./anchor-go -src=./idl/dex/saros_dlmm.json -pkg=saros_dlmm -dst=./generated/prd/saros_dlmm

7HrX:build
	./anchor-go -src=./idl/dex/hrx.json -pkg=hrx -dst=./generated/prd/hrx

wavebreak:build
	./anchor-go -type-id=uint8 -src=./idl/dex/wavebreak.json -pkg=wavebreak -dst=./generated/prd/wavebreak

mastermind:build
	./anchor-go -src=./idl/dex/mastermind.json -pkg=mastermind -dst=./generated/prd/mastermind

heaven:build
	./anchor-go -src=./idl/dex/heaven.json -pkg=heaven -dst=./generated/prd/heaven


upgrade-anchor:
	avm install 0.30.1 && anchor --version 
	
all: pump_curve pump_amm raydium_amm raydium_clmm raydium_cpmm whirlpool meteora_dynamic_bonding_curve meteora_damm_v2 meteora_dlmm meteora_pools meteora_vault raydium_launchpad moonit

convert:
	# anchor idl convert ./test/old.json >./test/new.json
	anchor idl convert ./test/raydium.json >./test/raydium_new.json