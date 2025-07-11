// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_pools

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Update activation slot
type UpdateActivationPoint struct {
	NewActivationPoint *uint64

	// [0] = [WRITE] pool
	// ··········· Pool account (PDA)
	//
	// [1] = [SIGNER] admin
	// ··········· Admin account.
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateActivationPointInstructionBuilder creates a new `UpdateActivationPoint` instruction builder.
func NewUpdateActivationPointInstructionBuilder() *UpdateActivationPoint {
	nd := &UpdateActivationPoint{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewActivationPoint sets the "new_activation_point" parameter.
func (inst *UpdateActivationPoint) SetNewActivationPoint(new_activation_point uint64) *UpdateActivationPoint {
	inst.NewActivationPoint = &new_activation_point
	return inst
}

// SetPoolAccount sets the "pool" account.
// Pool account (PDA)
func (inst *UpdateActivationPoint) SetPoolAccount(pool ag_solanago.PublicKey) *UpdateActivationPoint {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
// Pool account (PDA)
func (inst *UpdateActivationPoint) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
// Admin account.
func (inst *UpdateActivationPoint) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateActivationPoint {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
// Admin account.
func (inst *UpdateActivationPoint) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst *UpdateActivationPoint) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *UpdateActivationPoint) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *UpdateActivationPoint {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:2], metas...)
	return inst
}

func (inst *UpdateActivationPoint) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2:]
}

func (inst UpdateActivationPoint) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateActivationPoint,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateActivationPoint) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateActivationPoint) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewActivationPoint == nil {
			return errors.New("newActivationPoint parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 2 {
		return errors.New("accounts slice has wrong length: expected 2 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *UpdateActivationPoint) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateActivationPoint")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  NewActivationPoint", *inst.NewActivationPoint))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("admin", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj UpdateActivationPoint) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewActivationPoint` param:
	err = encoder.Encode(obj.NewActivationPoint)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateActivationPoint) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewActivationPoint`:
	err = decoder.Decode(&obj.NewActivationPoint)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateActivationPointInstruction declares a new UpdateActivationPoint instruction with the provided parameters and accounts.
func NewUpdateActivationPointInstruction(
	// Parameters:
	new_activation_point uint64,
	// Accounts:
	pool ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *UpdateActivationPoint {
	return NewUpdateActivationPointInstructionBuilder().
		SetNewActivationPoint(new_activation_point).
		SetPoolAccount(pool).
		SetAdminAccount(admin)
}

// NewSimpleUpdateActivationPointInstruction declares a new UpdateActivationPoint instruction with the provided parameters and accounts.
func NewSimpleUpdateActivationPointInstruction(
	// Parameters:
	new_activation_point uint64,
	// Accounts:
	pool ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *UpdateActivationPoint {
	return NewUpdateActivationPointInstructionBuilder().
		SetNewActivationPoint(new_activation_point).
		SetPoolAccount(pool).
		SetAdminAccount(admin)
}
