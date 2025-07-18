// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_curve

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MigrationDammV2CreateMetadata is the `migration_damm_v2_create_metadata` instruction.
type MigrationDammV2CreateMetadata struct {

	// [0] = [] virtual_pool
	//
	// [1] = [] config
	//
	// [2] = [WRITE] migration_metadata
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] system_program
	//
	// [5] = [] event_authority
	//
	// [6] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMigrationDammV2CreateMetadataInstructionBuilder creates a new `MigrationDammV2CreateMetadata` instruction builder.
func NewMigrationDammV2CreateMetadataInstructionBuilder() *MigrationDammV2CreateMetadata {
	nd := &MigrationDammV2CreateMetadata{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(SystemProgram)
	nd.AccountMetaSlice[5] = ag_solanago.Meta(EventAuthorityPDA)
	nd.AccountMetaSlice[6] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetVirtualPoolAccount sets the "virtual_pool" account.
func (inst *MigrationDammV2CreateMetadata) SetVirtualPoolAccount(virtualPool ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(virtualPool)
	return inst
}

// GetVirtualPoolAccount gets the "virtual_pool" account.
func (inst *MigrationDammV2CreateMetadata) GetVirtualPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *MigrationDammV2CreateMetadata) SetConfigAccount(config ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *MigrationDammV2CreateMetadata) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMigrationMetadataAccount sets the "migration_metadata" account.
func (inst *MigrationDammV2CreateMetadata) SetMigrationMetadataAccount(migrationMetadata ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(migrationMetadata).WRITE()
	return inst
}

// GetMigrationMetadataAccount gets the "migration_metadata" account.
func (inst *MigrationDammV2CreateMetadata) GetMigrationMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *MigrationDammV2CreateMetadata) SetPayerAccount(payer ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *MigrationDammV2CreateMetadata) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *MigrationDammV2CreateMetadata) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *MigrationDammV2CreateMetadata) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *MigrationDammV2CreateMetadata) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *MigrationDammV2CreateMetadata) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetProgramAccount sets the "program" account.
func (inst *MigrationDammV2CreateMetadata) SetProgramAccount(program ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *MigrationDammV2CreateMetadata) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst *MigrationDammV2CreateMetadata) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *MigrationDammV2CreateMetadata) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *MigrationDammV2CreateMetadata {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:7], metas...)
	return inst
}

func (inst *MigrationDammV2CreateMetadata) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7:]
}

func (inst MigrationDammV2CreateMetadata) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MigrationDammV2CreateMetadata,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MigrationDammV2CreateMetadata) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MigrationDammV2CreateMetadata) Validate() error {
	if len(inst.AccountMetaSlice) < 7 {
		return errors.New("accounts slice has wrong length: expected 7 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.VirtualPool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MigrationMetadata is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *MigrationDammV2CreateMetadata) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MigrationDammV2CreateMetadata")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      virtual_pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("migration_metadata", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   event_authority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj MigrationDammV2CreateMetadata) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *MigrationDammV2CreateMetadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewMigrationDammV2CreateMetadataInstruction declares a new MigrationDammV2CreateMetadata instruction with the provided parameters and accounts.
func NewMigrationDammV2CreateMetadataInstruction(
	// Accounts:
	virtualPool ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	migrationMetadata ag_solanago.PublicKey,
	payer ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	return NewMigrationDammV2CreateMetadataInstructionBuilder().
		SetVirtualPoolAccount(virtualPool).
		SetConfigAccount(config).
		SetMigrationMetadataAccount(migrationMetadata).
		SetPayerAccount(payer)
}

// NewSimpleMigrationDammV2CreateMetadataInstruction declares a new MigrationDammV2CreateMetadata instruction with the provided parameters and accounts.
func NewSimpleMigrationDammV2CreateMetadataInstruction(
	// Accounts:
	virtualPool ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	payer ag_solanago.PublicKey) *MigrationDammV2CreateMetadata {
	migrationMetadata := MustFindMigrationMetadataAddress(virtualPool)
	return NewMigrationDammV2CreateMetadataInstructionBuilder().
		SetVirtualPoolAccount(virtualPool).
		SetConfigAccount(config).
		SetMigrationMetadataAccount(migrationMetadata).
		SetPayerAccount(payer)
}
