// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// enable vault
type EnableVault struct {
	Enabled *uint8

	// [0] = [WRITE] vault
	// ··········· Vault account
	//
	// [1] = [SIGNER] admin
	// ··········· Admin account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEnableVaultInstructionBuilder creates a new `EnableVault` instruction builder.
func NewEnableVaultInstructionBuilder() *EnableVault {
	nd := &EnableVault{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetEnabled sets the "enabled" parameter.
func (inst *EnableVault) SetEnabled(enabled uint8) *EnableVault {
	inst.Enabled = &enabled
	return inst
}

// SetVaultAccount sets the "vault" account.
// Vault account
func (inst *EnableVault) SetVaultAccount(vault ag_solanago.PublicKey) *EnableVault {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
// Vault account
func (inst *EnableVault) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
// Admin account
func (inst *EnableVault) SetAdminAccount(admin ag_solanago.PublicKey) *EnableVault {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
// Admin account
func (inst *EnableVault) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst *EnableVault) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *EnableVault) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *EnableVault {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:2], metas...)
	return inst
}

func (inst *EnableVault) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2:]
}

func (inst EnableVault) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EnableVault,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EnableVault) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EnableVault) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Enabled == nil {
			return errors.New("enabled parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 2 {
		return errors.New("accounts slice has wrong length: expected 2 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *EnableVault) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EnableVault")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Enabled", *inst.Enabled))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("vault", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("admin", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj EnableVault) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}
func (obj *EnableVault) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}

// NewEnableVaultInstruction declares a new EnableVault instruction with the provided parameters and accounts.
func NewEnableVaultInstruction(
	// Parameters:
	enabled uint8,
	// Accounts:
	vault ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *EnableVault {
	return NewEnableVaultInstructionBuilder().
		SetEnabled(enabled).
		SetVaultAccount(vault).
		SetAdminAccount(admin)
}

// NewSimpleEnableVaultInstruction declares a new EnableVault instruction with the provided parameters and accounts.
func NewSimpleEnableVaultInstruction(
	// Parameters:
	enabled uint8,
	// Accounts:
	vault ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *EnableVault {
	return NewEnableVaultInstructionBuilder().
		SetEnabled(enabled).
		SetVaultAccount(vault).
		SetAdminAccount(admin)
}
