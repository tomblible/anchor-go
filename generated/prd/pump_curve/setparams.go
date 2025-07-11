// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package pump_curve

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Sets the global state parameters.
type SetParams struct {
	InitialVirtualTokenReserves *uint64
	InitialVirtualSolReserves   *uint64
	InitialRealTokenReserves    *uint64
	TokenTotalSupply            *uint64
	FeeBasisPoints              *uint64
	WithdrawAuthority           *ag_solanago.PublicKey
	EnableMigrate               *bool
	PoolMigrationFee            *uint64
	CreatorFeeBasisPoints       *uint64
	SetCreatorAuthority         *ag_solanago.PublicKey

	// [0] = [WRITE] global
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [] event_authority
	//
	// [3] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetParamsInstructionBuilder creates a new `SetParams` instruction builder.
func NewSetParamsInstructionBuilder() *SetParams {
	nd := &SetParams{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(GlobalPDA).WRITE()
	nd.AccountMetaSlice[2] = ag_solanago.Meta(EventAuthorityPDA)
	nd.AccountMetaSlice[3] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetInitialVirtualTokenReserves sets the "initial_virtual_token_reserves" parameter.
func (inst *SetParams) SetInitialVirtualTokenReserves(initial_virtual_token_reserves uint64) *SetParams {
	inst.InitialVirtualTokenReserves = &initial_virtual_token_reserves
	return inst
}

// SetInitialVirtualSolReserves sets the "initial_virtual_sol_reserves" parameter.
func (inst *SetParams) SetInitialVirtualSolReserves(initial_virtual_sol_reserves uint64) *SetParams {
	inst.InitialVirtualSolReserves = &initial_virtual_sol_reserves
	return inst
}

// SetInitialRealTokenReserves sets the "initial_real_token_reserves" parameter.
func (inst *SetParams) SetInitialRealTokenReserves(initial_real_token_reserves uint64) *SetParams {
	inst.InitialRealTokenReserves = &initial_real_token_reserves
	return inst
}

// SetTokenTotalSupply sets the "token_total_supply" parameter.
func (inst *SetParams) SetTokenTotalSupply(token_total_supply uint64) *SetParams {
	inst.TokenTotalSupply = &token_total_supply
	return inst
}

// SetFeeBasisPoints sets the "fee_basis_points" parameter.
func (inst *SetParams) SetFeeBasisPoints(fee_basis_points uint64) *SetParams {
	inst.FeeBasisPoints = &fee_basis_points
	return inst
}

// SetWithdrawAuthority sets the "withdraw_authority" parameter.
func (inst *SetParams) SetWithdrawAuthority(withdraw_authority ag_solanago.PublicKey) *SetParams {
	inst.WithdrawAuthority = &withdraw_authority
	return inst
}

// SetEnableMigrate sets the "enable_migrate" parameter.
func (inst *SetParams) SetEnableMigrate(enable_migrate bool) *SetParams {
	inst.EnableMigrate = &enable_migrate
	return inst
}

// SetPoolMigrationFee sets the "pool_migration_fee" parameter.
func (inst *SetParams) SetPoolMigrationFee(pool_migration_fee uint64) *SetParams {
	inst.PoolMigrationFee = &pool_migration_fee
	return inst
}

// SetCreatorFeeBasisPoints sets the "creator_fee_basis_points" parameter.
func (inst *SetParams) SetCreatorFeeBasisPoints(creator_fee_basis_points uint64) *SetParams {
	inst.CreatorFeeBasisPoints = &creator_fee_basis_points
	return inst
}

// SetSetCreatorAuthority sets the "set_creator_authority" parameter.
func (inst *SetParams) SetSetCreatorAuthority(set_creator_authority ag_solanago.PublicKey) *SetParams {
	inst.SetCreatorAuthority = &set_creator_authority
	return inst
}

// SetGlobalAccount sets the "global" account.
func (inst *SetParams) SetGlobalAccount(global ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(global).WRITE()
	return inst
}

// GetGlobalAccount gets the "global" account.
func (inst *SetParams) GetGlobalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetParams) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetParams) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *SetParams) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *SetParams) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProgramAccount sets the "program" account.
func (inst *SetParams) SetProgramAccount(program ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *SetParams) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst *SetParams) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *SetParams) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *SetParams {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:4], metas...)
	return inst
}

func (inst *SetParams) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4:]
}

func (inst SetParams) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetParams,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetParams) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetParams) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.InitialVirtualTokenReserves == nil {
			return errors.New("initialVirtualTokenReserves parameter is not set")
		}
		if inst.InitialVirtualSolReserves == nil {
			return errors.New("initialVirtualSolReserves parameter is not set")
		}
		if inst.InitialRealTokenReserves == nil {
			return errors.New("initialRealTokenReserves parameter is not set")
		}
		if inst.TokenTotalSupply == nil {
			return errors.New("tokenTotalSupply parameter is not set")
		}
		if inst.FeeBasisPoints == nil {
			return errors.New("feeBasisPoints parameter is not set")
		}
		if inst.WithdrawAuthority == nil {
			return errors.New("withdrawAuthority parameter is not set")
		}
		if inst.EnableMigrate == nil {
			return errors.New("enableMigrate parameter is not set")
		}
		if inst.PoolMigrationFee == nil {
			return errors.New("poolMigrationFee parameter is not set")
		}
		if inst.CreatorFeeBasisPoints == nil {
			return errors.New("creatorFeeBasisPoints parameter is not set")
		}
		if inst.SetCreatorAuthority == nil {
			return errors.New("setCreatorAuthority parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 4 {
		return errors.New("accounts slice has wrong length: expected 4 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Global is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *SetParams) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetParams")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=10]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   InitialVirtualTokenReserves", *inst.InitialVirtualTokenReserves))
						paramsBranch.Child(ag_format.Param("     InitialVirtualSolReserves", *inst.InitialVirtualSolReserves))
						paramsBranch.Child(ag_format.Param("      InitialRealTokenReserves", *inst.InitialRealTokenReserves))
						paramsBranch.Child(ag_format.Param("              TokenTotalSupply", *inst.TokenTotalSupply))
						paramsBranch.Child(ag_format.Param("                FeeBasisPoints", *inst.FeeBasisPoints))
						paramsBranch.Child(ag_format.Param("             WithdrawAuthority", *inst.WithdrawAuthority))
						paramsBranch.Child(ag_format.Param("                 EnableMigrate", *inst.EnableMigrate))
						paramsBranch.Child(ag_format.Param("              PoolMigrationFee", *inst.PoolMigrationFee))
						paramsBranch.Child(ag_format.Param("         CreatorFeeBasisPoints", *inst.CreatorFeeBasisPoints))
						paramsBranch.Child(ag_format.Param("           SetCreatorAuthority", *inst.SetCreatorAuthority))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         global", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("event_authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        program", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj SetParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `InitialVirtualTokenReserves` param:
	err = encoder.Encode(obj.InitialVirtualTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `InitialVirtualSolReserves` param:
	err = encoder.Encode(obj.InitialVirtualSolReserves)
	if err != nil {
		return err
	}
	// Serialize `InitialRealTokenReserves` param:
	err = encoder.Encode(obj.InitialRealTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `TokenTotalSupply` param:
	err = encoder.Encode(obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Serialize `FeeBasisPoints` param:
	err = encoder.Encode(obj.FeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `WithdrawAuthority` param:
	err = encoder.Encode(obj.WithdrawAuthority)
	if err != nil {
		return err
	}
	// Serialize `EnableMigrate` param:
	err = encoder.Encode(obj.EnableMigrate)
	if err != nil {
		return err
	}
	// Serialize `PoolMigrationFee` param:
	err = encoder.Encode(obj.PoolMigrationFee)
	if err != nil {
		return err
	}
	// Serialize `CreatorFeeBasisPoints` param:
	err = encoder.Encode(obj.CreatorFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `SetCreatorAuthority` param:
	err = encoder.Encode(obj.SetCreatorAuthority)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `InitialVirtualTokenReserves`:
	err = decoder.Decode(&obj.InitialVirtualTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `InitialVirtualSolReserves`:
	err = decoder.Decode(&obj.InitialVirtualSolReserves)
	if err != nil {
		return err
	}
	// Deserialize `InitialRealTokenReserves`:
	err = decoder.Decode(&obj.InitialRealTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `TokenTotalSupply`:
	err = decoder.Decode(&obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Deserialize `FeeBasisPoints`:
	err = decoder.Decode(&obj.FeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawAuthority`:
	err = decoder.Decode(&obj.WithdrawAuthority)
	if err != nil {
		return err
	}
	// Deserialize `EnableMigrate`:
	err = decoder.Decode(&obj.EnableMigrate)
	if err != nil {
		return err
	}
	// Deserialize `PoolMigrationFee`:
	err = decoder.Decode(&obj.PoolMigrationFee)
	if err != nil {
		return err
	}
	// Deserialize `CreatorFeeBasisPoints`:
	err = decoder.Decode(&obj.CreatorFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `SetCreatorAuthority`:
	err = decoder.Decode(&obj.SetCreatorAuthority)
	if err != nil {
		return err
	}
	return nil
}

// NewSetParamsInstruction declares a new SetParams instruction with the provided parameters and accounts.
func NewSetParamsInstruction(
	// Parameters:
	initial_virtual_token_reserves uint64,
	initial_virtual_sol_reserves uint64,
	initial_real_token_reserves uint64,
	token_total_supply uint64,
	fee_basis_points uint64,
	withdraw_authority ag_solanago.PublicKey,
	enable_migrate bool,
	pool_migration_fee uint64,
	creator_fee_basis_points uint64,
	set_creator_authority ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetParams {
	return NewSetParamsInstructionBuilder().
		SetInitialVirtualTokenReserves(initial_virtual_token_reserves).
		SetInitialVirtualSolReserves(initial_virtual_sol_reserves).
		SetInitialRealTokenReserves(initial_real_token_reserves).
		SetTokenTotalSupply(token_total_supply).
		SetFeeBasisPoints(fee_basis_points).
		SetWithdrawAuthority(withdraw_authority).
		SetEnableMigrate(enable_migrate).
		SetPoolMigrationFee(pool_migration_fee).
		SetCreatorFeeBasisPoints(creator_fee_basis_points).
		SetSetCreatorAuthority(set_creator_authority).
		SetAuthorityAccount(authority)
}

// NewSimpleSetParamsInstruction declares a new SetParams instruction with the provided parameters and accounts.
func NewSimpleSetParamsInstruction(
	// Parameters:
	initial_virtual_token_reserves uint64,
	initial_virtual_sol_reserves uint64,
	initial_real_token_reserves uint64,
	token_total_supply uint64,
	fee_basis_points uint64,
	withdraw_authority ag_solanago.PublicKey,
	enable_migrate bool,
	pool_migration_fee uint64,
	creator_fee_basis_points uint64,
	set_creator_authority ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetParams {
	return NewSetParamsInstructionBuilder().
		SetInitialVirtualTokenReserves(initial_virtual_token_reserves).
		SetInitialVirtualSolReserves(initial_virtual_sol_reserves).
		SetInitialRealTokenReserves(initial_real_token_reserves).
		SetTokenTotalSupply(token_total_supply).
		SetFeeBasisPoints(fee_basis_points).
		SetWithdrawAuthority(withdraw_authority).
		SetEnableMigrate(enable_migrate).
		SetPoolMigrationFee(pool_migration_fee).
		SetCreatorFeeBasisPoints(creator_fee_basis_points).
		SetSetCreatorAuthority(set_creator_authority).
		SetAuthorityAccount(authority)
}
