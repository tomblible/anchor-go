// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package hpnf

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Close the user's position and NFT account. If the NFT mint belongs to token2022, it will also be closed and the funds returned to the NFT owner.
//
// # Arguments
//
// * `ctx` - The context of accounts
//
type ClosePosition struct {

	// [0] = [WRITE, SIGNER] nft_owner
	// ··········· The position nft owner
	//
	// [1] = [WRITE] position_nft_mint
	// ··········· Mint address bound to the personal position.
	//
	// [2] = [WRITE] position_nft_account
	// ··········· User token account where position NFT be minted to
	//
	// [3] = [WRITE] personal_position
	//
	// [4] = [] system_program
	// ··········· System program to close the position state account
	//
	// [5] = [] token_program
	// ··········· Token/Token2022 program to close token/mint account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClosePositionInstructionBuilder creates a new `ClosePosition` instruction builder.
func NewClosePositionInstructionBuilder() *ClosePosition {
	nd := &ClosePosition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(SystemProgram)
	return nd
}

// SetNftOwnerAccount sets the "nft_owner" account.
// The position nft owner
func (inst *ClosePosition) SetNftOwnerAccount(nftOwner ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(nftOwner).WRITE().SIGNER()
	return inst
}

// GetNftOwnerAccount gets the "nft_owner" account.
// The position nft owner
func (inst *ClosePosition) GetNftOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPositionNftMintAccount sets the "position_nft_mint" account.
// Mint address bound to the personal position.
func (inst *ClosePosition) SetPositionNftMintAccount(positionNftMint ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(positionNftMint).WRITE()
	return inst
}

// GetPositionNftMintAccount gets the "position_nft_mint" account.
// Mint address bound to the personal position.
func (inst *ClosePosition) GetPositionNftMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPositionNftAccountAccount sets the "position_nft_account" account.
// User token account where position NFT be minted to
func (inst *ClosePosition) SetPositionNftAccountAccount(positionNftAccount ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(positionNftAccount).WRITE()
	return inst
}

// GetPositionNftAccountAccount gets the "position_nft_account" account.
// User token account where position NFT be minted to
func (inst *ClosePosition) GetPositionNftAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPersonalPositionAccount sets the "personal_position" account.
func (inst *ClosePosition) SetPersonalPositionAccount(personalPosition ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(personalPosition).WRITE()
	return inst
}

// GetPersonalPositionAccount gets the "personal_position" account.
func (inst *ClosePosition) GetPersonalPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
// System program to close the position state account
func (inst *ClosePosition) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
// System program to close the position state account
func (inst *ClosePosition) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "token_program" account.
// Token/Token2022 program to close token/mint account
func (inst *ClosePosition) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
// Token/Token2022 program to close token/mint account
func (inst *ClosePosition) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst *ClosePosition) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *ClosePosition) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *ClosePosition {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:6], metas...)
	return inst
}

func (inst *ClosePosition) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6:]
}

func (inst ClosePosition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClosePosition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClosePosition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClosePosition) Validate() error {
	if len(inst.AccountMetaSlice) < 6 {
		return errors.New("accounts slice has wrong length: expected 6 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.NftOwner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PositionNftMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PositionNftAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PersonalPosition is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *ClosePosition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClosePosition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        nft_owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("position_nft_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    position_nft_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("personal_position", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   system_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    token_program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj ClosePosition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClosePosition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClosePositionInstruction declares a new ClosePosition instruction with the provided parameters and accounts.
func NewClosePositionInstruction(
	// Accounts:
	nftOwner ag_solanago.PublicKey,
	positionNftMint ag_solanago.PublicKey,
	positionNftAccount ag_solanago.PublicKey,
	personalPosition ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *ClosePosition {
	return NewClosePositionInstructionBuilder().
		SetNftOwnerAccount(nftOwner).
		SetPositionNftMintAccount(positionNftMint).
		SetPositionNftAccountAccount(positionNftAccount).
		SetPersonalPositionAccount(personalPosition).
		SetTokenProgramAccount(tokenProgram)
}

// NewSimpleClosePositionInstruction declares a new ClosePosition instruction with the provided parameters and accounts.
func NewSimpleClosePositionInstruction(
	// Accounts:
	nftOwner ag_solanago.PublicKey,
	positionNftMint ag_solanago.PublicKey,
	positionNftAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *ClosePosition {
	personalPosition := MustFindPersonalPositionAddress(positionNftMint)
	return NewClosePositionInstructionBuilder().
		SetNftOwnerAccount(nftOwner).
		SetPositionNftMintAccount(positionNftMint).
		SetPositionNftAccountAccount(positionNftAccount).
		SetPersonalPositionAccount(personalPosition).
		SetTokenProgramAccount(tokenProgram)
}
