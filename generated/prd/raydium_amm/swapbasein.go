// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package raydium_amm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SwapBaseIn is the `swap_base_in` instruction.
type SwapBaseIn struct {
	AmountIn         *uint64
	MinimumAmountOut *uint64

	// [0] = [] token_program
	//
	// [1] = [WRITE] amm
	//
	// [2] = [] amm_authority
	//
	// [3] = [WRITE] amm_open_orders
	//
	// [4] = [WRITE] amm_target_orders
	//
	// [5] = [WRITE] pool_coin_token_account
	//
	// [6] = [WRITE] pool_pc_token_account
	//
	// [7] = [] serum_program
	//
	// [8] = [WRITE] serum_market
	//
	// [9] = [WRITE] serum_bids
	//
	// [10] = [WRITE] serum_asks
	//
	// [11] = [WRITE] serum_event_queue
	//
	// [12] = [WRITE] serum_coin_vault_account
	//
	// [13] = [WRITE] serum_pc_vault_account
	//
	// [14] = [] serum_vault_signer
	//
	// [15] = [WRITE] uer_source_token_account
	//
	// [16] = [WRITE] uer_destination_token_account
	//
	// [17] = [SIGNER] user_source_owner
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSwapBaseInInstructionBuilder creates a new `SwapBaseIn` instruction builder.
func NewSwapBaseInInstructionBuilder() *SwapBaseIn {
	nd := &SwapBaseIn{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 18),
	}
	return nd
}

// SetAmountIn sets the "amount_in" parameter.
func (inst *SwapBaseIn) SetAmountIn(amount_in uint64) *SwapBaseIn {
	inst.AmountIn = &amount_in
	return inst
}

// SetMinimumAmountOut sets the "minimum_amount_out" parameter.
func (inst *SwapBaseIn) SetMinimumAmountOut(minimum_amount_out uint64) *SwapBaseIn {
	inst.MinimumAmountOut = &minimum_amount_out
	return inst
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *SwapBaseIn) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *SwapBaseIn) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAmmAccount sets the "amm" account.
func (inst *SwapBaseIn) SetAmmAccount(amm ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(amm).WRITE()
	return inst
}

// GetAmmAccount gets the "amm" account.
func (inst *SwapBaseIn) GetAmmAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAmmAuthorityAccount sets the "amm_authority" account.
func (inst *SwapBaseIn) SetAmmAuthorityAccount(ammAuthority ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ammAuthority)
	return inst
}

// GetAmmAuthorityAccount gets the "amm_authority" account.
func (inst *SwapBaseIn) GetAmmAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAmmOpenOrdersAccount sets the "amm_open_orders" account.
func (inst *SwapBaseIn) SetAmmOpenOrdersAccount(ammOpenOrders ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(ammOpenOrders).WRITE()
	return inst
}

// GetAmmOpenOrdersAccount gets the "amm_open_orders" account.
func (inst *SwapBaseIn) GetAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAmmTargetOrdersAccount sets the "amm_target_orders" account.
func (inst *SwapBaseIn) SetAmmTargetOrdersAccount(ammTargetOrders ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(ammTargetOrders).WRITE()
	return inst
}

// GetAmmTargetOrdersAccount gets the "amm_target_orders" account.
func (inst *SwapBaseIn) GetAmmTargetOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPoolCoinTokenAccountAccount sets the "pool_coin_token_account" account.
func (inst *SwapBaseIn) SetPoolCoinTokenAccountAccount(poolCoinTokenAccount ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(poolCoinTokenAccount).WRITE()
	return inst
}

// GetPoolCoinTokenAccountAccount gets the "pool_coin_token_account" account.
func (inst *SwapBaseIn) GetPoolCoinTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolPcTokenAccountAccount sets the "pool_pc_token_account" account.
func (inst *SwapBaseIn) SetPoolPcTokenAccountAccount(poolPcTokenAccount ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolPcTokenAccount).WRITE()
	return inst
}

// GetPoolPcTokenAccountAccount gets the "pool_pc_token_account" account.
func (inst *SwapBaseIn) GetPoolPcTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSerumProgramAccount sets the "serum_program" account.
func (inst *SwapBaseIn) SetSerumProgramAccount(serumProgram ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(serumProgram)
	return inst
}

// GetSerumProgramAccount gets the "serum_program" account.
func (inst *SwapBaseIn) GetSerumProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSerumMarketAccount sets the "serum_market" account.
func (inst *SwapBaseIn) SetSerumMarketAccount(serumMarket ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(serumMarket).WRITE()
	return inst
}

// GetSerumMarketAccount gets the "serum_market" account.
func (inst *SwapBaseIn) GetSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSerumBidsAccount sets the "serum_bids" account.
func (inst *SwapBaseIn) SetSerumBidsAccount(serumBids ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(serumBids).WRITE()
	return inst
}

// GetSerumBidsAccount gets the "serum_bids" account.
func (inst *SwapBaseIn) GetSerumBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSerumAsksAccount sets the "serum_asks" account.
func (inst *SwapBaseIn) SetSerumAsksAccount(serumAsks ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(serumAsks).WRITE()
	return inst
}

// GetSerumAsksAccount gets the "serum_asks" account.
func (inst *SwapBaseIn) GetSerumAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSerumEventQueueAccount sets the "serum_event_queue" account.
func (inst *SwapBaseIn) SetSerumEventQueueAccount(serumEventQueue ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(serumEventQueue).WRITE()
	return inst
}

// GetSerumEventQueueAccount gets the "serum_event_queue" account.
func (inst *SwapBaseIn) GetSerumEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSerumCoinVaultAccountAccount sets the "serum_coin_vault_account" account.
func (inst *SwapBaseIn) SetSerumCoinVaultAccountAccount(serumCoinVaultAccount ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(serumCoinVaultAccount).WRITE()
	return inst
}

// GetSerumCoinVaultAccountAccount gets the "serum_coin_vault_account" account.
func (inst *SwapBaseIn) GetSerumCoinVaultAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSerumPcVaultAccountAccount sets the "serum_pc_vault_account" account.
func (inst *SwapBaseIn) SetSerumPcVaultAccountAccount(serumPcVaultAccount ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(serumPcVaultAccount).WRITE()
	return inst
}

// GetSerumPcVaultAccountAccount gets the "serum_pc_vault_account" account.
func (inst *SwapBaseIn) GetSerumPcVaultAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSerumVaultSignerAccount sets the "serum_vault_signer" account.
func (inst *SwapBaseIn) SetSerumVaultSignerAccount(serumVaultSigner ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(serumVaultSigner)
	return inst
}

// GetSerumVaultSignerAccount gets the "serum_vault_signer" account.
func (inst *SwapBaseIn) GetSerumVaultSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetUerSourceTokenAccountAccount sets the "uer_source_token_account" account.
func (inst *SwapBaseIn) SetUerSourceTokenAccountAccount(uerSourceTokenAccount ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(uerSourceTokenAccount).WRITE()
	return inst
}

// GetUerSourceTokenAccountAccount gets the "uer_source_token_account" account.
func (inst *SwapBaseIn) GetUerSourceTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetUerDestinationTokenAccountAccount sets the "uer_destination_token_account" account.
func (inst *SwapBaseIn) SetUerDestinationTokenAccountAccount(uerDestinationTokenAccount ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(uerDestinationTokenAccount).WRITE()
	return inst
}

// GetUerDestinationTokenAccountAccount gets the "uer_destination_token_account" account.
func (inst *SwapBaseIn) GetUerDestinationTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetUserSourceOwnerAccount sets the "user_source_owner" account.
func (inst *SwapBaseIn) SetUserSourceOwnerAccount(userSourceOwner ag_solanago.PublicKey) *SwapBaseIn {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(userSourceOwner).SIGNER()
	return inst
}

// GetUserSourceOwnerAccount gets the "user_source_owner" account.
func (inst *SwapBaseIn) GetUserSourceOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

func (inst *SwapBaseIn) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *SwapBaseIn) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *SwapBaseIn {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:18], metas...)
	return inst
}

func (inst *SwapBaseIn) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[18:]
}

func (inst SwapBaseIn) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_SwapBaseIn),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SwapBaseIn) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SwapBaseIn) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AmountIn == nil {
			return errors.New("amountIn parameter is not set")
		}
		if inst.MinimumAmountOut == nil {
			return errors.New("minimumAmountOut parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 18 {
		return errors.New("accounts slice has wrong length: expected 18 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Amm is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AmmAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AmmOpenOrders is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AmmTargetOrders is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PoolCoinTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolPcTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SerumProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SerumMarket is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SerumBids is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SerumAsks is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SerumEventQueue is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SerumCoinVaultAccount is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SerumPcVaultAccount is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SerumVaultSigner is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.UerSourceTokenAccount is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.UerDestinationTokenAccount is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.UserSourceOwner is not set")
		}
	}
	return nil
}

func (inst *SwapBaseIn) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SwapBaseIn")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("          AmountIn", *inst.AmountIn))
						paramsBranch.Child(ag_format.Param("  MinimumAmountOut", *inst.MinimumAmountOut))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=18]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         token_program", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                   amm", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         amm_authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       amm_open_orders", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     amm_target_orders", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      pool_coin_token_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        pool_pc_token_", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         serum_program", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          serum_market", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("            serum_bids", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            serum_asks", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("     serum_event_queue", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("     serum_coin_vault_", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("       serum_pc_vault_", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("    serum_vault_signer", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("     uer_source_token_", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("uer_destination_token_", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("     user_source_owner", inst.AccountMetaSlice.Get(17)))
					})
				})
		})
}

func (obj SwapBaseIn) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AmountIn` param:
	err = encoder.Encode(obj.AmountIn)
	if err != nil {
		return err
	}
	// Serialize `MinimumAmountOut` param:
	err = encoder.Encode(obj.MinimumAmountOut)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SwapBaseIn) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AmountIn`:
	err = decoder.Decode(&obj.AmountIn)
	if err != nil {
		return err
	}
	// Deserialize `MinimumAmountOut`:
	err = decoder.Decode(&obj.MinimumAmountOut)
	if err != nil {
		return err
	}
	return nil
}

// NewSwapBaseInInstruction declares a new SwapBaseIn instruction with the provided parameters and accounts.
func NewSwapBaseInInstruction(
	// Parameters:
	amount_in uint64,
	minimum_amount_out uint64,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	ammTargetOrders ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	serumProgram ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	serumBids ag_solanago.PublicKey,
	serumAsks ag_solanago.PublicKey,
	serumEventQueue ag_solanago.PublicKey,
	serumCoinVaultAccount ag_solanago.PublicKey,
	serumPcVaultAccount ag_solanago.PublicKey,
	serumVaultSigner ag_solanago.PublicKey,
	uerSourceTokenAccount ag_solanago.PublicKey,
	uerDestinationTokenAccount ag_solanago.PublicKey,
	userSourceOwner ag_solanago.PublicKey) *SwapBaseIn {
	return NewSwapBaseInInstructionBuilder().
		SetAmountIn(amount_in).
		SetMinimumAmountOut(minimum_amount_out).
		SetTokenProgramAccount(tokenProgram).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetAmmTargetOrdersAccount(ammTargetOrders).
		SetPoolCoinTokenAccountAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccountAccount(poolPcTokenAccount).
		SetSerumProgramAccount(serumProgram).
		SetSerumMarketAccount(serumMarket).
		SetSerumBidsAccount(serumBids).
		SetSerumAsksAccount(serumAsks).
		SetSerumEventQueueAccount(serumEventQueue).
		SetSerumCoinVaultAccountAccount(serumCoinVaultAccount).
		SetSerumPcVaultAccountAccount(serumPcVaultAccount).
		SetSerumVaultSignerAccount(serumVaultSigner).
		SetUerSourceTokenAccountAccount(uerSourceTokenAccount).
		SetUerDestinationTokenAccountAccount(uerDestinationTokenAccount).
		SetUserSourceOwnerAccount(userSourceOwner)
}

// NewSimpleSwapBaseInInstruction declares a new SwapBaseIn instruction with the provided parameters and accounts.
func NewSimpleSwapBaseInInstruction(
	// Parameters:
	amount_in uint64,
	minimum_amount_out uint64,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	ammTargetOrders ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	serumProgram ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	serumBids ag_solanago.PublicKey,
	serumAsks ag_solanago.PublicKey,
	serumEventQueue ag_solanago.PublicKey,
	serumCoinVaultAccount ag_solanago.PublicKey,
	serumPcVaultAccount ag_solanago.PublicKey,
	serumVaultSigner ag_solanago.PublicKey,
	uerSourceTokenAccount ag_solanago.PublicKey,
	uerDestinationTokenAccount ag_solanago.PublicKey,
	userSourceOwner ag_solanago.PublicKey) *SwapBaseIn {
	return NewSwapBaseInInstructionBuilder().
		SetAmountIn(amount_in).
		SetMinimumAmountOut(minimum_amount_out).
		SetTokenProgramAccount(tokenProgram).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetAmmTargetOrdersAccount(ammTargetOrders).
		SetPoolCoinTokenAccountAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccountAccount(poolPcTokenAccount).
		SetSerumProgramAccount(serumProgram).
		SetSerumMarketAccount(serumMarket).
		SetSerumBidsAccount(serumBids).
		SetSerumAsksAccount(serumAsks).
		SetSerumEventQueueAccount(serumEventQueue).
		SetSerumCoinVaultAccountAccount(serumCoinVaultAccount).
		SetSerumPcVaultAccountAccount(serumPcVaultAccount).
		SetSerumVaultSignerAccount(serumVaultSigner).
		SetUerSourceTokenAccountAccount(uerSourceTokenAccount).
		SetUerDestinationTokenAccountAccount(uerDestinationTokenAccount).
		SetUserSourceOwnerAccount(userSourceOwner)
}
