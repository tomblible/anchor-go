// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package raydium_amm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MonitorStep is the `monitor_step` instruction.
type MonitorStep struct {
	PlanOrderLimit   *uint16
	PlaceOrderLimit  *uint16
	CancelOrderLimit *uint16

	// [0] = [] token_program
	//
	// [1] = [] rent
	//
	// [2] = [] clock
	//
	// [3] = [WRITE] amm
	//
	// [4] = [] amm_authority
	//
	// [5] = [WRITE] amm_open_orders
	//
	// [6] = [WRITE] amm_target_orders
	//
	// [7] = [WRITE] pool_coin_token_account
	//
	// [8] = [WRITE] pool_pc_token_account
	//
	// [9] = [WRITE] pool_withdraw_queue
	//
	// [10] = [] serum_program
	//
	// [11] = [WRITE] serum_market
	//
	// [12] = [WRITE] serum_coin_vault_account
	//
	// [13] = [WRITE] serum_pc_vault_account
	//
	// [14] = [] serum_vault_signer
	//
	// [15] = [WRITE] serum_req_q
	//
	// [16] = [WRITE] serum_event_q
	//
	// [17] = [WRITE] serum_bids
	//
	// [18] = [WRITE] serum_asks
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMonitorStepInstructionBuilder creates a new `MonitorStep` instruction builder.
func NewMonitorStepInstructionBuilder() *MonitorStep {
	nd := &MonitorStep{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 19),
	}
	return nd
}

// SetPlanOrderLimit sets the "plan_order_limit" parameter.
func (inst *MonitorStep) SetPlanOrderLimit(plan_order_limit uint16) *MonitorStep {
	inst.PlanOrderLimit = &plan_order_limit
	return inst
}

// SetPlaceOrderLimit sets the "place_order_limit" parameter.
func (inst *MonitorStep) SetPlaceOrderLimit(place_order_limit uint16) *MonitorStep {
	inst.PlaceOrderLimit = &place_order_limit
	return inst
}

// SetCancelOrderLimit sets the "cancel_order_limit" parameter.
func (inst *MonitorStep) SetCancelOrderLimit(cancel_order_limit uint16) *MonitorStep {
	inst.CancelOrderLimit = &cancel_order_limit
	return inst
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *MonitorStep) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *MonitorStep) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRentAccount sets the "rent" account.
func (inst *MonitorStep) SetRentAccount(rent ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *MonitorStep) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetClockAccount sets the "clock" account.
func (inst *MonitorStep) SetClockAccount(clock ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *MonitorStep) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAmmAccount sets the "amm" account.
func (inst *MonitorStep) SetAmmAccount(amm ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(amm).WRITE()
	return inst
}

// GetAmmAccount gets the "amm" account.
func (inst *MonitorStep) GetAmmAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAmmAuthorityAccount sets the "amm_authority" account.
func (inst *MonitorStep) SetAmmAuthorityAccount(ammAuthority ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(ammAuthority)
	return inst
}

// GetAmmAuthorityAccount gets the "amm_authority" account.
func (inst *MonitorStep) GetAmmAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAmmOpenOrdersAccount sets the "amm_open_orders" account.
func (inst *MonitorStep) SetAmmOpenOrdersAccount(ammOpenOrders ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(ammOpenOrders).WRITE()
	return inst
}

// GetAmmOpenOrdersAccount gets the "amm_open_orders" account.
func (inst *MonitorStep) GetAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAmmTargetOrdersAccount sets the "amm_target_orders" account.
func (inst *MonitorStep) SetAmmTargetOrdersAccount(ammTargetOrders ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(ammTargetOrders).WRITE()
	return inst
}

// GetAmmTargetOrdersAccount gets the "amm_target_orders" account.
func (inst *MonitorStep) GetAmmTargetOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetPoolCoinTokenAccountAccount sets the "pool_coin_token_account" account.
func (inst *MonitorStep) SetPoolCoinTokenAccountAccount(poolCoinTokenAccount ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(poolCoinTokenAccount).WRITE()
	return inst
}

// GetPoolCoinTokenAccountAccount gets the "pool_coin_token_account" account.
func (inst *MonitorStep) GetPoolCoinTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetPoolPcTokenAccountAccount sets the "pool_pc_token_account" account.
func (inst *MonitorStep) SetPoolPcTokenAccountAccount(poolPcTokenAccount ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(poolPcTokenAccount).WRITE()
	return inst
}

// GetPoolPcTokenAccountAccount gets the "pool_pc_token_account" account.
func (inst *MonitorStep) GetPoolPcTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetPoolWithdrawQueueAccount sets the "pool_withdraw_queue" account.
func (inst *MonitorStep) SetPoolWithdrawQueueAccount(poolWithdrawQueue ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(poolWithdrawQueue).WRITE()
	return inst
}

// GetPoolWithdrawQueueAccount gets the "pool_withdraw_queue" account.
func (inst *MonitorStep) GetPoolWithdrawQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSerumProgramAccount sets the "serum_program" account.
func (inst *MonitorStep) SetSerumProgramAccount(serumProgram ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(serumProgram)
	return inst
}

// GetSerumProgramAccount gets the "serum_program" account.
func (inst *MonitorStep) GetSerumProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSerumMarketAccount sets the "serum_market" account.
func (inst *MonitorStep) SetSerumMarketAccount(serumMarket ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(serumMarket).WRITE()
	return inst
}

// GetSerumMarketAccount gets the "serum_market" account.
func (inst *MonitorStep) GetSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSerumCoinVaultAccountAccount sets the "serum_coin_vault_account" account.
func (inst *MonitorStep) SetSerumCoinVaultAccountAccount(serumCoinVaultAccount ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(serumCoinVaultAccount).WRITE()
	return inst
}

// GetSerumCoinVaultAccountAccount gets the "serum_coin_vault_account" account.
func (inst *MonitorStep) GetSerumCoinVaultAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSerumPcVaultAccountAccount sets the "serum_pc_vault_account" account.
func (inst *MonitorStep) SetSerumPcVaultAccountAccount(serumPcVaultAccount ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(serumPcVaultAccount).WRITE()
	return inst
}

// GetSerumPcVaultAccountAccount gets the "serum_pc_vault_account" account.
func (inst *MonitorStep) GetSerumPcVaultAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSerumVaultSignerAccount sets the "serum_vault_signer" account.
func (inst *MonitorStep) SetSerumVaultSignerAccount(serumVaultSigner ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(serumVaultSigner)
	return inst
}

// GetSerumVaultSignerAccount gets the "serum_vault_signer" account.
func (inst *MonitorStep) GetSerumVaultSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSerumReqQAccount sets the "serum_req_q" account.
func (inst *MonitorStep) SetSerumReqQAccount(serumReqQ ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(serumReqQ).WRITE()
	return inst
}

// GetSerumReqQAccount gets the "serum_req_q" account.
func (inst *MonitorStep) GetSerumReqQAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetSerumEventQAccount sets the "serum_event_q" account.
func (inst *MonitorStep) SetSerumEventQAccount(serumEventQ ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(serumEventQ).WRITE()
	return inst
}

// GetSerumEventQAccount gets the "serum_event_q" account.
func (inst *MonitorStep) GetSerumEventQAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetSerumBidsAccount sets the "serum_bids" account.
func (inst *MonitorStep) SetSerumBidsAccount(serumBids ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(serumBids).WRITE()
	return inst
}

// GetSerumBidsAccount gets the "serum_bids" account.
func (inst *MonitorStep) GetSerumBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetSerumAsksAccount sets the "serum_asks" account.
func (inst *MonitorStep) SetSerumAsksAccount(serumAsks ag_solanago.PublicKey) *MonitorStep {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(serumAsks).WRITE()
	return inst
}

// GetSerumAsksAccount gets the "serum_asks" account.
func (inst *MonitorStep) GetSerumAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

func (inst *MonitorStep) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *MonitorStep) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *MonitorStep {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:19], metas...)
	return inst
}

func (inst *MonitorStep) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[19:]
}

func (inst MonitorStep) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_MonitorStep),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MonitorStep) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MonitorStep) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PlanOrderLimit == nil {
			return errors.New("planOrderLimit parameter is not set")
		}
		if inst.PlaceOrderLimit == nil {
			return errors.New("placeOrderLimit parameter is not set")
		}
		if inst.CancelOrderLimit == nil {
			return errors.New("cancelOrderLimit parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 19 {
		return errors.New("accounts slice has wrong length: expected 19 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Clock is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Amm is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AmmAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AmmOpenOrders is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.AmmTargetOrders is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.PoolCoinTokenAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.PoolPcTokenAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.PoolWithdrawQueue is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SerumProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SerumMarket is not set")
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
			return errors.New("accounts.SerumReqQ is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.SerumEventQ is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.SerumBids is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.SerumAsks is not set")
		}
	}
	return nil
}

func (inst *MonitorStep) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MonitorStep")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    PlanOrderLimit", *inst.PlanOrderLimit))
						paramsBranch.Child(ag_format.Param("   PlaceOrderLimit", *inst.PlaceOrderLimit))
						paramsBranch.Child(ag_format.Param("  CancelOrderLimit", *inst.CancelOrderLimit))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=19]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      token_program", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               rent", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              clock", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                amm", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      amm_authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    amm_open_orders", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  amm_target_orders", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("   pool_coin_token_", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     pool_pc_token_", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("pool_withdraw_queue", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("      serum_program", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       serum_market", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("  serum_coin_vault_", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("    serum_pc_vault_", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta(" serum_vault_signer", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("        serum_req_q", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("      serum_event_q", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("         serum_bids", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("         serum_asks", inst.AccountMetaSlice.Get(18)))
					})
				})
		})
}

func (obj MonitorStep) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PlanOrderLimit` param:
	err = encoder.Encode(obj.PlanOrderLimit)
	if err != nil {
		return err
	}
	// Serialize `PlaceOrderLimit` param:
	err = encoder.Encode(obj.PlaceOrderLimit)
	if err != nil {
		return err
	}
	// Serialize `CancelOrderLimit` param:
	err = encoder.Encode(obj.CancelOrderLimit)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MonitorStep) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PlanOrderLimit`:
	err = decoder.Decode(&obj.PlanOrderLimit)
	if err != nil {
		return err
	}
	// Deserialize `PlaceOrderLimit`:
	err = decoder.Decode(&obj.PlaceOrderLimit)
	if err != nil {
		return err
	}
	// Deserialize `CancelOrderLimit`:
	err = decoder.Decode(&obj.CancelOrderLimit)
	if err != nil {
		return err
	}
	return nil
}

// NewMonitorStepInstruction declares a new MonitorStep instruction with the provided parameters and accounts.
func NewMonitorStepInstruction(
	// Parameters:
	plan_order_limit uint16,
	place_order_limit uint16,
	cancel_order_limit uint16,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	ammTargetOrders ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	poolWithdrawQueue ag_solanago.PublicKey,
	serumProgram ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	serumCoinVaultAccount ag_solanago.PublicKey,
	serumPcVaultAccount ag_solanago.PublicKey,
	serumVaultSigner ag_solanago.PublicKey,
	serumReqQ ag_solanago.PublicKey,
	serumEventQ ag_solanago.PublicKey,
	serumBids ag_solanago.PublicKey,
	serumAsks ag_solanago.PublicKey) *MonitorStep {
	return NewMonitorStepInstructionBuilder().
		SetPlanOrderLimit(plan_order_limit).
		SetPlaceOrderLimit(place_order_limit).
		SetCancelOrderLimit(cancel_order_limit).
		SetTokenProgramAccount(tokenProgram).
		SetRentAccount(rent).
		SetClockAccount(clock).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetAmmTargetOrdersAccount(ammTargetOrders).
		SetPoolCoinTokenAccountAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccountAccount(poolPcTokenAccount).
		SetPoolWithdrawQueueAccount(poolWithdrawQueue).
		SetSerumProgramAccount(serumProgram).
		SetSerumMarketAccount(serumMarket).
		SetSerumCoinVaultAccountAccount(serumCoinVaultAccount).
		SetSerumPcVaultAccountAccount(serumPcVaultAccount).
		SetSerumVaultSignerAccount(serumVaultSigner).
		SetSerumReqQAccount(serumReqQ).
		SetSerumEventQAccount(serumEventQ).
		SetSerumBidsAccount(serumBids).
		SetSerumAsksAccount(serumAsks)
}

// NewSimpleMonitorStepInstruction declares a new MonitorStep instruction with the provided parameters and accounts.
func NewSimpleMonitorStepInstruction(
	// Parameters:
	plan_order_limit uint16,
	place_order_limit uint16,
	cancel_order_limit uint16,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	ammTargetOrders ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	poolWithdrawQueue ag_solanago.PublicKey,
	serumProgram ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	serumCoinVaultAccount ag_solanago.PublicKey,
	serumPcVaultAccount ag_solanago.PublicKey,
	serumVaultSigner ag_solanago.PublicKey,
	serumReqQ ag_solanago.PublicKey,
	serumEventQ ag_solanago.PublicKey,
	serumBids ag_solanago.PublicKey,
	serumAsks ag_solanago.PublicKey) *MonitorStep {
	return NewMonitorStepInstructionBuilder().
		SetPlanOrderLimit(plan_order_limit).
		SetPlaceOrderLimit(place_order_limit).
		SetCancelOrderLimit(cancel_order_limit).
		SetTokenProgramAccount(tokenProgram).
		SetRentAccount(rent).
		SetClockAccount(clock).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetAmmTargetOrdersAccount(ammTargetOrders).
		SetPoolCoinTokenAccountAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccountAccount(poolPcTokenAccount).
		SetPoolWithdrawQueueAccount(poolWithdrawQueue).
		SetSerumProgramAccount(serumProgram).
		SetSerumMarketAccount(serumMarket).
		SetSerumCoinVaultAccountAccount(serumCoinVaultAccount).
		SetSerumPcVaultAccountAccount(serumPcVaultAccount).
		SetSerumVaultSignerAccount(serumVaultSigner).
		SetSerumReqQAccount(serumReqQ).
		SetSerumEventQAccount(serumEventQ).
		SetSerumBidsAccount(serumBids).
		SetSerumAsksAccount(serumAsks)
}
