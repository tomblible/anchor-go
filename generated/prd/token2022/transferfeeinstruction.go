// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package token2022

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// TransferFeeInstruction is the `transferFeeInstruction` instruction.
type TransferFeeInstruction struct {
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTransferFeeInstructionInstructionBuilder creates a new `TransferFeeInstruction` instruction builder.
func NewTransferFeeInstructionInstructionBuilder() *TransferFeeInstruction {
	nd := &TransferFeeInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 0),
	}
	return nd
}

func (inst *TransferFeeInstruction) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *TransferFeeInstruction) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *TransferFeeInstruction {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:0], metas...)
	return inst
}

func (inst *TransferFeeInstruction) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0:]
}

func (inst TransferFeeInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_TransferFeeInstruction),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferFeeInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferFeeInstruction) Validate() error {
	return nil
}

func (inst *TransferFeeInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferFeeInstruction")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=0]").ParentFunc(func(accountsBranch ag_treeout.Branches) {})
				})
		})
}

func (obj TransferFeeInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *TransferFeeInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewTransferFeeInstructionInstruction declares a new TransferFeeInstruction instruction with the provided parameters and accounts.
func NewTransferFeeInstructionInstruction() *TransferFeeInstruction {
	return NewTransferFeeInstructionInstructionBuilder()
}

// NewSimpleTransferFeeInstructionInstruction declares a new TransferFeeInstruction instruction with the provided parameters and accounts.
func NewSimpleTransferFeeInstructionInstruction() *TransferFeeInstruction {
	return NewTransferFeeInstructionInstructionBuilder()
}
