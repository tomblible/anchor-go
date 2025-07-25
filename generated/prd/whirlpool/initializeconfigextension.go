// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package orca_whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeConfigExtension is the `initialize_config_extension` instruction.
type InitializeConfigExtension struct {

	// [0] = [] config
	//
	// [1] = [WRITE] config_extension
	//
	// [2] = [WRITE, SIGNER] funder
	//
	// [3] = [SIGNER] fee_authority
	//
	// [4] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeConfigExtensionInstructionBuilder creates a new `InitializeConfigExtension` instruction builder.
func NewInitializeConfigExtensionInstructionBuilder() *InitializeConfigExtension {
	nd := &InitializeConfigExtension{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetConfigAccount sets the "config" account.
func (inst *InitializeConfigExtension) SetConfigAccount(config ag_solanago.PublicKey) *InitializeConfigExtension {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *InitializeConfigExtension) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigExtensionAccount sets the "config_extension" account.
func (inst *InitializeConfigExtension) SetConfigExtensionAccount(configExtension ag_solanago.PublicKey) *InitializeConfigExtension {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(configExtension).WRITE()
	return inst
}

// GetConfigExtensionAccount gets the "config_extension" account.
func (inst *InitializeConfigExtension) GetConfigExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFunderAccount sets the "funder" account.
func (inst *InitializeConfigExtension) SetFunderAccount(funder ag_solanago.PublicKey) *InitializeConfigExtension {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *InitializeConfigExtension) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFeeAuthorityAccount sets the "fee_authority" account.
func (inst *InitializeConfigExtension) SetFeeAuthorityAccount(feeAuthority ag_solanago.PublicKey) *InitializeConfigExtension {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(feeAuthority).SIGNER()
	return inst
}

// GetFeeAuthorityAccount gets the "fee_authority" account.
func (inst *InitializeConfigExtension) GetFeeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *InitializeConfigExtension) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeConfigExtension {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *InitializeConfigExtension) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst *InitializeConfigExtension) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *InitializeConfigExtension) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *InitializeConfigExtension {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:5], metas...)
	return inst
}

func (inst *InitializeConfigExtension) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5:]
}

func (inst InitializeConfigExtension) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeConfigExtension,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeConfigExtension) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeConfigExtension) Validate() error {
	if len(inst.AccountMetaSlice) < 5 {
		return errors.New("accounts slice has wrong length: expected 5 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ConfigExtension is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FeeAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitializeConfigExtension) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeConfigExtension")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("config_extension", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          funder", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   fee_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  system_program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeConfigExtension) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitializeConfigExtension) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeConfigExtensionInstruction declares a new InitializeConfigExtension instruction with the provided parameters and accounts.
func NewInitializeConfigExtensionInstruction(
	// Accounts:
	config ag_solanago.PublicKey,
	configExtension ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	feeAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeConfigExtension {
	return NewInitializeConfigExtensionInstructionBuilder().
		SetConfigAccount(config).
		SetConfigExtensionAccount(configExtension).
		SetFunderAccount(funder).
		SetFeeAuthorityAccount(feeAuthority).
		SetSystemProgramAccount(systemProgram)
}

// NewSimpleInitializeConfigExtensionInstruction declares a new InitializeConfigExtension instruction with the provided parameters and accounts.
func NewSimpleInitializeConfigExtensionInstruction(
	// Accounts:
	config ag_solanago.PublicKey,
	configExtension ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	feeAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeConfigExtension {
	return NewInitializeConfigExtensionInstructionBuilder().
		SetConfigAccount(config).
		SetConfigExtensionAccount(configExtension).
		SetFunderAccount(funder).
		SetFeeAuthorityAccount(feeAuthority).
		SetSystemProgramAccount(systemProgram)
}
