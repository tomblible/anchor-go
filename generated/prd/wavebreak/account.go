package wavebreak

import (
	ag_binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type MintConfig struct {
	Bump                          uint8
	QuoteMint                     solana.PublicKey
	CreateRequiresPermission      bool
	CreatePermissionBitmap        [32]uint8
	DefaultBuyRequiresPermission  bool
	DefaultBuyPermissionBitmap    [32]uint8
	DefaultSellRequiresPermission bool
	DefaultSellPermissionBitmap   [32]uint8
	Padding1                      [3]uint8
	DefaultCreatorReward          uint64
	DefaultGraduationReward       uint64
	DefaultGraduationTarget       uint64
	DefaultMaxBuyAmount           uint64
	DefaultMaxSellAmount          uint64
	DefaultStartPrice             ag_binary.Uint128
	DefaultEndPrice               ag_binary.Uint128
	DefaultControlPoints          [4]uint16
	DefaultSwapFeeBps             uint16
	DefaultQuoteFeeBps            uint16
	DefaultBaseFeeBps             uint16
	Padding2                      [1826]uint8
}

func (obj MintConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `QuoteMint` param:
	err = encoder.Encode(obj.QuoteMint)
	if err != nil {
		return err
	}
	// Serialize `CreateRequiresPermission` param:
	err = encoder.Encode(obj.CreateRequiresPermission)
	if err != nil {
		return err
	}
	// Serialize `CreatePermissionBitmap` param:
	err = encoder.Encode(obj.CreatePermissionBitmap)
	if err != nil {
		return err
	}
	// Serialize `DefaultBuyRequiresPermission` param:
	err = encoder.Encode(obj.DefaultBuyRequiresPermission)
	if err != nil {
		return err
	}
	// Serialize `DefaultBuyPermissionBitmap` param:
	err = encoder.Encode(obj.DefaultBuyPermissionBitmap)
	if err != nil {
		return err
	}
	// Serialize `DefaultSellRequiresPermission` param:
	err = encoder.Encode(obj.DefaultSellRequiresPermission)
	if err != nil {
		return err
	}
	// Serialize `DefaultSellPermissionBitmap` param:
	err = encoder.Encode(obj.DefaultSellPermissionBitmap)
	if err != nil {
		return err
	}
	// Serialize `Padding1` param:
	err = encoder.Encode(obj.Padding1)
	if err != nil {
		return err
	}
	// Serialize `DefaultCreatorReward` param:
	err = encoder.Encode(obj.DefaultCreatorReward)
	if err != nil {
		return err
	}
	// Serialize `DefaultGraduationReward` param:
	err = encoder.Encode(obj.DefaultGraduationReward)
	if err != nil {
		return err
	}
	// Serialize `DefaultGraduationTarget` param:
	err = encoder.Encode(obj.DefaultGraduationTarget)
	if err != nil {
		return err
	}
	// Serialize `DefaultMaxBuyAmount` param:
	err = encoder.Encode(obj.DefaultMaxBuyAmount)
	if err != nil {
		return err
	}
	// Serialize `DefaultMaxSellAmount` param:
	err = encoder.Encode(obj.DefaultMaxSellAmount)
	if err != nil {
		return err
	}
	// Serialize `DefaultStartPrice` param:
	err = encoder.Encode(obj.DefaultStartPrice)
	if err != nil {
		return err
	}
	// Serialize `DefaultEndPrice` param:
	err = encoder.Encode(obj.DefaultEndPrice)
	if err != nil {
		return err
	}
	// Serialize `DefaultControlPoints` param:
	err = encoder.Encode(obj.DefaultControlPoints)
	if err != nil {
		return err
	}
	// Serialize `DefaultSwapFeeBps` param:
	err = encoder.Encode(obj.DefaultSwapFeeBps)
	if err != nil {
		return err
	}
	// Serialize `DefaultQuoteFeeBps` param:
	err = encoder.Encode(obj.DefaultQuoteFeeBps)
	if err != nil {
		return err
	}
	// Serialize `DefaultBaseFeeBps` param:
	err = encoder.Encode(obj.DefaultBaseFeeBps)
	if err != nil {
		return err
	}
	// Serialize `Padding2` param:
	err = encoder.Encode(obj.Padding2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MintConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `QuoteMint`:
	err = decoder.Decode(&obj.QuoteMint)
	if err != nil {
		return err
	}
	// Deserialize `CreateRequiresPermission`:
	err = decoder.Decode(&obj.CreateRequiresPermission)
	if err != nil {
		return err
	}
	// Deserialize `CreatePermissionBitmap`:
	err = decoder.Decode(&obj.CreatePermissionBitmap)
	if err != nil {
		return err
	}
	// Deserialize `DefaultBuyRequiresPermission`:
	err = decoder.Decode(&obj.DefaultBuyRequiresPermission)
	if err != nil {
		return err
	}
	// Deserialize `DefaultBuyPermissionBitmap`:
	err = decoder.Decode(&obj.DefaultBuyPermissionBitmap)
	if err != nil {
		return err
	}
	// Deserialize `DefaultSellRequiresPermission`:
	err = decoder.Decode(&obj.DefaultSellRequiresPermission)
	if err != nil {
		return err
	}
	// Deserialize `DefaultSellPermissionBitmap`:
	err = decoder.Decode(&obj.DefaultSellPermissionBitmap)
	if err != nil {
		return err
	}
	// Deserialize `Padding1`:
	err = decoder.Decode(&obj.Padding1)
	if err != nil {
		return err
	}
	// Deserialize `DefaultCreatorReward`:
	err = decoder.Decode(&obj.DefaultCreatorReward)
	if err != nil {
		return err
	}
	// Deserialize `DefaultGraduationReward`:
	err = decoder.Decode(&obj.DefaultGraduationReward)
	if err != nil {
		return err
	}
	// Deserialize `DefaultGraduationTarget`:
	err = decoder.Decode(&obj.DefaultGraduationTarget)
	if err != nil {
		return err
	}
	// Deserialize `DefaultMaxBuyAmount`:
	err = decoder.Decode(&obj.DefaultMaxBuyAmount)
	if err != nil {
		return err
	}
	// Deserialize `DefaultMaxSellAmount`:
	err = decoder.Decode(&obj.DefaultMaxSellAmount)
	if err != nil {
		return err
	}
	// Deserialize `DefaultStartPrice`:
	err = decoder.Decode(&obj.DefaultStartPrice)
	if err != nil {
		return err
	}
	// Deserialize `DefaultEndPrice`:
	err = decoder.Decode(&obj.DefaultEndPrice)
	if err != nil {
		return err
	}
	// Deserialize `DefaultControlPoints`:
	err = decoder.Decode(&obj.DefaultControlPoints)
	if err != nil {
		return err
	}
	// Deserialize `DefaultSwapFeeBps`:
	err = decoder.Decode(&obj.DefaultSwapFeeBps)
	if err != nil {
		return err
	}
	// Deserialize `DefaultQuoteFeeBps`:
	err = decoder.Decode(&obj.DefaultQuoteFeeBps)
	if err != nil {
		return err
	}
	// Deserialize `DefaultBaseFeeBps`:
	err = decoder.Decode(&obj.DefaultBaseFeeBps)
	if err != nil {
		return err
	}
	// Deserialize `Padding2`:
	err = decoder.Decode(&obj.Padding2)
	if err != nil {
		return err
	}
	return nil
}

// type BondingCurve struct {
// 	Discriminator          [8]byte          // Account discriminator (fixed)
// 	BaseMint               solana.PublicKey // 32 bytes
// 	QuoteMint              solana.PublicKey
// 	Creator                solana.PublicKey
// 	RetainMintAuthority    bool // 1 byte
// 	BuyRequiresPermission  bool // 1 byte
// 	BuyPermissionBitmap    [32]uint8
// 	SellRequiresPermission bool // 1 byte
// 	SellPermissionBitmap   [32]uint8
// 	QuoteFeeBps            uint16 // 2 bytes
// 	BaseFeeBps             uint16
// 	ControlPoints          [4]uint16
// 	StartPrice             ag_binary.Uint128 // 16 bytes (u128)
// 	EndPrice               ag_binary.Uint128
// 	QuoteAmount            uint64 // 8 bytes
// 	BaseAmount             uint64
// 	LaunchSlot             uint64
// 	CreatorReward          uint64
// 	GraduationTarget       uint64
// 	GraduationSlot         uint64
// 	GraduationReward       uint64
// 	MaxBuyAmount           uint64
// 	MaxSellAmount          uint64
// 	SwapFeeBps             uint16
// 	BaseAllocationBps      uint16
// 	GraduationMethods      [8]GraduationMethodData
// 	MinReserveBps          uint16
// 	Padding1               [2]uint8
// 	PremintedSupply        uint64
// 	Padding2               [728]uint8
// }
