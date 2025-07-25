// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ClosePosition2 is the `close_position2` instruction.
type ClosePosition2 struct {

	// [0] = [WRITE] position
	//
	// [1] = [SIGNER] sender
	//
	// [2] = [WRITE] rent_receiver
	//
	// [3] = [] event_authority
	//
	// [4] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClosePosition2InstructionBuilder creates a new `ClosePosition2` instruction builder.
func NewClosePosition2InstructionBuilder() *ClosePosition2 {
	nd := &ClosePosition2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetPositionAccount sets the "position" account.
func (inst *ClosePosition2) SetPositionAccount(position ag_solanago.PublicKey) *ClosePosition2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *ClosePosition2) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSenderAccount sets the "sender" account.
func (inst *ClosePosition2) SetSenderAccount(sender ag_solanago.PublicKey) *ClosePosition2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(sender).SIGNER()
	return inst
}

// GetSenderAccount gets the "sender" account.
func (inst *ClosePosition2) GetSenderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRentReceiverAccount sets the "rent_receiver" account.
func (inst *ClosePosition2) SetRentReceiverAccount(rentReceiver ag_solanago.PublicKey) *ClosePosition2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(rentReceiver).WRITE()
	return inst
}

// GetRentReceiverAccount gets the "rent_receiver" account.
func (inst *ClosePosition2) GetRentReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *ClosePosition2) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *ClosePosition2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *ClosePosition2) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProgramAccount sets the "program" account.
func (inst *ClosePosition2) SetProgramAccount(program ag_solanago.PublicKey) *ClosePosition2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *ClosePosition2) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst *ClosePosition2) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *ClosePosition2) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *ClosePosition2 {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:5], metas...)
	return inst
}

func (inst *ClosePosition2) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5:]
}

func (inst ClosePosition2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClosePosition2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClosePosition2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClosePosition2) Validate() error {
	if len(inst.AccountMetaSlice) < 5 {
		return errors.New("accounts slice has wrong length: expected 5 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Sender is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.RentReceiver is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *ClosePosition2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClosePosition2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       position", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         sender", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  rent_receiver", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("event_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj ClosePosition2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClosePosition2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClosePosition2Instruction declares a new ClosePosition2 instruction with the provided parameters and accounts.
func NewClosePosition2Instruction(
	// Accounts:
	position ag_solanago.PublicKey,
	sender ag_solanago.PublicKey,
	rentReceiver ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *ClosePosition2 {
	return NewClosePosition2InstructionBuilder().
		SetPositionAccount(position).
		SetSenderAccount(sender).
		SetRentReceiverAccount(rentReceiver).
		SetEventAuthorityAccount(eventAuthority)
}

// NewSimpleClosePosition2Instruction declares a new ClosePosition2 instruction with the provided parameters and accounts.
func NewSimpleClosePosition2Instruction(
	// Accounts:
	position ag_solanago.PublicKey,
	sender ag_solanago.PublicKey,
	rentReceiver ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *ClosePosition2 {
	return NewClosePosition2InstructionBuilder().
		SetPositionAccount(position).
		SetSenderAccount(sender).
		SetRentReceiverAccount(rentReceiver).
		SetEventAuthorityAccount(eventAuthority)
}
