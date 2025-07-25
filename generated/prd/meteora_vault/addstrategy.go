// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// add a strategy
type AddStrategy struct {

	// [0] = [WRITE] vault
	// ··········· vault
	//
	// [1] = [] strategy
	// ··········· strategy
	//
	// [2] = [SIGNER] admin
	// ··········· admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddStrategyInstructionBuilder creates a new `AddStrategy` instruction builder.
func NewAddStrategyInstructionBuilder() *AddStrategy {
	nd := &AddStrategy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetVaultAccount sets the "vault" account.
// vault
func (inst *AddStrategy) SetVaultAccount(vault ag_solanago.PublicKey) *AddStrategy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
// vault
func (inst *AddStrategy) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStrategyAccount sets the "strategy" account.
// strategy
func (inst *AddStrategy) SetStrategyAccount(strategy ag_solanago.PublicKey) *AddStrategy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(strategy)
	return inst
}

// GetStrategyAccount gets the "strategy" account.
// strategy
func (inst *AddStrategy) GetStrategyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAdminAccount sets the "admin" account.
// admin
func (inst *AddStrategy) SetAdminAccount(admin ag_solanago.PublicKey) *AddStrategy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
// admin
func (inst *AddStrategy) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst *AddStrategy) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *AddStrategy) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *AddStrategy {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:3], metas...)
	return inst
}

func (inst *AddStrategy) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3:]
}

func (inst AddStrategy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddStrategy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddStrategy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddStrategy) Validate() error {
	if len(inst.AccountMetaSlice) < 3 {
		return errors.New("accounts slice has wrong length: expected 3 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Strategy is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *AddStrategy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddStrategy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   vault", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("strategy", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   admin", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj AddStrategy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AddStrategy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAddStrategyInstruction declares a new AddStrategy instruction with the provided parameters and accounts.
func NewAddStrategyInstruction(
	// Accounts:
	vault ag_solanago.PublicKey,
	strategy ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *AddStrategy {
	return NewAddStrategyInstructionBuilder().
		SetVaultAccount(vault).
		SetStrategyAccount(strategy).
		SetAdminAccount(admin)
}

// NewSimpleAddStrategyInstruction declares a new AddStrategy instruction with the provided parameters and accounts.
func NewSimpleAddStrategyInstruction(
	// Accounts:
	vault ag_solanago.PublicKey,
	strategy ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *AddStrategy {
	return NewAddStrategyInstructionBuilder().
		SetVaultAccount(vault).
		SetStrategyAccount(strategy).
		SetAdminAccount(admin)
}
