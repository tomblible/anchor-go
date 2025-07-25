// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package hpnf

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Update pool status for given value
//
// # Arguments
//
// * `ctx`- The context of accounts
// * `status` - The value of status
//
type UpdatePoolStatus struct {
	Status *uint8

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] pool_state
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePoolStatusInstructionBuilder creates a new `UpdatePoolStatus` instruction builder.
func NewUpdatePoolStatusInstructionBuilder() *UpdatePoolStatus {
	nd := &UpdatePoolStatus{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Owner).SIGNER()
	return nd
}

// SetStatus sets the "status" parameter.
func (inst *UpdatePoolStatus) SetStatus(status uint8) *UpdatePoolStatus {
	inst.Status = &status
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdatePoolStatus) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdatePoolStatus {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdatePoolStatus) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPoolStateAccount sets the "pool_state" account.
func (inst *UpdatePoolStatus) SetPoolStateAccount(poolState ag_solanago.PublicKey) *UpdatePoolStatus {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "pool_state" account.
func (inst *UpdatePoolStatus) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst *UpdatePoolStatus) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *UpdatePoolStatus) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *UpdatePoolStatus {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:2], metas...)
	return inst
}

func (inst *UpdatePoolStatus) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2:]
}

func (inst UpdatePoolStatus) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePoolStatus,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePoolStatus) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePoolStatus) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Status == nil {
			return errors.New("status parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 2 {
		return errors.New("accounts slice has wrong length: expected 2 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PoolState is not set")
		}
	}
	return nil
}

func (inst *UpdatePoolStatus) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePoolStatus")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Status", *inst.Status))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("pool_state", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj UpdatePoolStatus) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdatePoolStatus) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdatePoolStatusInstruction declares a new UpdatePoolStatus instruction with the provided parameters and accounts.
func NewUpdatePoolStatusInstruction(
	// Parameters:
	status uint8,
	poolState ag_solanago.PublicKey) *UpdatePoolStatus {
	return NewUpdatePoolStatusInstructionBuilder().
		SetStatus(status).
		SetPoolStateAccount(poolState)
}

// NewSimpleUpdatePoolStatusInstruction declares a new UpdatePoolStatus instruction with the provided parameters and accounts.
func NewSimpleUpdatePoolStatusInstruction(
	// Parameters:
	status uint8,
	poolState ag_solanago.PublicKey) *UpdatePoolStatus {
	return NewUpdatePoolStatusInstructionBuilder().
		SetStatus(status).
		SetPoolStateAccount(poolState)
}
