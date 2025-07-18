// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package hpnf

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Transfer reward owner
//
// # Arguments
//
// * `ctx`- The context of accounts
// * `new_owner`- new owner pubkey
//
type TransferRewardOwner struct {
	NewOwner *ag_solanago.PublicKey

	// [0] = [SIGNER] authority
	// ··········· Address to be set as operation account owner.
	//
	// [1] = [WRITE] pool_state
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTransferRewardOwnerInstructionBuilder creates a new `TransferRewardOwner` instruction builder.
func NewTransferRewardOwnerInstructionBuilder() *TransferRewardOwner {
	nd := &TransferRewardOwner{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Owner).SIGNER()
	return nd
}

// SetNewOwner sets the "new_owner" parameter.
func (inst *TransferRewardOwner) SetNewOwner(new_owner ag_solanago.PublicKey) *TransferRewardOwner {
	inst.NewOwner = &new_owner
	return inst
}

// SetAuthorityAccount sets the "authority" account.
// Address to be set as operation account owner.
func (inst *TransferRewardOwner) SetAuthorityAccount(authority ag_solanago.PublicKey) *TransferRewardOwner {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// Address to be set as operation account owner.
func (inst *TransferRewardOwner) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPoolStateAccount sets the "pool_state" account.
func (inst *TransferRewardOwner) SetPoolStateAccount(poolState ag_solanago.PublicKey) *TransferRewardOwner {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "pool_state" account.
func (inst *TransferRewardOwner) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst *TransferRewardOwner) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *TransferRewardOwner) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *TransferRewardOwner {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:2], metas...)
	return inst
}

func (inst *TransferRewardOwner) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2:]
}

func (inst TransferRewardOwner) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TransferRewardOwner,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferRewardOwner) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferRewardOwner) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewOwner == nil {
			return errors.New("newOwner parameter is not set")
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

func (inst *TransferRewardOwner) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferRewardOwner")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" NewOwner", *inst.NewOwner))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("pool_state", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj TransferRewardOwner) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewOwner` param:
	err = encoder.Encode(obj.NewOwner)
	if err != nil {
		return err
	}
	return nil
}
func (obj *TransferRewardOwner) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewOwner`:
	err = decoder.Decode(&obj.NewOwner)
	if err != nil {
		return err
	}
	return nil
}

// NewTransferRewardOwnerInstruction declares a new TransferRewardOwner instruction with the provided parameters and accounts.
func NewTransferRewardOwnerInstruction(
	// Parameters:
	new_owner ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey) *TransferRewardOwner {
	return NewTransferRewardOwnerInstructionBuilder().
		SetNewOwner(new_owner).
		SetPoolStateAccount(poolState)
}

// NewSimpleTransferRewardOwnerInstruction declares a new TransferRewardOwner instruction with the provided parameters and accounts.
func NewSimpleTransferRewardOwnerInstruction(
	// Parameters:
	new_owner ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey) *TransferRewardOwner {
	return NewTransferRewardOwnerInstructionBuilder().
		SetNewOwner(new_owner).
		SetPoolStateAccount(poolState)
}
