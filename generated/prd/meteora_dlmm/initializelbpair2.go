// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeLbPair2 is the `initialize_lb_pair2` instruction.
type InitializeLbPair2 struct {
	Params *InitializeLbPair2Params

	// [0] = [WRITE] lb_pair
	//
	// [1] = [WRITE] bin_array_bitmap_extension
	//
	// [2] = [] token_mint_x
	//
	// [3] = [] token_mint_y
	//
	// [4] = [WRITE] reserve_x
	//
	// [5] = [WRITE] reserve_y
	//
	// [6] = [WRITE] oracle
	//
	// [7] = [] preset_parameter
	//
	// [8] = [WRITE, SIGNER] funder
	//
	// [9] = [] token_badge_x
	//
	// [10] = [] token_badge_y
	//
	// [11] = [] token_program_x
	//
	// [12] = [] token_program_y
	//
	// [13] = [] system_program
	//
	// [14] = [] event_authority
	//
	// [15] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeLbPair2InstructionBuilder creates a new `InitializeLbPair2` instruction builder.
func NewInitializeLbPair2InstructionBuilder() *InitializeLbPair2 {
	nd := &InitializeLbPair2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	nd.AccountMetaSlice[15] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetParams sets the "params" parameter.
func (inst *InitializeLbPair2) SetParams(params InitializeLbPair2Params) *InitializeLbPair2 {
	inst.Params = &params
	return inst
}

// SetLbPairAccount sets the "lb_pair" account.
func (inst *InitializeLbPair2) SetLbPairAccount(lbPair ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lb_pair" account.
func (inst *InitializeLbPair2) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBinArrayBitmapExtensionAccount sets the "bin_array_bitmap_extension" account.
func (inst *InitializeLbPair2) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(binArrayBitmapExtension).WRITE()
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "bin_array_bitmap_extension" account (optional).
func (inst *InitializeLbPair2) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenMintXAccount sets the "token_mint_x" account.
func (inst *InitializeLbPair2) SetTokenMintXAccount(tokenMintX ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenMintX)
	return inst
}

// GetTokenMintXAccount gets the "token_mint_x" account.
func (inst *InitializeLbPair2) GetTokenMintXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMintYAccount sets the "token_mint_y" account.
func (inst *InitializeLbPair2) SetTokenMintYAccount(tokenMintY ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMintY)
	return inst
}

// GetTokenMintYAccount gets the "token_mint_y" account.
func (inst *InitializeLbPair2) GetTokenMintYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReserveXAccount sets the "reserve_x" account.
func (inst *InitializeLbPair2) SetReserveXAccount(reserveX ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(reserveX).WRITE()
	return inst
}

// GetReserveXAccount gets the "reserve_x" account.
func (inst *InitializeLbPair2) GetReserveXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReserveYAccount sets the "reserve_y" account.
func (inst *InitializeLbPair2) SetReserveYAccount(reserveY ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(reserveY).WRITE()
	return inst
}

// GetReserveYAccount gets the "reserve_y" account.
func (inst *InitializeLbPair2) GetReserveYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetOracleAccount sets the "oracle" account.
func (inst *InitializeLbPair2) SetOracleAccount(oracle ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(oracle).WRITE()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *InitializeLbPair2) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPresetParameterAccount sets the "preset_parameter" account.
func (inst *InitializeLbPair2) SetPresetParameterAccount(presetParameter ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(presetParameter)
	return inst
}

// GetPresetParameterAccount gets the "preset_parameter" account.
func (inst *InitializeLbPair2) GetPresetParameterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetFunderAccount sets the "funder" account.
func (inst *InitializeLbPair2) SetFunderAccount(funder ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *InitializeLbPair2) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenBadgeXAccount sets the "token_badge_x" account.
func (inst *InitializeLbPair2) SetTokenBadgeXAccount(tokenBadgeX ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenBadgeX)
	return inst
}

// GetTokenBadgeXAccount gets the "token_badge_x" account (optional).
func (inst *InitializeLbPair2) GetTokenBadgeXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenBadgeYAccount sets the "token_badge_y" account.
func (inst *InitializeLbPair2) SetTokenBadgeYAccount(tokenBadgeY ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenBadgeY)
	return inst
}

// GetTokenBadgeYAccount gets the "token_badge_y" account (optional).
func (inst *InitializeLbPair2) GetTokenBadgeYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramXAccount sets the "token_program_x" account.
func (inst *InitializeLbPair2) SetTokenProgramXAccount(tokenProgramX ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgramX)
	return inst
}

// GetTokenProgramXAccount gets the "token_program_x" account.
func (inst *InitializeLbPair2) GetTokenProgramXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenProgramYAccount sets the "token_program_y" account.
func (inst *InitializeLbPair2) SetTokenProgramYAccount(tokenProgramY ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenProgramY)
	return inst
}

// GetTokenProgramYAccount gets the "token_program_y" account.
func (inst *InitializeLbPair2) GetTokenProgramYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *InitializeLbPair2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *InitializeLbPair2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *InitializeLbPair2) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *InitializeLbPair2) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetProgramAccount sets the "program" account.
func (inst *InitializeLbPair2) SetProgramAccount(program ag_solanago.PublicKey) *InitializeLbPair2 {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *InitializeLbPair2) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst *InitializeLbPair2) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *InitializeLbPair2) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *InitializeLbPair2 {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:16], metas...)
	return inst
}

func (inst *InitializeLbPair2) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16:]
}

func (inst InitializeLbPair2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeLbPair2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeLbPair2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeLbPair2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("params parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 16 {
		return errors.New("accounts slice has wrong length: expected 16 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LbPair is not set")
		}

		// [1] = BinArrayBitmapExtension is optional

		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenMintX is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenMintY is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReserveX is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReserveY is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PresetParameter is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Funder is not set")
		}

		// [9] = TokenBadgeX is optional

		// [10] = TokenBadgeY is optional

		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgramX is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenProgramY is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SystemProgram is not set")
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

func (inst *InitializeLbPair2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeLbPair2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   lb_pair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("bin_array_bitmap_extension", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              token_mint_x", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              token_mint_y", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 reserve_x", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 reserve_y", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                    oracle", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          preset_parameter", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                    funder", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("             token_badge_x", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("             token_badge_y", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("           token_program_x", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("           token_program_y", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("            system_program", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("           event_authority", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                   program", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj InitializeLbPair2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeLbPair2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeLbPair2Instruction declares a new InitializeLbPair2 instruction with the provided parameters and accounts.
func NewInitializeLbPair2Instruction(
	// Parameters:
	params InitializeLbPair2Params,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	tokenMintX ag_solanago.PublicKey,
	tokenMintY ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	presetParameter ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	tokenBadgeX ag_solanago.PublicKey,
	tokenBadgeY ag_solanago.PublicKey,
	tokenProgramX ag_solanago.PublicKey,
	tokenProgramY ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *InitializeLbPair2 {
	return NewInitializeLbPair2InstructionBuilder().
		SetParams(params).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetTokenMintXAccount(tokenMintX).
		SetTokenMintYAccount(tokenMintY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetOracleAccount(oracle).
		SetPresetParameterAccount(presetParameter).
		SetFunderAccount(funder).
		SetTokenBadgeXAccount(tokenBadgeX).
		SetTokenBadgeYAccount(tokenBadgeY).
		SetTokenProgramXAccount(tokenProgramX).
		SetTokenProgramYAccount(tokenProgramY).
		SetSystemProgramAccount(systemProgram).
		SetEventAuthorityAccount(eventAuthority)
}

// NewSimpleInitializeLbPair2Instruction declares a new InitializeLbPair2 instruction with the provided parameters and accounts.
func NewSimpleInitializeLbPair2Instruction(
	// Parameters:
	params InitializeLbPair2Params,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	tokenMintX ag_solanago.PublicKey,
	tokenMintY ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	presetParameter ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	tokenBadgeX ag_solanago.PublicKey,
	tokenBadgeY ag_solanago.PublicKey,
	tokenProgramX ag_solanago.PublicKey,
	tokenProgramY ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *InitializeLbPair2 {
	return NewInitializeLbPair2InstructionBuilder().
		SetParams(params).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetTokenMintXAccount(tokenMintX).
		SetTokenMintYAccount(tokenMintY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetOracleAccount(oracle).
		SetPresetParameterAccount(presetParameter).
		SetFunderAccount(funder).
		SetTokenBadgeXAccount(tokenBadgeX).
		SetTokenBadgeYAccount(tokenBadgeY).
		SetTokenProgramXAccount(tokenProgramX).
		SetTokenProgramYAccount(tokenProgramY).
		SetSystemProgramAccount(systemProgram).
		SetEventAuthorityAccount(eventAuthority)
}
