// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package hpnf

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// #[deprecated(note = "Use `decrease_liquidity_v2` instead.")]
// Decreases liquidity for an existing position
//
// # Arguments
//
// * `ctx` -  The context of accounts
// * `liquidity` - The amount by which liquidity will be decreased
// * `amount_0_min` - The minimum amount of token_0 that should be accounted for the burned liquidity
// * `amount_1_min` - The minimum amount of token_1 that should be accounted for the burned liquidity
//
type DecreaseLiquidity struct {
	Liquidity  *ag_binary.Uint128
	Amount0Min *uint64
	Amount1Min *uint64

	// [0] = [SIGNER] nft_owner
	// ··········· The position owner or delegated authority
	//
	// [1] = [] nft_account
	// ··········· The token account for the tokenized position
	//
	// [2] = [WRITE] personal_position
	// ··········· Decrease liquidity for this position
	//
	// [3] = [WRITE] pool_state
	//
	// [4] = [WRITE] protocol_position
	//
	// [5] = [WRITE] token_vault_0
	// ··········· Token_0 vault
	//
	// [6] = [WRITE] token_vault_1
	// ··········· Token_1 vault
	//
	// [7] = [WRITE] tick_array_lower
	// ··········· Stores init state for the lower tick
	//
	// [8] = [WRITE] tick_array_upper
	// ··········· Stores init state for the upper tick
	//
	// [9] = [WRITE] recipient_token_account_0
	// ··········· The destination token account for receive amount_0
	//
	// [10] = [WRITE] recipient_token_account_1
	// ··········· The destination token account for receive amount_1
	//
	// [11] = [] token_program
	// ··········· SPL program to transfer out tokens
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDecreaseLiquidityInstructionBuilder creates a new `DecreaseLiquidity` instruction builder.
func NewDecreaseLiquidityInstructionBuilder() *DecreaseLiquidity {
	nd := &DecreaseLiquidity{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	nd.AccountMetaSlice[11] = ag_solanago.Meta(TokenProgram)
	return nd
}

// SetLiquidity sets the "liquidity" parameter.
func (inst *DecreaseLiquidity) SetLiquidity(liquidity ag_binary.Uint128) *DecreaseLiquidity {
	inst.Liquidity = &liquidity
	return inst
}

// SetAmount0Min sets the "amount_0_min" parameter.
func (inst *DecreaseLiquidity) SetAmount0Min(amount_0_min uint64) *DecreaseLiquidity {
	inst.Amount0Min = &amount_0_min
	return inst
}

// SetAmount1Min sets the "amount_1_min" parameter.
func (inst *DecreaseLiquidity) SetAmount1Min(amount_1_min uint64) *DecreaseLiquidity {
	inst.Amount1Min = &amount_1_min
	return inst
}

// SetNftOwnerAccount sets the "nft_owner" account.
// The position owner or delegated authority
func (inst *DecreaseLiquidity) SetNftOwnerAccount(nftOwner ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(nftOwner).SIGNER()
	return inst
}

// GetNftOwnerAccount gets the "nft_owner" account.
// The position owner or delegated authority
func (inst *DecreaseLiquidity) GetNftOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetNftAccountAccount sets the "nft_account" account.
// The token account for the tokenized position
func (inst *DecreaseLiquidity) SetNftAccountAccount(nftAccount ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(nftAccount)
	return inst
}

// GetNftAccountAccount gets the "nft_account" account.
// The token account for the tokenized position
func (inst *DecreaseLiquidity) GetNftAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPersonalPositionAccount sets the "personal_position" account.
// Decrease liquidity for this position
func (inst *DecreaseLiquidity) SetPersonalPositionAccount(personalPosition ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(personalPosition).WRITE()
	return inst
}

// GetPersonalPositionAccount gets the "personal_position" account.
// Decrease liquidity for this position
func (inst *DecreaseLiquidity) GetPersonalPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolStateAccount sets the "pool_state" account.
func (inst *DecreaseLiquidity) SetPoolStateAccount(poolState ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "pool_state" account.
func (inst *DecreaseLiquidity) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProtocolPositionAccount sets the "protocol_position" account.
func (inst *DecreaseLiquidity) SetProtocolPositionAccount(protocolPosition ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(protocolPosition).WRITE()
	return inst
}

// GetProtocolPositionAccount gets the "protocol_position" account.
func (inst *DecreaseLiquidity) GetProtocolPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenVault0Account sets the "token_vault_0" account.
// Token_0 vault
func (inst *DecreaseLiquidity) SetTokenVault0Account(tokenVault0 ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenVault0).WRITE()
	return inst
}

// GetTokenVault0Account gets the "token_vault_0" account.
// Token_0 vault
func (inst *DecreaseLiquidity) GetTokenVault0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenVault1Account sets the "token_vault_1" account.
// Token_1 vault
func (inst *DecreaseLiquidity) SetTokenVault1Account(tokenVault1 ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenVault1).WRITE()
	return inst
}

// GetTokenVault1Account gets the "token_vault_1" account.
// Token_1 vault
func (inst *DecreaseLiquidity) GetTokenVault1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTickArrayLowerAccount sets the "tick_array_lower" account.
// Stores init state for the lower tick
func (inst *DecreaseLiquidity) SetTickArrayLowerAccount(tickArrayLower ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tickArrayLower).WRITE()
	return inst
}

// GetTickArrayLowerAccount gets the "tick_array_lower" account.
// Stores init state for the lower tick
func (inst *DecreaseLiquidity) GetTickArrayLowerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTickArrayUpperAccount sets the "tick_array_upper" account.
// Stores init state for the upper tick
func (inst *DecreaseLiquidity) SetTickArrayUpperAccount(tickArrayUpper ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tickArrayUpper).WRITE()
	return inst
}

// GetTickArrayUpperAccount gets the "tick_array_upper" account.
// Stores init state for the upper tick
func (inst *DecreaseLiquidity) GetTickArrayUpperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRecipientTokenAccount0Account sets the "recipient_token_account_0" account.
// The destination token account for receive amount_0
func (inst *DecreaseLiquidity) SetRecipientTokenAccount0Account(recipientTokenAccount0 ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(recipientTokenAccount0).WRITE()
	return inst
}

// GetRecipientTokenAccount0Account gets the "recipient_token_account_0" account.
// The destination token account for receive amount_0
func (inst *DecreaseLiquidity) GetRecipientTokenAccount0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRecipientTokenAccount1Account sets the "recipient_token_account_1" account.
// The destination token account for receive amount_1
func (inst *DecreaseLiquidity) SetRecipientTokenAccount1Account(recipientTokenAccount1 ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(recipientTokenAccount1).WRITE()
	return inst
}

// GetRecipientTokenAccount1Account gets the "recipient_token_account_1" account.
// The destination token account for receive amount_1
func (inst *DecreaseLiquidity) GetRecipientTokenAccount1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "token_program" account.
// SPL program to transfer out tokens
func (inst *DecreaseLiquidity) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DecreaseLiquidity {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
// SPL program to transfer out tokens
func (inst *DecreaseLiquidity) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst *DecreaseLiquidity) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *DecreaseLiquidity) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *DecreaseLiquidity {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:12], metas...)
	return inst
}

func (inst *DecreaseLiquidity) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12:]
}

func (inst DecreaseLiquidity) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DecreaseLiquidity,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DecreaseLiquidity) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DecreaseLiquidity) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Liquidity == nil {
			return errors.New("liquidity parameter is not set")
		}
		if inst.Amount0Min == nil {
			return errors.New("amount0Min parameter is not set")
		}
		if inst.Amount1Min == nil {
			return errors.New("amount1Min parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 12 {
		return errors.New("accounts slice has wrong length: expected 12 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.NftOwner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.NftAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PersonalPosition is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ProtocolPosition is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenVault0 is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenVault1 is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TickArrayLower is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TickArrayUpper is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.RecipientTokenAccount0 is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.RecipientTokenAccount1 is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *DecreaseLiquidity) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DecreaseLiquidity")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   Liquidity", *inst.Liquidity))
						paramsBranch.Child(ag_format.Param("  Amount0Min", *inst.Amount0Min))
						paramsBranch.Child(ag_format.Param("  Amount1Min", *inst.Amount1Min))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                nft_owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                     nft_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        personal_position", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("               pool_state", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        protocol_position", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("            token_vault_0", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            token_vault_1", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         tick_array_lower", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("         tick_array_upper", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("recipient_token_account_0", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("recipient_token_account_1", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("            token_program", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj DecreaseLiquidity) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Liquidity` param:
	err = encoder.Encode(obj.Liquidity)
	if err != nil {
		return err
	}
	// Serialize `Amount0Min` param:
	err = encoder.Encode(obj.Amount0Min)
	if err != nil {
		return err
	}
	// Serialize `Amount1Min` param:
	err = encoder.Encode(obj.Amount1Min)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DecreaseLiquidity) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Liquidity`:
	err = decoder.Decode(&obj.Liquidity)
	if err != nil {
		return err
	}
	// Deserialize `Amount0Min`:
	err = decoder.Decode(&obj.Amount0Min)
	if err != nil {
		return err
	}
	// Deserialize `Amount1Min`:
	err = decoder.Decode(&obj.Amount1Min)
	if err != nil {
		return err
	}
	return nil
}

// NewDecreaseLiquidityInstruction declares a new DecreaseLiquidity instruction with the provided parameters and accounts.
func NewDecreaseLiquidityInstruction(
	// Parameters:
	liquidity ag_binary.Uint128,
	amount_0_min uint64,
	amount_1_min uint64,
	// Accounts:
	nftOwner ag_solanago.PublicKey,
	nftAccount ag_solanago.PublicKey,
	personalPosition ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	protocolPosition ag_solanago.PublicKey,
	tokenVault0 ag_solanago.PublicKey,
	tokenVault1 ag_solanago.PublicKey,
	tickArrayLower ag_solanago.PublicKey,
	tickArrayUpper ag_solanago.PublicKey,
	recipientTokenAccount0 ag_solanago.PublicKey,
	recipientTokenAccount1 ag_solanago.PublicKey) *DecreaseLiquidity {
	return NewDecreaseLiquidityInstructionBuilder().
		SetLiquidity(liquidity).
		SetAmount0Min(amount_0_min).
		SetAmount1Min(amount_1_min).
		SetNftOwnerAccount(nftOwner).
		SetNftAccountAccount(nftAccount).
		SetPersonalPositionAccount(personalPosition).
		SetPoolStateAccount(poolState).
		SetProtocolPositionAccount(protocolPosition).
		SetTokenVault0Account(tokenVault0).
		SetTokenVault1Account(tokenVault1).
		SetTickArrayLowerAccount(tickArrayLower).
		SetTickArrayUpperAccount(tickArrayUpper).
		SetRecipientTokenAccount0Account(recipientTokenAccount0).
		SetRecipientTokenAccount1Account(recipientTokenAccount1)
}

// NewSimpleDecreaseLiquidityInstruction declares a new DecreaseLiquidity instruction with the provided parameters and accounts.
func NewSimpleDecreaseLiquidityInstruction(
	// Parameters:
	liquidity ag_binary.Uint128,
	amount_0_min uint64,
	amount_1_min uint64,
	// Accounts:
	nftOwner ag_solanago.PublicKey,
	nftAccount ag_solanago.PublicKey,
	personalPosition ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	protocolPosition ag_solanago.PublicKey,
	tokenVault0 ag_solanago.PublicKey,
	tokenVault1 ag_solanago.PublicKey,
	tickArrayLower ag_solanago.PublicKey,
	tickArrayUpper ag_solanago.PublicKey,
	recipientTokenAccount0 ag_solanago.PublicKey,
	recipientTokenAccount1 ag_solanago.PublicKey) *DecreaseLiquidity {
	return NewDecreaseLiquidityInstructionBuilder().
		SetLiquidity(liquidity).
		SetAmount0Min(amount_0_min).
		SetAmount1Min(amount_1_min).
		SetNftOwnerAccount(nftOwner).
		SetNftAccountAccount(nftAccount).
		SetPersonalPositionAccount(personalPosition).
		SetPoolStateAccount(poolState).
		SetProtocolPositionAccount(protocolPosition).
		SetTokenVault0Account(tokenVault0).
		SetTokenVault1Account(tokenVault1).
		SetTickArrayLowerAccount(tickArrayLower).
		SetTickArrayUpperAccount(tickArrayUpper).
		SetRecipientTokenAccount0Account(recipientTokenAccount0).
		SetRecipientTokenAccount1Account(recipientTokenAccount1)
}
