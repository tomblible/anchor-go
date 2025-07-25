// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package pump_curve

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Buys tokens from a bonding curve.
type Buy struct {
	Amount     *uint64
	MaxSolCost *uint64

	// [0] = [] global
	//
	// [1] = [WRITE] fee_recipient
	//
	// [2] = [] mint
	//
	// [3] = [WRITE] bonding_curve
	//
	// [4] = [WRITE] associated_bonding_curve
	//
	// [5] = [WRITE] associated_user
	//
	// [6] = [WRITE, SIGNER] user
	//
	// [7] = [] system_program
	//
	// [8] = [] token_program
	//
	// [9] = [WRITE] creator_vault
	//
	// [10] = [] event_authority
	//
	// [11] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func NewBuyInstructionBuilder() *Buy {
	nd := &Buy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(GlobalPDA)
	nd.AccountMetaSlice[7] = ag_solanago.Meta(SystemProgram)
	nd.AccountMetaSlice[8] = ag_solanago.Meta(TokenProgram)
	nd.AccountMetaSlice[10] = ag_solanago.Meta(EventAuthorityPDA)
	nd.AccountMetaSlice[11] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *Buy) SetAmount(amount uint64) *Buy {
	inst.Amount = &amount
	return inst
}

// SetMaxSolCost sets the "max_sol_cost" parameter.
func (inst *Buy) SetMaxSolCost(max_sol_cost uint64) *Buy {
	inst.MaxSolCost = &max_sol_cost
	return inst
}

// SetGlobalAccount sets the "global" account.
func (inst *Buy) SetGlobalAccount(global ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(global)
	return inst
}

// GetGlobalAccount gets the "global" account.
func (inst *Buy) GetGlobalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFeeRecipientAccount sets the "fee_recipient" account.
func (inst *Buy) SetFeeRecipientAccount(feeRecipient ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(feeRecipient).WRITE()
	return inst
}

// GetFeeRecipientAccount gets the "fee_recipient" account.
func (inst *Buy) GetFeeRecipientAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" account.
func (inst *Buy) SetMintAccount(mint ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *Buy) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBondingCurveAccount sets the "bonding_curve" account.
func (inst *Buy) SetBondingCurveAccount(bondingCurve ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bondingCurve).WRITE()
	return inst
}

// GetBondingCurveAccount gets the "bonding_curve" account.
func (inst *Buy) GetBondingCurveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAssociatedBondingCurveAccount sets the "associated_bonding_curve" account.
func (inst *Buy) SetAssociatedBondingCurveAccount(associatedBondingCurve ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(associatedBondingCurve).WRITE()
	return inst
}

// GetAssociatedBondingCurveAccount gets the "associated_bonding_curve" account.
func (inst *Buy) GetAssociatedBondingCurveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAssociatedUserAccount sets the "associated_user" account.
func (inst *Buy) SetAssociatedUserAccount(associatedUser ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(associatedUser).WRITE()
	return inst
}

// GetAssociatedUserAccount gets the "associated_user" account.
func (inst *Buy) GetAssociatedUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetUserAccount sets the "user" account.
func (inst *Buy) SetUserAccount(user ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *Buy) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Buy) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Buy) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *Buy) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *Buy) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetCreatorVaultAccount sets the "creator_vault" account.
func (inst *Buy) SetCreatorVaultAccount(creatorVault ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(creatorVault).WRITE()
	return inst
}

// GetCreatorVaultAccount gets the "creator_vault" account.
func (inst *Buy) GetCreatorVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *Buy) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *Buy) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetProgramAccount sets the "program" account.
func (inst *Buy) SetProgramAccount(program ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Buy) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst *Buy) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *Buy) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *Buy {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:12], metas...)
	return inst
}

func (inst *Buy) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12:]
}

func (inst Buy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Buy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Buy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Buy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("amount parameter is not set")
		}
		if inst.MaxSolCost == nil {
			return errors.New("maxSolCost parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 12 {
		return errors.New("accounts slice has wrong length: expected 12 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Global is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FeeRecipient is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BondingCurve is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AssociatedBondingCurve is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AssociatedUser is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CreatorVault is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *Buy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Buy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("      Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("  MaxSolCost", *inst.MaxSolCost))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  global", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           fee_recipient", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                    mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           bonding_curve", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("associated_bonding_curve", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         associated_user", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                    user", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           creator_vault", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("         event_authority", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                 program", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj Buy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `MaxSolCost` param:
	err = encoder.Encode(obj.MaxSolCost)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Buy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `MaxSolCost`:
	err = decoder.Decode(&obj.MaxSolCost)
	if err != nil {
		return err
	}
	return nil
}

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewBuyInstruction(
	// Parameters:
	amount uint64,
	max_sol_cost uint64,
	feeRecipient ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	bondingCurve ag_solanago.PublicKey,
	associatedBondingCurve ag_solanago.PublicKey,
	associatedUser ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	creatorVault ag_solanago.PublicKey) *Buy {
	return NewBuyInstructionBuilder().
		SetAmount(amount).
		SetMaxSolCost(max_sol_cost).
		SetFeeRecipientAccount(feeRecipient).
		SetMintAccount(mint).
		SetBondingCurveAccount(bondingCurve).
		SetAssociatedBondingCurveAccount(associatedBondingCurve).
		SetAssociatedUserAccount(associatedUser).
		SetUserAccount(user).
		SetCreatorVaultAccount(creatorVault)
}

// NewSimpleBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewSimpleBuyInstruction(
	// Parameters:
	amount uint64,
	max_sol_cost uint64,
	feeRecipient ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	associatedBondingCurve ag_solanago.PublicKey,
	associatedUser ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	creatorVault ag_solanago.PublicKey) *Buy {
	bondingCurve := MustFindBondingCurveAddress(mint)
	return NewBuyInstructionBuilder().
		SetAmount(amount).
		SetMaxSolCost(max_sol_cost).
		SetFeeRecipientAccount(feeRecipient).
		SetMintAccount(mint).
		SetBondingCurveAccount(bondingCurve).
		SetAssociatedBondingCurveAccount(associatedBondingCurve).
		SetAssociatedUserAccount(associatedUser).
		SetUserAccount(user).
		SetCreatorVaultAccount(creatorVault)
}
