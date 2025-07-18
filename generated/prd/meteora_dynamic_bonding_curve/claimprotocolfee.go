// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_curve

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ClaimProtocolFee is the `claim_protocol_fee` instruction.
type ClaimProtocolFee struct {

	// [0] = [] pool_authority
	//
	// [1] = [] config
	//
	// [2] = [WRITE] pool
	//
	// [3] = [WRITE] base_vault
	// ··········· The vault token account for input token
	//
	// [4] = [WRITE] quote_vault
	// ··········· The vault token account for output token
	//
	// [5] = [] base_mint
	// ··········· The mint of token a
	//
	// [6] = [] quote_mint
	// ··········· The mint of token b
	//
	// [7] = [WRITE] token_base_account
	// ··········· The treasury token a account
	//
	// [8] = [WRITE] token_quote_account
	// ··········· The treasury token b account
	//
	// [9] = [] claim_fee_operator
	// ··········· Claim fee operator
	//
	// [10] = [SIGNER] operator
	// ··········· Operator
	//
	// [11] = [] token_base_program
	// ··········· Token a program
	//
	// [12] = [] token_quote_program
	// ··········· Token b program
	//
	// [13] = [] event_authority
	//
	// [14] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClaimProtocolFeeInstructionBuilder creates a new `ClaimProtocolFee` instruction builder.
func NewClaimProtocolFeeInstructionBuilder() *ClaimProtocolFee {
	nd := &ClaimProtocolFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(PoolAuthority)
	nd.AccountMetaSlice[13] = ag_solanago.Meta(EventAuthorityPDA)
	nd.AccountMetaSlice[14] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetPoolAuthorityAccount sets the "pool_authority" account.
func (inst *ClaimProtocolFee) SetPoolAuthorityAccount(poolAuthority ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(poolAuthority)
	return inst
}

// GetPoolAuthorityAccount gets the "pool_authority" account.
func (inst *ClaimProtocolFee) GetPoolAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *ClaimProtocolFee) SetConfigAccount(config ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *ClaimProtocolFee) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPoolAccount sets the "pool" account.
func (inst *ClaimProtocolFee) SetPoolAccount(pool ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *ClaimProtocolFee) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBaseVaultAccount sets the "base_vault" account.
// The vault token account for input token
func (inst *ClaimProtocolFee) SetBaseVaultAccount(baseVault ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(baseVault).WRITE()
	return inst
}

// GetBaseVaultAccount gets the "base_vault" account.
// The vault token account for input token
func (inst *ClaimProtocolFee) GetBaseVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetQuoteVaultAccount sets the "quote_vault" account.
// The vault token account for output token
func (inst *ClaimProtocolFee) SetQuoteVaultAccount(quoteVault ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(quoteVault).WRITE()
	return inst
}

// GetQuoteVaultAccount gets the "quote_vault" account.
// The vault token account for output token
func (inst *ClaimProtocolFee) GetQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetBaseMintAccount sets the "base_mint" account.
// The mint of token a
func (inst *ClaimProtocolFee) SetBaseMintAccount(baseMint ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(baseMint)
	return inst
}

// GetBaseMintAccount gets the "base_mint" account.
// The mint of token a
func (inst *ClaimProtocolFee) GetBaseMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetQuoteMintAccount sets the "quote_mint" account.
// The mint of token b
func (inst *ClaimProtocolFee) SetQuoteMintAccount(quoteMint ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(quoteMint)
	return inst
}

// GetQuoteMintAccount gets the "quote_mint" account.
// The mint of token b
func (inst *ClaimProtocolFee) GetQuoteMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenBaseAccountAccount sets the "token_base_account" account.
// The treasury token a account
func (inst *ClaimProtocolFee) SetTokenBaseAccountAccount(tokenBaseAccount ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenBaseAccount).WRITE()
	return inst
}

// GetTokenBaseAccountAccount gets the "token_base_account" account.
// The treasury token a account
func (inst *ClaimProtocolFee) GetTokenBaseAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenQuoteAccountAccount sets the "token_quote_account" account.
// The treasury token b account
func (inst *ClaimProtocolFee) SetTokenQuoteAccountAccount(tokenQuoteAccount ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenQuoteAccount).WRITE()
	return inst
}

// GetTokenQuoteAccountAccount gets the "token_quote_account" account.
// The treasury token b account
func (inst *ClaimProtocolFee) GetTokenQuoteAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetClaimFeeOperatorAccount sets the "claim_fee_operator" account.
// Claim fee operator
func (inst *ClaimProtocolFee) SetClaimFeeOperatorAccount(claimFeeOperator ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(claimFeeOperator)
	return inst
}

// GetClaimFeeOperatorAccount gets the "claim_fee_operator" account.
// Claim fee operator
func (inst *ClaimProtocolFee) GetClaimFeeOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetOperatorAccount sets the "operator" account.
// Operator
func (inst *ClaimProtocolFee) SetOperatorAccount(operator ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(operator).SIGNER()
	return inst
}

// GetOperatorAccount gets the "operator" account.
// Operator
func (inst *ClaimProtocolFee) GetOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenBaseProgramAccount sets the "token_base_program" account.
// Token a program
func (inst *ClaimProtocolFee) SetTokenBaseProgramAccount(tokenBaseProgram ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenBaseProgram)
	return inst
}

// GetTokenBaseProgramAccount gets the "token_base_program" account.
// Token a program
func (inst *ClaimProtocolFee) GetTokenBaseProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenQuoteProgramAccount sets the "token_quote_program" account.
// Token b program
func (inst *ClaimProtocolFee) SetTokenQuoteProgramAccount(tokenQuoteProgram ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenQuoteProgram)
	return inst
}

// GetTokenQuoteProgramAccount gets the "token_quote_program" account.
// Token b program
func (inst *ClaimProtocolFee) GetTokenQuoteProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *ClaimProtocolFee) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *ClaimProtocolFee) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetProgramAccount sets the "program" account.
func (inst *ClaimProtocolFee) SetProgramAccount(program ag_solanago.PublicKey) *ClaimProtocolFee {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *ClaimProtocolFee) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst *ClaimProtocolFee) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *ClaimProtocolFee) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *ClaimProtocolFee {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:15], metas...)
	return inst
}

func (inst *ClaimProtocolFee) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[15:]
}

func (inst ClaimProtocolFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClaimProtocolFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimProtocolFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimProtocolFee) Validate() error {
	if len(inst.AccountMetaSlice) < 15 {
		return errors.New("accounts slice has wrong length: expected 15 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PoolAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BaseVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.QuoteVault is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.BaseMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.QuoteMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenBaseAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenQuoteAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.ClaimFeeOperator is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Operator is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenBaseProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenQuoteProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *ClaimProtocolFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClaimProtocolFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     pool_authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               pool", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         base_vault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        quote_vault", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          base_mint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         quote_mint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("        token_base_", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       token_quote_", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta(" claim_fee_operator", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("           operator", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta(" token_base_program", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("token_quote_program", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("    event_authority", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("            program", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj ClaimProtocolFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClaimProtocolFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClaimProtocolFeeInstruction declares a new ClaimProtocolFee instruction with the provided parameters and accounts.
func NewClaimProtocolFeeInstruction(
	config ag_solanago.PublicKey,
	pool ag_solanago.PublicKey,
	baseVault ag_solanago.PublicKey,
	quoteVault ag_solanago.PublicKey,
	baseMint ag_solanago.PublicKey,
	quoteMint ag_solanago.PublicKey,
	tokenBaseAccount ag_solanago.PublicKey,
	tokenQuoteAccount ag_solanago.PublicKey,
	claimFeeOperator ag_solanago.PublicKey,
	operator ag_solanago.PublicKey,
	tokenBaseProgram ag_solanago.PublicKey,
	tokenQuoteProgram ag_solanago.PublicKey) *ClaimProtocolFee {
	return NewClaimProtocolFeeInstructionBuilder().
		SetConfigAccount(config).
		SetPoolAccount(pool).
		SetBaseVaultAccount(baseVault).
		SetQuoteVaultAccount(quoteVault).
		SetBaseMintAccount(baseMint).
		SetQuoteMintAccount(quoteMint).
		SetTokenBaseAccountAccount(tokenBaseAccount).
		SetTokenQuoteAccountAccount(tokenQuoteAccount).
		SetClaimFeeOperatorAccount(claimFeeOperator).
		SetOperatorAccount(operator).
		SetTokenBaseProgramAccount(tokenBaseProgram).
		SetTokenQuoteProgramAccount(tokenQuoteProgram)
}

// NewSimpleClaimProtocolFeeInstruction declares a new ClaimProtocolFee instruction with the provided parameters and accounts.
func NewSimpleClaimProtocolFeeInstruction(
	config ag_solanago.PublicKey,
	pool ag_solanago.PublicKey,
	baseVault ag_solanago.PublicKey,
	quoteVault ag_solanago.PublicKey,
	baseMint ag_solanago.PublicKey,
	quoteMint ag_solanago.PublicKey,
	tokenBaseAccount ag_solanago.PublicKey,
	tokenQuoteAccount ag_solanago.PublicKey,
	claimFeeOperator ag_solanago.PublicKey,
	operator ag_solanago.PublicKey,
	tokenBaseProgram ag_solanago.PublicKey,
	tokenQuoteProgram ag_solanago.PublicKey) *ClaimProtocolFee {
	return NewClaimProtocolFeeInstructionBuilder().
		SetConfigAccount(config).
		SetPoolAccount(pool).
		SetBaseVaultAccount(baseVault).
		SetQuoteVaultAccount(quoteVault).
		SetBaseMintAccount(baseMint).
		SetQuoteMintAccount(quoteMint).
		SetTokenBaseAccountAccount(tokenBaseAccount).
		SetTokenQuoteAccountAccount(tokenQuoteAccount).
		SetClaimFeeOperatorAccount(claimFeeOperator).
		SetOperatorAccount(operator).
		SetTokenBaseProgramAccount(tokenBaseProgram).
		SetTokenQuoteProgramAccount(tokenQuoteProgram)
}
