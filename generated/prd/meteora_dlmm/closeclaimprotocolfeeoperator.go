// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CloseClaimProtocolFeeOperator is the `close_claim_protocol_fee_operator` instruction.
type CloseClaimProtocolFeeOperator struct {

	// [0] = [WRITE] claim_fee_operator
	//
	// [1] = [WRITE] rent_receiver
	//
	// [2] = [SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCloseClaimProtocolFeeOperatorInstructionBuilder creates a new `CloseClaimProtocolFeeOperator` instruction builder.
func NewCloseClaimProtocolFeeOperatorInstructionBuilder() *CloseClaimProtocolFeeOperator {
	nd := &CloseClaimProtocolFeeOperator{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetClaimFeeOperatorAccount sets the "claim_fee_operator" account.
func (inst *CloseClaimProtocolFeeOperator) SetClaimFeeOperatorAccount(claimFeeOperator ag_solanago.PublicKey) *CloseClaimProtocolFeeOperator {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(claimFeeOperator).WRITE()
	return inst
}

// GetClaimFeeOperatorAccount gets the "claim_fee_operator" account.
func (inst *CloseClaimProtocolFeeOperator) GetClaimFeeOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRentReceiverAccount sets the "rent_receiver" account.
func (inst *CloseClaimProtocolFeeOperator) SetRentReceiverAccount(rentReceiver ag_solanago.PublicKey) *CloseClaimProtocolFeeOperator {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(rentReceiver).WRITE()
	return inst
}

// GetRentReceiverAccount gets the "rent_receiver" account.
func (inst *CloseClaimProtocolFeeOperator) GetRentReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAdminAccount sets the "admin" account.
func (inst *CloseClaimProtocolFeeOperator) SetAdminAccount(admin ag_solanago.PublicKey) *CloseClaimProtocolFeeOperator {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *CloseClaimProtocolFeeOperator) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst *CloseClaimProtocolFeeOperator) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *CloseClaimProtocolFeeOperator) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *CloseClaimProtocolFeeOperator {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:3], metas...)
	return inst
}

func (inst *CloseClaimProtocolFeeOperator) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3:]
}

func (inst CloseClaimProtocolFeeOperator) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CloseClaimProtocolFeeOperator,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CloseClaimProtocolFeeOperator) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CloseClaimProtocolFeeOperator) Validate() error {
	if len(inst.AccountMetaSlice) < 3 {
		return errors.New("accounts slice has wrong length: expected 3 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.ClaimFeeOperator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.RentReceiver is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *CloseClaimProtocolFeeOperator) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CloseClaimProtocolFeeOperator")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("claim_fee_operator", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     rent_receiver", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             admin", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CloseClaimProtocolFeeOperator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CloseClaimProtocolFeeOperator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCloseClaimProtocolFeeOperatorInstruction declares a new CloseClaimProtocolFeeOperator instruction with the provided parameters and accounts.
func NewCloseClaimProtocolFeeOperatorInstruction(
	// Accounts:
	claimFeeOperator ag_solanago.PublicKey,
	rentReceiver ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *CloseClaimProtocolFeeOperator {
	return NewCloseClaimProtocolFeeOperatorInstructionBuilder().
		SetClaimFeeOperatorAccount(claimFeeOperator).
		SetRentReceiverAccount(rentReceiver).
		SetAdminAccount(admin)
}

// NewSimpleCloseClaimProtocolFeeOperatorInstruction declares a new CloseClaimProtocolFeeOperator instruction with the provided parameters and accounts.
func NewSimpleCloseClaimProtocolFeeOperatorInstruction(
	// Accounts:
	claimFeeOperator ag_solanago.PublicKey,
	rentReceiver ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *CloseClaimProtocolFeeOperator {
	return NewCloseClaimProtocolFeeOperatorInstructionBuilder().
		SetClaimFeeOperatorAccount(claimFeeOperator).
		SetRentReceiverAccount(rentReceiver).
		SetAdminAccount(admin)
}
