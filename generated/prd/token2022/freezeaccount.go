// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package token2022

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Freeze an Initialized account using the Mint's freeze_authority (if set).
type FreezeAccount struct {

	// [0] = [WRITE] account
	// ··········· The account to freeze.
	//
	// [1] = [] mint
	// ··········· The token mint.
	//
	// [2] = [] authority
	// ··········· The mint freeze authority.
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFreezeAccountInstructionBuilder creates a new `FreezeAccount` instruction builder.
func NewFreezeAccountInstructionBuilder() *FreezeAccount {
	nd := &FreezeAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAccountAccount sets the "account" account.
// The account to freeze.
func (inst *FreezeAccount) SetAccountAccount(account ag_solanago.PublicKey) *FreezeAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(account).WRITE()
	return inst
}

// GetAccountAccount gets the "account" account.
// The account to freeze.
func (inst *FreezeAccount) GetAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
// The token mint.
func (inst *FreezeAccount) SetMintAccount(mint ag_solanago.PublicKey) *FreezeAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// The token mint.
func (inst *FreezeAccount) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
// The mint freeze authority.
func (inst *FreezeAccount) SetAuthorityAccount(authority ag_solanago.PublicKey) *FreezeAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// The mint freeze authority.
func (inst *FreezeAccount) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst *FreezeAccount) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *FreezeAccount) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *FreezeAccount {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:3], metas...)
	return inst
}

func (inst *FreezeAccount) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3:]
}

func (inst FreezeAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_FreezeAccount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FreezeAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FreezeAccount) Validate() error {
	if len(inst.AccountMetaSlice) < 3 {
		return errors.New("accounts slice has wrong length: expected 3 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Account is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *FreezeAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FreezeAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  account", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj FreezeAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *FreezeAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewFreezeAccountInstruction declares a new FreezeAccount instruction with the provided parameters and accounts.
func NewFreezeAccountInstruction(
	// Accounts:
	account ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *FreezeAccount {
	return NewFreezeAccountInstructionBuilder().
		SetAccountAccount(account).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}

// NewSimpleFreezeAccountInstruction declares a new FreezeAccount instruction with the provided parameters and accounts.
func NewSimpleFreezeAccountInstruction(
	// Accounts:
	account ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *FreezeAccount {
	return NewFreezeAccountInstructionBuilder().
		SetAccountAccount(account).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}
