// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_pools

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Withdraw protocol fee
type WithdrawProtocolFees struct {

	// [0] = [] pool
	// ··········· Pool account (PDA)
	//
	// [1] = [] a_vault_lp
	//
	// [2] = [WRITE] protocol_token_a_fee
	//
	// [3] = [WRITE] protocol_token_b_fee
	//
	// [4] = [WRITE] treasury_token_a
	//
	// [5] = [WRITE] treasury_token_b
	//
	// [6] = [] token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawProtocolFeesInstructionBuilder creates a new `WithdrawProtocolFees` instruction builder.
func NewWithdrawProtocolFeesInstructionBuilder() *WithdrawProtocolFees {
	nd := &WithdrawProtocolFees{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetPoolAccount sets the "pool" account.
// Pool account (PDA)
func (inst *WithdrawProtocolFees) SetPoolAccount(pool ag_solanago.PublicKey) *WithdrawProtocolFees {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool)
	return inst
}

// GetPoolAccount gets the "pool" account.
// Pool account (PDA)
func (inst *WithdrawProtocolFees) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAVaultLpAccount sets the "a_vault_lp" account.
func (inst *WithdrawProtocolFees) SetAVaultLpAccount(aVaultLp ag_solanago.PublicKey) *WithdrawProtocolFees {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(aVaultLp)
	return inst
}

// GetAVaultLpAccount gets the "a_vault_lp" account.
func (inst *WithdrawProtocolFees) GetAVaultLpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProtocolTokenAFeeAccount sets the "protocol_token_a_fee" account.
func (inst *WithdrawProtocolFees) SetProtocolTokenAFeeAccount(protocolTokenAFee ag_solanago.PublicKey) *WithdrawProtocolFees {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(protocolTokenAFee).WRITE()
	return inst
}

// GetProtocolTokenAFeeAccount gets the "protocol_token_a_fee" account.
func (inst *WithdrawProtocolFees) GetProtocolTokenAFeeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProtocolTokenBFeeAccount sets the "protocol_token_b_fee" account.
func (inst *WithdrawProtocolFees) SetProtocolTokenBFeeAccount(protocolTokenBFee ag_solanago.PublicKey) *WithdrawProtocolFees {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(protocolTokenBFee).WRITE()
	return inst
}

// GetProtocolTokenBFeeAccount gets the "protocol_token_b_fee" account.
func (inst *WithdrawProtocolFees) GetProtocolTokenBFeeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTreasuryTokenAAccount sets the "treasury_token_a" account.
func (inst *WithdrawProtocolFees) SetTreasuryTokenAAccount(treasuryTokenA ag_solanago.PublicKey) *WithdrawProtocolFees {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(treasuryTokenA).WRITE()
	return inst
}

// GetTreasuryTokenAAccount gets the "treasury_token_a" account.
func (inst *WithdrawProtocolFees) GetTreasuryTokenAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTreasuryTokenBAccount sets the "treasury_token_b" account.
func (inst *WithdrawProtocolFees) SetTreasuryTokenBAccount(treasuryTokenB ag_solanago.PublicKey) *WithdrawProtocolFees {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(treasuryTokenB).WRITE()
	return inst
}

// GetTreasuryTokenBAccount gets the "treasury_token_b" account.
func (inst *WithdrawProtocolFees) GetTreasuryTokenBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *WithdrawProtocolFees) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawProtocolFees {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *WithdrawProtocolFees) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst *WithdrawProtocolFees) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *WithdrawProtocolFees) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *WithdrawProtocolFees {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:7], metas...)
	return inst
}

func (inst *WithdrawProtocolFees) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7:]
}

func (inst WithdrawProtocolFees) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawProtocolFees,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawProtocolFees) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawProtocolFees) Validate() error {
	if len(inst.AccountMetaSlice) < 7 {
		return errors.New("accounts slice has wrong length: expected 7 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AVaultLp is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ProtocolTokenAFee is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ProtocolTokenBFee is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TreasuryTokenA is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TreasuryTokenB is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *WithdrawProtocolFees) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawProtocolFees")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          a_vault_lp", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("protocol_token_a_fee", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("protocol_token_b_fee", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    treasury_token_a", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    treasury_token_b", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       token_program", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj WithdrawProtocolFees) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *WithdrawProtocolFees) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewWithdrawProtocolFeesInstruction declares a new WithdrawProtocolFees instruction with the provided parameters and accounts.
func NewWithdrawProtocolFeesInstruction(
	// Accounts:
	pool ag_solanago.PublicKey,
	aVaultLp ag_solanago.PublicKey,
	protocolTokenAFee ag_solanago.PublicKey,
	protocolTokenBFee ag_solanago.PublicKey,
	treasuryTokenA ag_solanago.PublicKey,
	treasuryTokenB ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *WithdrawProtocolFees {
	return NewWithdrawProtocolFeesInstructionBuilder().
		SetPoolAccount(pool).
		SetAVaultLpAccount(aVaultLp).
		SetProtocolTokenAFeeAccount(protocolTokenAFee).
		SetProtocolTokenBFeeAccount(protocolTokenBFee).
		SetTreasuryTokenAAccount(treasuryTokenA).
		SetTreasuryTokenBAccount(treasuryTokenB).
		SetTokenProgramAccount(tokenProgram)
}

// NewSimpleWithdrawProtocolFeesInstruction declares a new WithdrawProtocolFees instruction with the provided parameters and accounts.
func NewSimpleWithdrawProtocolFeesInstruction(
	// Accounts:
	pool ag_solanago.PublicKey,
	aVaultLp ag_solanago.PublicKey,
	protocolTokenAFee ag_solanago.PublicKey,
	protocolTokenBFee ag_solanago.PublicKey,
	treasuryTokenA ag_solanago.PublicKey,
	treasuryTokenB ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *WithdrawProtocolFees {
	return NewWithdrawProtocolFeesInstructionBuilder().
		SetPoolAccount(pool).
		SetAVaultLpAccount(aVaultLp).
		SetProtocolTokenAFeeAccount(protocolTokenAFee).
		SetProtocolTokenBFeeAccount(protocolTokenBFee).
		SetTreasuryTokenAAccount(treasuryTokenA).
		SetTreasuryTokenBAccount(treasuryTokenB).
		SetTokenProgramAccount(tokenProgram)
}
