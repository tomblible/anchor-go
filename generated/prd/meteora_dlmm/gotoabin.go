// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// GoToABin is the `go_to_a_bin` instruction.
type GoToABin struct {
	BinId *int32

	// [0] = [WRITE] lb_pair
	//
	// [1] = [] bin_array_bitmap_extension
	//
	// [2] = [] from_bin_array
	//
	// [3] = [] to_bin_array
	//
	// [4] = [] event_authority
	//
	// [5] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewGoToABinInstructionBuilder creates a new `GoToABin` instruction builder.
func NewGoToABinInstructionBuilder() *GoToABin {
	nd := &GoToABin{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[5] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetBinId sets the "bin_id" parameter.
func (inst *GoToABin) SetBinId(bin_id int32) *GoToABin {
	inst.BinId = &bin_id
	return inst
}

// SetLbPairAccount sets the "lb_pair" account.
func (inst *GoToABin) SetLbPairAccount(lbPair ag_solanago.PublicKey) *GoToABin {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lb_pair" account.
func (inst *GoToABin) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBinArrayBitmapExtensionAccount sets the "bin_array_bitmap_extension" account.
func (inst *GoToABin) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *GoToABin {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(binArrayBitmapExtension)
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "bin_array_bitmap_extension" account (optional).
func (inst *GoToABin) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFromBinArrayAccount sets the "from_bin_array" account.
func (inst *GoToABin) SetFromBinArrayAccount(fromBinArray ag_solanago.PublicKey) *GoToABin {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fromBinArray)
	return inst
}

// GetFromBinArrayAccount gets the "from_bin_array" account (optional).
func (inst *GoToABin) GetFromBinArrayAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetToBinArrayAccount sets the "to_bin_array" account.
func (inst *GoToABin) SetToBinArrayAccount(toBinArray ag_solanago.PublicKey) *GoToABin {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(toBinArray)
	return inst
}

// GetToBinArrayAccount gets the "to_bin_array" account (optional).
func (inst *GoToABin) GetToBinArrayAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *GoToABin) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *GoToABin {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *GoToABin) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProgramAccount sets the "program" account.
func (inst *GoToABin) SetProgramAccount(program ag_solanago.PublicKey) *GoToABin {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *GoToABin) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst *GoToABin) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *GoToABin) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *GoToABin {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:6], metas...)
	return inst
}

func (inst *GoToABin) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6:]
}

func (inst GoToABin) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_GoToABin,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst GoToABin) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *GoToABin) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.BinId == nil {
			return errors.New("binId parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 6 {
		return errors.New("accounts slice has wrong length: expected 6 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LbPair is not set")
		}

		// [1] = BinArrayBitmapExtension is optional

		// [2] = FromBinArray is optional

		// [3] = ToBinArray is optional

		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *GoToABin) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("GoToABin")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" BinId", *inst.BinId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   lb_pair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("bin_array_bitmap_extension", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            from_bin_array", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              to_bin_array", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           event_authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                   program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj GoToABin) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BinId` param:
	err = encoder.Encode(obj.BinId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *GoToABin) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BinId`:
	err = decoder.Decode(&obj.BinId)
	if err != nil {
		return err
	}
	return nil
}

// NewGoToABinInstruction declares a new GoToABin instruction with the provided parameters and accounts.
func NewGoToABinInstruction(
	// Parameters:
	bin_id int32,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	fromBinArray ag_solanago.PublicKey,
	toBinArray ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *GoToABin {
	return NewGoToABinInstructionBuilder().
		SetBinId(bin_id).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetFromBinArrayAccount(fromBinArray).
		SetToBinArrayAccount(toBinArray).
		SetEventAuthorityAccount(eventAuthority)
}

// NewSimpleGoToABinInstruction declares a new GoToABin instruction with the provided parameters and accounts.
func NewSimpleGoToABinInstruction(
	// Parameters:
	bin_id int32,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	fromBinArray ag_solanago.PublicKey,
	toBinArray ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey) *GoToABin {
	return NewGoToABinInstructionBuilder().
		SetBinId(bin_id).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetFromBinArrayAccount(fromBinArray).
		SetToBinArrayAccount(toBinArray).
		SetEventAuthorityAccount(eventAuthority)
}
