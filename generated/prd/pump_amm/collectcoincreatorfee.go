// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package pump_amm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CollectCoinCreatorFee is the `collect_coin_creator_fee` instruction.
type CollectCoinCreatorFee struct {

	// [0] = [] quote_mint
	//
	// [1] = [] quote_token_program
	//
	// [2] = [SIGNER] coin_creator
	//
	// [3] = [] coin_creator_vault_authority
	//
	// [4] = [WRITE] coin_creator_vault_ata
	//
	// [5] = [WRITE] coin_creator_token_account
	//
	// [6] = [] event_authority
	//
	// [7] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCollectCoinCreatorFeeInstructionBuilder creates a new `CollectCoinCreatorFee` instruction builder.
func NewCollectCoinCreatorFeeInstructionBuilder() *CollectCoinCreatorFee {
	nd := &CollectCoinCreatorFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	nd.AccountMetaSlice[6] = ag_solanago.Meta(EventAuthorityPDA)
	nd.AccountMetaSlice[7] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetQuoteMintAccount sets the "quote_mint" account.
func (inst *CollectCoinCreatorFee) SetQuoteMintAccount(quoteMint ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(quoteMint)
	return inst
}

// GetQuoteMintAccount gets the "quote_mint" account.
func (inst *CollectCoinCreatorFee) GetQuoteMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetQuoteTokenProgramAccount sets the "quote_token_program" account.
func (inst *CollectCoinCreatorFee) SetQuoteTokenProgramAccount(quoteTokenProgram ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(quoteTokenProgram)
	return inst
}

// GetQuoteTokenProgramAccount gets the "quote_token_program" account.
func (inst *CollectCoinCreatorFee) GetQuoteTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetCoinCreatorAccount sets the "coin_creator" account.
func (inst *CollectCoinCreatorFee) SetCoinCreatorAccount(coinCreator ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(coinCreator).SIGNER()
	return inst
}

// GetCoinCreatorAccount gets the "coin_creator" account.
func (inst *CollectCoinCreatorFee) GetCoinCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCoinCreatorVaultAuthorityAccount sets the "coin_creator_vault_authority" account.
func (inst *CollectCoinCreatorFee) SetCoinCreatorVaultAuthorityAccount(coinCreatorVaultAuthority ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(coinCreatorVaultAuthority)
	return inst
}

// GetCoinCreatorVaultAuthorityAccount gets the "coin_creator_vault_authority" account.
func (inst *CollectCoinCreatorFee) GetCoinCreatorVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCoinCreatorVaultAtaAccount sets the "coin_creator_vault_ata" account.
func (inst *CollectCoinCreatorFee) SetCoinCreatorVaultAtaAccount(coinCreatorVaultAta ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(coinCreatorVaultAta).WRITE()
	return inst
}

// GetCoinCreatorVaultAtaAccount gets the "coin_creator_vault_ata" account.
func (inst *CollectCoinCreatorFee) GetCoinCreatorVaultAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetCoinCreatorTokenAccountAccount sets the "coin_creator_token_account" account.
func (inst *CollectCoinCreatorFee) SetCoinCreatorTokenAccountAccount(coinCreatorTokenAccount ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(coinCreatorTokenAccount).WRITE()
	return inst
}

// GetCoinCreatorTokenAccountAccount gets the "coin_creator_token_account" account.
func (inst *CollectCoinCreatorFee) GetCoinCreatorTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *CollectCoinCreatorFee) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *CollectCoinCreatorFee) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetProgramAccount sets the "program" account.
func (inst *CollectCoinCreatorFee) SetProgramAccount(program ag_solanago.PublicKey) *CollectCoinCreatorFee {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *CollectCoinCreatorFee) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst *CollectCoinCreatorFee) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *CollectCoinCreatorFee) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *CollectCoinCreatorFee {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:8], metas...)
	return inst
}

func (inst *CollectCoinCreatorFee) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8:]
}

func (inst CollectCoinCreatorFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CollectCoinCreatorFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CollectCoinCreatorFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CollectCoinCreatorFee) Validate() error {
	if len(inst.AccountMetaSlice) < 8 {
		return errors.New("accounts slice has wrong length: expected 8 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.QuoteMint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.QuoteTokenProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.CoinCreator is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CoinCreatorVaultAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CoinCreatorVaultAta is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CoinCreatorTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *CollectCoinCreatorFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CollectCoinCreatorFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  quote_mint", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         quote_token_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                coin_creator", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("coin_creator_vault_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      coin_creator_vault_ata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         coin_creator_token_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("             event_authority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                     program", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj CollectCoinCreatorFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CollectCoinCreatorFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCollectCoinCreatorFeeInstruction declares a new CollectCoinCreatorFee instruction with the provided parameters and accounts.
func NewCollectCoinCreatorFeeInstruction(
	// Accounts:
	quoteMint ag_solanago.PublicKey,
	quoteTokenProgram ag_solanago.PublicKey,
	coinCreator ag_solanago.PublicKey,
	coinCreatorVaultAuthority ag_solanago.PublicKey,
	coinCreatorVaultAta ag_solanago.PublicKey,
	coinCreatorTokenAccount ag_solanago.PublicKey) *CollectCoinCreatorFee {
	return NewCollectCoinCreatorFeeInstructionBuilder().
		SetQuoteMintAccount(quoteMint).
		SetQuoteTokenProgramAccount(quoteTokenProgram).
		SetCoinCreatorAccount(coinCreator).
		SetCoinCreatorVaultAuthorityAccount(coinCreatorVaultAuthority).
		SetCoinCreatorVaultAtaAccount(coinCreatorVaultAta).
		SetCoinCreatorTokenAccountAccount(coinCreatorTokenAccount)
}

// NewSimpleCollectCoinCreatorFeeInstruction declares a new CollectCoinCreatorFee instruction with the provided parameters and accounts.
func NewSimpleCollectCoinCreatorFeeInstruction(
	// Accounts:
	quoteMint ag_solanago.PublicKey,
	quoteTokenProgram ag_solanago.PublicKey,
	coinCreator ag_solanago.PublicKey,
	coinCreatorVaultAta ag_solanago.PublicKey,
	coinCreatorTokenAccount ag_solanago.PublicKey) *CollectCoinCreatorFee {
	coinCreatorVaultAuthority := MustFindCoinCreatorVaultAuthorityAddress(coinCreator)
	return NewCollectCoinCreatorFeeInstructionBuilder().
		SetQuoteMintAccount(quoteMint).
		SetQuoteTokenProgramAccount(quoteTokenProgram).
		SetCoinCreatorAccount(coinCreator).
		SetCoinCreatorVaultAuthorityAccount(coinCreatorVaultAuthority).
		SetCoinCreatorVaultAtaAccount(coinCreatorVaultAta).
		SetCoinCreatorTokenAccountAccount(coinCreatorTokenAccount)
}
