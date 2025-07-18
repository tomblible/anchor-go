// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AddLiquidityByStrategy is the `add_liquidity_by_strategy` instruction.
type AddLiquidityByStrategy struct {
	LiquidityParameter *LiquidityParameterByStrategy

	// [0] = [WRITE] position
	//
	// [1] = [WRITE] lb_pair
	//
	// [2] = [WRITE] bin_array_bitmap_extension
	//
	// [3] = [WRITE] user_token_x
	//
	// [4] = [WRITE] user_token_y
	//
	// [5] = [WRITE] reserve_x
	//
	// [6] = [WRITE] reserve_y
	//
	// [7] = [] token_x_mint
	//
	// [8] = [] token_y_mint
	//
	// [9] = [WRITE] bin_array_lower
	//
	// [10] = [WRITE] bin_array_upper
	//
	// [11] = [SIGNER] sender
	//
	// [12] = [] token_x_program
	//
	// [13] = [] token_y_program
	//
	// [14] = [] event_authority
	//
	// [15] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddLiquidityByStrategyInstructionBuilder creates a new `AddLiquidityByStrategy` instruction builder.
func NewAddLiquidityByStrategyInstructionBuilder() *AddLiquidityByStrategy {
	nd := &AddLiquidityByStrategy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	nd.AccountMetaSlice[15] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetLiquidityParameter sets the "liquidity_parameter" parameter.
func (inst *AddLiquidityByStrategy) SetLiquidityParameter(liquidity_parameter LiquidityParameterByStrategy) *AddLiquidityByStrategy {
	inst.LiquidityParameter = &liquidity_parameter
	return inst
}

// SetPositionAccount sets the "position" account.
func (inst *AddLiquidityByStrategy) SetPositionAccount(position ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *AddLiquidityByStrategy) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLbPairAccount sets the "lb_pair" account.
func (inst *AddLiquidityByStrategy) SetLbPairAccount(lbPair ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lb_pair" account.
func (inst *AddLiquidityByStrategy) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBinArrayBitmapExtensionAccount sets the "bin_array_bitmap_extension" account.
func (inst *AddLiquidityByStrategy) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(binArrayBitmapExtension).WRITE()
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "bin_array_bitmap_extension" account (optional).
func (inst *AddLiquidityByStrategy) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserTokenXAccount sets the "user_token_x" account.
func (inst *AddLiquidityByStrategy) SetUserTokenXAccount(userTokenX ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userTokenX).WRITE()
	return inst
}

// GetUserTokenXAccount gets the "user_token_x" account.
func (inst *AddLiquidityByStrategy) GetUserTokenXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserTokenYAccount sets the "user_token_y" account.
func (inst *AddLiquidityByStrategy) SetUserTokenYAccount(userTokenY ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userTokenY).WRITE()
	return inst
}

// GetUserTokenYAccount gets the "user_token_y" account.
func (inst *AddLiquidityByStrategy) GetUserTokenYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReserveXAccount sets the "reserve_x" account.
func (inst *AddLiquidityByStrategy) SetReserveXAccount(reserveX ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(reserveX).WRITE()
	return inst
}

// GetReserveXAccount gets the "reserve_x" account.
func (inst *AddLiquidityByStrategy) GetReserveXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetReserveYAccount sets the "reserve_y" account.
func (inst *AddLiquidityByStrategy) SetReserveYAccount(reserveY ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(reserveY).WRITE()
	return inst
}

// GetReserveYAccount gets the "reserve_y" account.
func (inst *AddLiquidityByStrategy) GetReserveYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenXMintAccount sets the "token_x_mint" account.
func (inst *AddLiquidityByStrategy) SetTokenXMintAccount(tokenXMint ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenXMint)
	return inst
}

// GetTokenXMintAccount gets the "token_x_mint" account.
func (inst *AddLiquidityByStrategy) GetTokenXMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenYMintAccount sets the "token_y_mint" account.
func (inst *AddLiquidityByStrategy) SetTokenYMintAccount(tokenYMint ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenYMint)
	return inst
}

// GetTokenYMintAccount gets the "token_y_mint" account.
func (inst *AddLiquidityByStrategy) GetTokenYMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetBinArrayLowerAccount sets the "bin_array_lower" account.
func (inst *AddLiquidityByStrategy) SetBinArrayLowerAccount(binArrayLower ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(binArrayLower).WRITE()
	return inst
}

// GetBinArrayLowerAccount gets the "bin_array_lower" account.
func (inst *AddLiquidityByStrategy) GetBinArrayLowerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetBinArrayUpperAccount sets the "bin_array_upper" account.
func (inst *AddLiquidityByStrategy) SetBinArrayUpperAccount(binArrayUpper ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(binArrayUpper).WRITE()
	return inst
}

// GetBinArrayUpperAccount gets the "bin_array_upper" account.
func (inst *AddLiquidityByStrategy) GetBinArrayUpperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSenderAccount sets the "sender" account.
func (inst *AddLiquidityByStrategy) SetSenderAccount(sender ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(sender).SIGNER()
	return inst
}

// GetSenderAccount gets the "sender" account.
func (inst *AddLiquidityByStrategy) GetSenderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenXProgramAccount sets the "token_x_program" account.
func (inst *AddLiquidityByStrategy) SetTokenXProgramAccount(tokenXProgram ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenXProgram)
	return inst
}

// GetTokenXProgramAccount gets the "token_x_program" account.
func (inst *AddLiquidityByStrategy) GetTokenXProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenYProgramAccount sets the "token_y_program" account.
func (inst *AddLiquidityByStrategy) SetTokenYProgramAccount(tokenYProgram ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenYProgram)
	return inst
}

// GetTokenYProgramAccount gets the "token_y_program" account.
func (inst *AddLiquidityByStrategy) GetTokenYProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *AddLiquidityByStrategy) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *AddLiquidityByStrategy) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetProgramAccount sets the "program" account.
func (inst *AddLiquidityByStrategy) SetProgramAccount(program ag_solanago.PublicKey) *AddLiquidityByStrategy {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *AddLiquidityByStrategy) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst *AddLiquidityByStrategy) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *AddLiquidityByStrategy) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *AddLiquidityByStrategy {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:16], metas...)
	return inst
}

func (inst *AddLiquidityByStrategy) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16:]
}

func (inst AddLiquidityByStrategy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddLiquidityByStrategy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddLiquidityByStrategy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddLiquidityByStrategy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.LiquidityParameter == nil {
			return errors.New("liquidityParameter parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 16 {
		return errors.New("accounts slice has wrong length: expected 16 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LbPair is not set")
		}

		// [2] = BinArrayBitmapExtension is optional

		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserTokenX is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserTokenY is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReserveX is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ReserveY is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenXMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenYMint is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.BinArrayLower is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.BinArrayUpper is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Sender is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenXProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenYProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *AddLiquidityByStrategy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddLiquidityByStrategy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" LiquidityParameter", *inst.LiquidityParameter))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  position", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                   lb_pair", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("bin_array_bitmap_extension", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              user_token_x", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("              user_token_y", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 reserve_x", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                 reserve_y", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("              token_x_mint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              token_y_mint", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           bin_array_lower", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("           bin_array_upper", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                    sender", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("           token_x_program", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("           token_y_program", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("           event_authority", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                   program", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj AddLiquidityByStrategy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LiquidityParameter` param:
	err = encoder.Encode(obj.LiquidityParameter)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddLiquidityByStrategy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LiquidityParameter`:
	err = decoder.Decode(&obj.LiquidityParameter)
	if err != nil {
		return err
	}
	return nil
}

// NewAddLiquidityByStrategyInstruction declares a new AddLiquidityByStrategy instruction with the provided parameters and accounts.
func NewAddLiquidityByStrategyInstruction(
	// Parameters:
	liquidity_parameter LiquidityParameterByStrategy,
	// Accounts:
	position ag_solanago.PublicKey,
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	userTokenX ag_solanago.PublicKey,
	userTokenY ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	tokenXMint ag_solanago.PublicKey,
	tokenYMint ag_solanago.PublicKey,
	binArrayLower ag_solanago.PublicKey,
	binArrayUpper ag_solanago.PublicKey,
	sender ag_solanago.PublicKey,
	tokenXProgram ag_solanago.PublicKey,
	tokenYProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *AddLiquidityByStrategy {
	return NewAddLiquidityByStrategyInstructionBuilder().
		SetLiquidityParameter(liquidity_parameter).
		SetPositionAccount(position).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetUserTokenXAccount(userTokenX).
		SetUserTokenYAccount(userTokenY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetTokenXMintAccount(tokenXMint).
		SetTokenYMintAccount(tokenYMint).
		SetBinArrayLowerAccount(binArrayLower).
		SetBinArrayUpperAccount(binArrayUpper).
		SetSenderAccount(sender).
		SetTokenXProgramAccount(tokenXProgram).
		SetTokenYProgramAccount(tokenYProgram).
		SetEventAuthorityAccount(eventAuthority)
}

// NewSimpleAddLiquidityByStrategyInstruction declares a new AddLiquidityByStrategy instruction with the provided parameters and accounts.
func NewSimpleAddLiquidityByStrategyInstruction(
	// Parameters:
	liquidity_parameter LiquidityParameterByStrategy,
	// Accounts:
	position ag_solanago.PublicKey,
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	userTokenX ag_solanago.PublicKey,
	userTokenY ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	tokenXMint ag_solanago.PublicKey,
	tokenYMint ag_solanago.PublicKey,
	binArrayLower ag_solanago.PublicKey,
	binArrayUpper ag_solanago.PublicKey,
	sender ag_solanago.PublicKey,
	tokenXProgram ag_solanago.PublicKey,
	tokenYProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *AddLiquidityByStrategy {
	return NewAddLiquidityByStrategyInstructionBuilder().
		SetLiquidityParameter(liquidity_parameter).
		SetPositionAccount(position).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetUserTokenXAccount(userTokenX).
		SetUserTokenYAccount(userTokenY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetTokenXMintAccount(tokenXMint).
		SetTokenYMintAccount(tokenYMint).
		SetBinArrayLowerAccount(binArrayLower).
		SetBinArrayUpperAccount(binArrayUpper).
		SetSenderAccount(sender).
		SetTokenXProgramAccount(tokenXProgram).
		SetTokenYProgramAccount(tokenYProgram).
		SetEventAuthorityAccount(eventAuthority)
}
