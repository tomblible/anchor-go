// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeLbPair is the `initialize_lb_pair` instruction.
type InitializeLbPair struct {
	ActiveId *int32
	BinStep  *uint16

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
	// [9] = [] token_program
	//
	// [10] = [] system_program
	//
	// [11] = [] rent
	//
	// [12] = [] event_authority
	//
	// [13] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeLbPairInstructionBuilder creates a new `InitializeLbPair` instruction builder.
func NewInitializeLbPairInstructionBuilder() *InitializeLbPair {
	nd := &InitializeLbPair{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	nd.AccountMetaSlice[13] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetActiveId sets the "active_id" parameter.
func (inst *InitializeLbPair) SetActiveId(active_id int32) *InitializeLbPair {
	inst.ActiveId = &active_id
	return inst
}

// SetBinStep sets the "bin_step" parameter.
func (inst *InitializeLbPair) SetBinStep(bin_step uint16) *InitializeLbPair {
	inst.BinStep = &bin_step
	return inst
}

// SetLbPairAccount sets the "lb_pair" account.
func (inst *InitializeLbPair) SetLbPairAccount(lbPair ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lb_pair" account.
func (inst *InitializeLbPair) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBinArrayBitmapExtensionAccount sets the "bin_array_bitmap_extension" account.
func (inst *InitializeLbPair) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(binArrayBitmapExtension).WRITE()
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "bin_array_bitmap_extension" account (optional).
func (inst *InitializeLbPair) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenMintXAccount sets the "token_mint_x" account.
func (inst *InitializeLbPair) SetTokenMintXAccount(tokenMintX ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenMintX)
	return inst
}

// GetTokenMintXAccount gets the "token_mint_x" account.
func (inst *InitializeLbPair) GetTokenMintXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMintYAccount sets the "token_mint_y" account.
func (inst *InitializeLbPair) SetTokenMintYAccount(tokenMintY ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMintY)
	return inst
}

// GetTokenMintYAccount gets the "token_mint_y" account.
func (inst *InitializeLbPair) GetTokenMintYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReserveXAccount sets the "reserve_x" account.
func (inst *InitializeLbPair) SetReserveXAccount(reserveX ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(reserveX).WRITE()
	return inst
}

// GetReserveXAccount gets the "reserve_x" account.
func (inst *InitializeLbPair) GetReserveXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReserveYAccount sets the "reserve_y" account.
func (inst *InitializeLbPair) SetReserveYAccount(reserveY ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(reserveY).WRITE()
	return inst
}

// GetReserveYAccount gets the "reserve_y" account.
func (inst *InitializeLbPair) GetReserveYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetOracleAccount sets the "oracle" account.
func (inst *InitializeLbPair) SetOracleAccount(oracle ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(oracle).WRITE()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *InitializeLbPair) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPresetParameterAccount sets the "preset_parameter" account.
func (inst *InitializeLbPair) SetPresetParameterAccount(presetParameter ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(presetParameter)
	return inst
}

// GetPresetParameterAccount gets the "preset_parameter" account.
func (inst *InitializeLbPair) GetPresetParameterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetFunderAccount sets the "funder" account.
func (inst *InitializeLbPair) SetFunderAccount(funder ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *InitializeLbPair) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *InitializeLbPair) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *InitializeLbPair) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *InitializeLbPair) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *InitializeLbPair) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeLbPair) SetRentAccount(rent ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeLbPair) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *InitializeLbPair) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *InitializeLbPair) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetProgramAccount sets the "program" account.
func (inst *InitializeLbPair) SetProgramAccount(program ag_solanago.PublicKey) *InitializeLbPair {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *InitializeLbPair) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst *InitializeLbPair) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *InitializeLbPair) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *InitializeLbPair {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:14], metas...)
	return inst
}

func (inst *InitializeLbPair) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[14:]
}

func (inst InitializeLbPair) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeLbPair,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeLbPair) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeLbPair) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ActiveId == nil {
			return errors.New("activeId parameter is not set")
		}
		if inst.BinStep == nil {
			return errors.New("binStep parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 14 {
		return errors.New("accounts slice has wrong length: expected 14 accounts")
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
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *InitializeLbPair) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeLbPair")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" ActiveId", *inst.ActiveId))
						paramsBranch.Child(ag_format.Param("  BinStep", *inst.BinStep))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   lb_pair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("bin_array_bitmap_extension", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              token_mint_x", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              token_mint_y", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 reserve_x", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 reserve_y", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                    oracle", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          preset_parameter", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                    funder", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("             token_program", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            system_program", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                      rent", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("           event_authority", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                   program", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj InitializeLbPair) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ActiveId` param:
	err = encoder.Encode(obj.ActiveId)
	if err != nil {
		return err
	}
	// Serialize `BinStep` param:
	err = encoder.Encode(obj.BinStep)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeLbPair) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ActiveId`:
	err = decoder.Decode(&obj.ActiveId)
	if err != nil {
		return err
	}
	// Deserialize `BinStep`:
	err = decoder.Decode(&obj.BinStep)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeLbPairInstruction declares a new InitializeLbPair instruction with the provided parameters and accounts.
func NewInitializeLbPairInstruction(
	// Parameters:
	active_id int32,
	bin_step uint16,
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
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *InitializeLbPair {
	return NewInitializeLbPairInstructionBuilder().
		SetActiveId(active_id).
		SetBinStep(bin_step).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetTokenMintXAccount(tokenMintX).
		SetTokenMintYAccount(tokenMintY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetOracleAccount(oracle).
		SetPresetParameterAccount(presetParameter).
		SetFunderAccount(funder).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetEventAuthorityAccount(eventAuthority)
}

// NewSimpleInitializeLbPairInstruction declares a new InitializeLbPair instruction with the provided parameters and accounts.
func NewSimpleInitializeLbPairInstruction(
	// Parameters:
	active_id int32,
	bin_step uint16,
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
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *InitializeLbPair {
	return NewInitializeLbPairInstructionBuilder().
		SetActiveId(active_id).
		SetBinStep(bin_step).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetTokenMintXAccount(tokenMintX).
		SetTokenMintYAccount(tokenMintY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetOracleAccount(oracle).
		SetPresetParameterAccount(presetParameter).
		SetFunderAccount(funder).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetEventAuthorityAccount(eventAuthority)
}
