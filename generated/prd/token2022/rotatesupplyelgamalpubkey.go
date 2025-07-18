// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package token2022

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RotateSupplyElgamalPubkey is the `rotate_supply_elgamal_pubkey` instruction.
type RotateSupplyElgamalPubkey struct {
	NewSupplyElgamalPubkey *ag_solanago.PublicKey `bin:"optional"`
	ProofInstructionOffset *int8

	// [0] = [WRITE] mint
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRotateSupplyElgamalPubkeyInstructionBuilder creates a new `RotateSupplyElgamalPubkey` instruction builder.
func NewRotateSupplyElgamalPubkeyInstructionBuilder() *RotateSupplyElgamalPubkey {
	nd := &RotateSupplyElgamalPubkey{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewSupplyElgamalPubkey sets the "newSupplyElgamalPubkey" parameter.
func (inst *RotateSupplyElgamalPubkey) SetNewSupplyElgamalPubkey(newSupplyElgamalPubkey ag_solanago.PublicKey) *RotateSupplyElgamalPubkey {
	inst.NewSupplyElgamalPubkey = &newSupplyElgamalPubkey
	return inst
}

// SetProofInstructionOffset sets the "proofInstructionOffset" parameter.
func (inst *RotateSupplyElgamalPubkey) SetProofInstructionOffset(proofInstructionOffset int8) *RotateSupplyElgamalPubkey {
	inst.ProofInstructionOffset = &proofInstructionOffset
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *RotateSupplyElgamalPubkey) SetMintAccount(mint ag_solanago.PublicKey) *RotateSupplyElgamalPubkey {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *RotateSupplyElgamalPubkey) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *RotateSupplyElgamalPubkey) SetAuthorityAccount(authority ag_solanago.PublicKey) *RotateSupplyElgamalPubkey {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *RotateSupplyElgamalPubkey) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst *RotateSupplyElgamalPubkey) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *RotateSupplyElgamalPubkey) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *RotateSupplyElgamalPubkey {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:2], metas...)
	return inst
}

func (inst *RotateSupplyElgamalPubkey) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2:]
}

func (inst RotateSupplyElgamalPubkey) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_RotateSupplyElgamalPubkey),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RotateSupplyElgamalPubkey) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RotateSupplyElgamalPubkey) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ProofInstructionOffset == nil {
			return errors.New("proofInstructionOffset parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 2 {
		return errors.New("accounts slice has wrong length: expected 2 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *RotateSupplyElgamalPubkey) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RotateSupplyElgamalPubkey")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewSupplyElgamalPubkey (OPT)", inst.NewSupplyElgamalPubkey))
						paramsBranch.Child(ag_format.Param("ProofInstructionOffset", *inst.ProofInstructionOffset))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     mint", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj RotateSupplyElgamalPubkey) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewSupplyElgamalPubkey` param (optional):
	{
		if obj.NewSupplyElgamalPubkey == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.NewSupplyElgamalPubkey)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `ProofInstructionOffset` param:
	err = encoder.Encode(obj.ProofInstructionOffset)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RotateSupplyElgamalPubkey) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewSupplyElgamalPubkey` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.NewSupplyElgamalPubkey)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `ProofInstructionOffset`:
	err = decoder.Decode(&obj.ProofInstructionOffset)
	if err != nil {
		return err
	}
	return nil
}

// NewRotateSupplyElgamalPubkeyInstruction declares a new RotateSupplyElgamalPubkey instruction with the provided parameters and accounts.
func NewRotateSupplyElgamalPubkeyInstruction(
	// Parameters:
	newSupplyElgamalPubkey ag_solanago.PublicKey,
	proofInstructionOffset int8,
	// Accounts:
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *RotateSupplyElgamalPubkey {
	return NewRotateSupplyElgamalPubkeyInstructionBuilder().
		SetNewSupplyElgamalPubkey(newSupplyElgamalPubkey).
		SetProofInstructionOffset(proofInstructionOffset).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}

// NewSimpleRotateSupplyElgamalPubkeyInstruction declares a new RotateSupplyElgamalPubkey instruction with the provided parameters and accounts.
func NewSimpleRotateSupplyElgamalPubkeyInstruction(
	// Parameters:
	newSupplyElgamalPubkey ag_solanago.PublicKey,
	proofInstructionOffset int8,
	// Accounts:
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *RotateSupplyElgamalPubkey {
	return NewRotateSupplyElgamalPubkeyInstructionBuilder().
		SetNewSupplyElgamalPubkey(newSupplyElgamalPubkey).
		SetProofInstructionOffset(proofInstructionOffset).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}
