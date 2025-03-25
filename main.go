package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gagliardetto/anchor-go/sighash"
	"github.com/gagliardetto/solana-go"

	"github.com/dave/jennifer/jen"
	. "github.com/dave/jennifer/jen"
	bin "github.com/gagliardetto/binary"
	. "github.com/gagliardetto/utilz"

	"golang.org/x/mod/modfile"
)

const generatedDir = "generated"

// TODO:
// - tests where type has field that is a complex enum (represented as an interface): assign a random concrete value from the possible enum variants.
// - when printing tree, check for len before accessing array indexes.

func main() {
	conf.Encoding = EncodingBorsh
	conf.TypeID = TypeIDAnchor

	filenames := FlagStringArray{}
	flag.Var(&filenames, "src", "Path to source; can use multiple times.")
	flag.StringVar(&conf.DstDir, "dst", generatedDir, "Destination folder")
	flag.StringVar(&conf.Package, "pkg", "", "Set package name to generate, default value is metadata.name of the source IDL.")
	flag.BoolVar(&conf.Debug, "debug", false, "debug mode")
	flag.BoolVar(&conf.RemoveAccountSuffix, "remove-account-suffix", false, "Remove \"Account\" suffix from accessors (if leads to duplication, e.g. \"SetFooAccountAccount\")")

	flag.StringVar((*string)(&conf.Encoding), "codec", string(EncodingBorsh), "Choose codec")
	flag.StringVar((*string)(&conf.TypeID), "type-id", string(TypeIDAnchor), "Choose typeID kind")
	flag.StringVar(&conf.ModPath, "mod", "", "Generate a go.mod file with the necessary dependencies, and this module")
	flag.Parse()

	if err := conf.Validate(); err != nil {
		panic(fmt.Errorf("error while validating config: %w", err))
	}

	var ts time.Time
	if GetConfig().Debug {
		ts = time.Unix(0, 0)
	} else {
		ts = time.Now()
	}
	if len(filenames) == 0 {
		Sfln(
			"[%s] No IDL files provided",
			Red(XMark),
		)
		os.Exit(1)
	}
	{
		exists, err := DirExists(GetConfig().DstDir)
		if err != nil {
			panic(err)
		}
		if !exists {
			MustCreateFolderIfNotExists(GetConfig().DstDir, os.ModePerm)
		}
	}

	callbacks := make([]func(), 0)
	defer func() {
		for _, cb := range callbacks {
			cb()
		}
	}()

	for _, idlFilepath := range filenames {
		Sfln(
			"[%s] Generating client from IDL: %s",
			Shakespeare("+"),
			Shakespeare(idlFilepath),
		)
		idlFile, err := os.Open(idlFilepath)
		if err != nil {
			panic(err)
		}

		dec := json.NewDecoder(idlFile)

		var idl IDL

		err = dec.Decode(&idl)
		if err != nil {
			panic(err)
		}
		{
			if idl.State != nil {
				Sfln(
					"%s idl.State is defined, but generator is not implemented yet.",
					OrangeBG("[?]"),
				)
			}
		}

		// spew.Dump(idl)

		// Create subfolder for package for generated assets:
		packageAssetFolderName := sighash.ToRustSnakeCase(idl.Metadata.Name)
		var dstDirForFiles string
		if GetConfig().Debug {
			packageAssetFolderPath := path.Join(GetConfig().DstDir, packageAssetFolderName)
			MustCreateFolderIfNotExists(packageAssetFolderPath, os.ModePerm)
			// Create folder for assets generated during this run:
			thisRunAssetFolderName := ToLowerCamel(idl.Metadata.Name) + "_" + ts.Format(FilenameTimeFormat)
			thisRunAssetFolderPath := path.Join(packageAssetFolderPath, thisRunAssetFolderName)
			// Create a new assets folder inside the main assets folder:
			MustCreateFolderIfNotExists(thisRunAssetFolderPath, os.ModePerm)
			dstDirForFiles = thisRunAssetFolderPath
		} else {
			if GetConfig().DstDir == generatedDir {
				dstDirForFiles = filepath.Join(GetConfig().DstDir, packageAssetFolderName)
			} else {
				dstDirForFiles = GetConfig().DstDir
			}
		}
		MustCreateFolderIfNotExists(dstDirForFiles, os.ModePerm)

		files, err := GenerateClientFromProgramIDL(idl)
		if err != nil {
			panic(err)
		}

		{
			mdf := &modfile.File{}
			mdf.AddModuleStmt(GetConfig().ModPath)

			mdf.AddNewRequire("github.com/gagliardetto/solana-go", "v1.5.0", false)
			mdf.AddNewRequire("github.com/fragmetric-labs/solana-binary-go", "v0.8.0", false)
			mdf.AddNewRequire("github.com/gagliardetto/treeout", "v0.1.4", false)
			mdf.AddNewRequire("github.com/gagliardetto/gofuzz", "v1.2.2", false)
			mdf.AddNewRequire("github.com/stretchr/testify", "v1.6.1", false)
			mdf.AddNewRequire("github.com/davecgh/go-spew", "v1.1.1", false)
			mdf.Cleanup()

			//callbacks = append(callbacks, func() {
			//	Ln()
			//	Ln(Bold("Don't forget to import the necessary dependencies!"))
			//	Ln()
			//	for _, v := range mdf.Require {
			//		Sfln(
			//			"	go get %s@%s",
			//			v.Mod.Path,
			//			v.Mod.Version,
			//		)
			//	}
			//	Ln()
			//})

			if GetConfig().ModPath != "" {
				mfBytes, err := mdf.Format()
				if err != nil {
					panic(err)
				}
				gomodFilepath := filepath.Join(dstDirForFiles, "go.mod")
				Sfln(
					"[%s] %s",
					Lime(Checkmark),
					MustAbs(gomodFilepath),
				)
				// Write `go.mod` file:
				err = os.WriteFile(gomodFilepath, mfBytes, 0666)
				if err != nil {
					panic(err)
				}
			}
		}

		for _, file := range files {
			// err := file.Render(os.Stdout)
			// if err != nil {
			// 	panic(err)
			// }
			file.File.HeaderComment("Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.")
			{
				// Save assets:
				assetFileName := file.Name + ".go"
				assetFilepath := path.Join(dstDirForFiles, assetFileName)

				// Create file:
				goFile, err := os.Create(assetFilepath)
				if err != nil {
					panic(err)
				}
				defer goFile.Close()

				// Write generated code file:
				Sfln(
					"[%s] %s",
					Lime(Checkmark),
					MustAbs(assetFilepath),
				)
				err = file.File.Render(goFile)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func FormatSighash(buf []byte) string {
	elems := make([]string, 0)
	for _, v := range buf {
		elems = append(elems, strconv.Itoa(int(v)))
	}

	return "[" + strings.Join(elems, ", ") + "]"
}

func GenerateClientFromProgramIDL(idl IDL) ([]*FileWrapper, error) {
	if idl.Address == "" {
		idl.Address = idl.Metadata.Address
	}

	if GetConfig().Package != "" {
		idl.Metadata.Name = GetConfig().Package
	}

	if err := idl.Validate(); err != nil {
		return nil, err
	}

	// configurable address map
	addresses := make(map[string]string)

	// Create and populate Go file that holds all the basic
	// elements of an instruction client:
	files := make([]*FileWrapper, 0)

	// register complex enums:
	{
		for _, typ := range idl.Types {
			registerComplexEnums(&idl, typ)
		}
		for _, typ := range idl.Accounts {
			registerComplexEnums(&idl, typ)
		}
	}
	//遍历指令的名称，首字母大写驼峰命名格式
	typesMap := make(map[string]string)
	for _, instruction := range idl.Instructions {
		typesMap[ToCamel(instruction.Name)] = instruction.Name
	}
	defs := make(map[string]IdlTypeDef)
	{
		file := NewGoFile(idl.Metadata.Name, true)
		// Declare types from IDL:
		for _, typ := range idl.Types {
			if _, isExist := typesMap[ToCamel(typ.Name)]; isExist {
				continue
			}
			defs[typ.Name] = typ
			file.Add(genTypeDef(&idl, nil, IdlTypeDef{
				Name: ToCamel(typ.Name),
				Type: typ.Type,
			}))
		}
		files = append(files, &FileWrapper{
			Name: "types",
			File: file,
		})
	}

	{
		file := NewGoFile(idl.Metadata.Name, true)
		// Declare account layouts from IDL:
		for _, acc := range idl.Accounts {
			if _, ok := defs[acc.Name]; ok {
				file.Add(genTypeDef(&idl, acc.Discriminator, IdlTypeDef{
					Name: ToCamel(defs[acc.Name].Name) + "Account",
					Type: defs[acc.Name].Type,
				}))
			} else {
				panic(`not implemented - only IDL from ("anchor": ">=0.30.0") is available`)
			}
		}
		files = append(files, &FileWrapper{
			Name: "accounts",
			File: file,
		})
	}

	constantsFile := NewGoFile(idl.Metadata.Name, false)
	constantsCode := Empty()
	constantsCodeMap := make(map[string][]*Statement)
	// Instructions:
	for _, instruction := range idl.Instructions {
		file := NewGoFile(idl.Metadata.Name, true)
		insExportedName := ToCamel(instruction.Name)
		var args []IdlField
		for _, arg := range instruction.Args {
			idlFieldArg := IdlField{
				Name: arg.Name,
				Docs: arg.Docs,
				Type: IdlType{
					asString:         arg.Type.asString,
					asIdlTypeVec:     arg.Type.asIdlTypeVec,
					asIdlTypeOption:  arg.Type.asIdlTypeOption,
					asIdlTypeArray:   arg.Type.asIdlTypeArray,
					asIdlTypeDefined: nil,
				},
			}
			if arg.Type.asIdlTypeDefined != nil {
				idlFieldArg.Type.asIdlTypeDefined = &IdlTypeDefined{
					Defined: IdLTypeDefinedName{
						Name: arg.Type.asIdlTypeDefined.Defined.Name,
					},
				}
			}
			args = append(args, idlFieldArg)
		}

		{
			code := Empty().Line().Line()

			for _, doc := range instruction.Docs {
				code.Comment(doc).Line()
			}

			if len(instruction.Docs) == 0 {
				code.Commentf(
					"%s is the `%s` instruction.",
					insExportedName,
					instruction.Name,
				).Line()
			}

			code.Type().Id(insExportedName).StructFunc(func(fieldsGroup *Group) {
				for argIndex, arg := range args {
					if len(arg.Docs) > 0 {
						if argIndex > 0 {
							fieldsGroup.Line()
						}
						for _, doc := range arg.Docs {
							fieldsGroup.Comment(doc)
						}
					}
					fieldsGroup.Add(genField(arg, true)).
						Add(func() Code {
							if arg.Type.IsIdlTypeOption() {
								return Tag(map[string]string{
									"bin": "optional",
								})
							}
							return nil
						}())
				}

				fieldsGroup.Line()

				{
					lastGroupName := ""
					// Add comments of the accounts from rust docs.
					instruction.Accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {
						comment := &strings.Builder{}
						indent := 6
						var prepend int

						if groupPath != "" {
							thisGroupName := filepath.Base(groupPath)
							indent = len(thisGroupName) + 2
							if strings.Count(groupPath, "/") == 0 {
								prepend = 6
							} else {
								prepend = 6 + (strings.Count(groupPath, "/") * 2) + len(strings.TrimSuffix(groupPath, thisGroupName)) - 1
							}
							if lastGroupName != groupPath {
								comment.WriteString(strings.Repeat("·", prepend-1) + Sf(" %s: ", thisGroupName))
							} else {
								comment.WriteString(strings.Repeat("·", prepend+indent-1) + " ")
							}
							lastGroupName = groupPath
						}

						comment.WriteString(Sf("[%v] = ", accountIndex))
						comment.WriteString("[")
						if ia.Writable {
							comment.WriteString("WRITE")
						}
						if ia.Signer {
							if ia.Writable {
								comment.WriteString(", ")
							}
							comment.WriteString("SIGNER")
						}
						comment.WriteString("] ")
						comment.WriteString(ia.Name)

						fieldsGroup.Comment(comment.String())
						for _, doc := range ia.Docs {
							fieldsGroup.Comment(strings.Repeat("·", prepend+indent-1+6) + " " + doc)
						}
						if accountIndex < instruction.Accounts.NumAccounts()-1 {
							fieldsGroup.Comment("")
						}

						accountIndex++
						return true
					})
				}
				fieldsGroup.Qual(PkgSolanaGo, "AccountMetaSlice").Tag(map[string]string{
					"bin": "-",
				})
			})

			file.Add(code.Line())
		}

		//NewInstructionBuilder
		{
			builderFuncName := formatBuilderFuncName(insExportedName)
			code := Empty()
			code.Commentf(
				"%s creates a new `%s` instruction builder.",
				builderFuncName,
				insExportedName,
			).Line()
			//
			code.Func().Id(builderFuncName).Params().Op("*").Id(insExportedName).
				BlockFunc(func(body *Group) {
					body.Id("nd").Op(":=").Op("&").Id(insExportedName).Block(
						Id("AccountMetaSlice").Op(":").Make(Qual(PkgSolanaGo, "AccountMetaSlice"), Lit(instruction.Accounts.NumAccounts())).Op(","),
					)
					instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
						if account.Name == "program" {
							body.Id("nd").Dot("AccountMetaSlice").Index(Lit(index)).Op("=").Qual(PkgSolanaGo, "Meta").Call(Op("ProgramID"))
							return true
						}
						if account.Address == "" {
							if account.PDA != nil && account.PDA.Program == nil {
								isOnlyConst := true
								var constSeeds [][]byte
								for _, seed := range account.PDA.Seeds { //看这个种子的参数是否都是const,如果都是const,放入addresses,在instruction.go中直接定义变量
									if seed.Value == nil {
										isOnlyConst = false
										break
									}
									constSeeds = append(constSeeds, seed.Value)
								}
								if isOnlyConst { //这个PDA是一个唯一账户，没有使用其他交易账户
									address, _, _ := solana.FindProgramAddress(constSeeds, solana.MustPublicKeyFromBase58(idl.Address))
									addresses[ToCamel(account.Name)+"Pda"] = address.String() //小写的，大写开头的驼峰会和types冲突
									def := Qual(PkgSolanaGo, "Meta").Call(Op(ToCamel(account.Name) + "Pda"))
									if account.Writable {
										def.Dot("WRITE").Call()
									}
									if account.Signer {
										def.Dot("SIGNER").Call()
									}
									body.Id("nd").Dot("AccountMetaSlice").Index(Lit(index)).Op("=").Add(def)
								}
							}
							return true
						}
						programName := getDefProgram(ToCamel(account.Name), account.Address)
						addresses[programName] = account.Address
						// fmt.Println("global pda:", programName)
						def := Qual(PkgSolanaGo, "Meta").Call(Op(programName))
						if account.Writable {
							def.Dot("WRITE").Call()
						}
						if account.Signer {
							def.Dot("SIGNER").Call()
						}
						body.Id("nd").Dot("AccountMetaSlice").Index(Lit(index)).Op("=").Add(def)
						return true
					})
					body.Return(Id("nd"))
				})
			file.Add(code.Line())
		}

		// 创建Set函数,set账户列表前的参数
		{
			// Declare methods that set the parameters of the instruction:
			code := Empty()
			for _, arg := range args {
				exportedArgName := ToCamel(arg.Name)

				code.Line().Line()
				name := "Set" + exportedArgName
				code.Commentf("%s sets the %q parameter.", name, arg.Name).Line()
				for _, doc := range arg.Docs {
					code.Comment(doc).Line()
				}

				code.Func().Params(Id("inst").Op("*").Id(insExportedName)).Id(name).
					Params(
						ListFunc(func(params *Group) {
							// Parameters:
							params.Id(arg.Name).Add(genTypeName(arg.Type))
						}),
					).
					Params(
						ListFunc(func(results *Group) {
							// Results:
							results.Op("*").Id(insExportedName)
						}),
					).
					BlockFunc(func(body *Group) {
						// Body:
						body.Id("inst").Dot(exportedArgName).Op("=").
							Add(func() Code {
								if isComplexEnum(arg.Type) {
									return nil
								}
								return Op("&")
							}()).
							Id(arg.Name)

						body.Return().Id("inst")
					})
			}

			file.Add(code.Line())
		}

		// 创建Set/get函数,账户列表
		{
			// Declare methods that set/get accounts for the instruction:
			code := Empty()

			declaredReceivers := []string{}
			groupMemberIndex := 0
			instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
				builderStructName := insExportedName + ToCamel(parentGroupPath) + "AccountsBuilder"
				hasNestedParent := parentGroupPath != ""
				isDeclaredReceiver := SliceContains(declaredReceivers, parentGroupPath)

				if !hasNestedParent {
					groupMemberIndex = index
				}
				if hasNestedParent && !isDeclaredReceiver {
					groupMemberIndex = 0
					declaredReceivers = append(declaredReceivers, parentGroupPath)
					// many accounts (???)
					// builder struct for this accounts group:

					code.Line().Line()
					for _, doc := range parentGroup.Docs {
						code.Comment(doc).Line()
					}
					code.Type().Id(builderStructName).Struct(
						Qual(PkgSolanaGo, "AccountMetaSlice").Tag(map[string]string{
							"bin": "-",
						}),
					)

					// func that returns a new builder for this account group:
					code.Line().Line().Func().Id("New" + builderStructName).Params().Op("*").Id(builderStructName).
						BlockFunc(func(gr *Group) {
							gr.Return().Op("&").Id(builderStructName).Block(
								Id("AccountMetaSlice").Op(":").Make(Qual(PkgSolanaGo, "AccountMetaSlice"), Lit(parentGroup.Accounts.NumAccounts())).Op(","),
							)
						}).Line().Line()

					// Method on intruction builder that accepts the accounts group builder, and copies the accounts:
					code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("Set" + ToCamel(parentGroup.Name) + "AccountsFromBuilder").
						Params(
							ListFunc(func(st *Group) {
								// Parameters:
								st.Id(ToLowerCamel(builderStructName)).Op("*").Id(builderStructName)
							}),
						).
						Params(
							ListFunc(func(st *Group) {
								// Results:
								st.Op("*").Id(insExportedName)
							}),
						).
						BlockFunc(func(gr *Group) {
							// Body:

							tpIndex := index
							// spew.Dump(parentGroup)
							for _, subAccount := range parentGroup.Accounts {
								if subAccount.IdlAccount != nil {
									exportedAccountName := ToCamel(subAccount.IdlAccount.Name)

									def := Id("inst").Dot("AccountMetaSlice").Index(Lit(tpIndex)).
										Op("=").Id(ToLowerCamel(builderStructName)).Dot(formatAccountAccessorName("Get", exportedAccountName)).Call()

									gr.Add(def)
								}
								tpIndex++
							}

							gr.Return().Id("inst")
						})
				}

				{
					exportedAccountName := ToCamel(account.Name)
					lowerAccountName := ToLowerCamel(account.Name)

					var receiverTypeName string
					if parentGroupPath == "" {
						receiverTypeName = insExportedName
					} else {
						receiverTypeName = builderStructName
					}

					code.Add(genAccountGettersSetters(
						receiverTypeName,
						account,
						groupMemberIndex,
						exportedAccountName,
						lowerAccountName,
						instruction.Accounts,
						// addresses,
						instruction.Args,
						// constantsCode,
						constantsCodeMap,
						&idl,
					))
					groupMemberIndex++
				}
				return true
			})

			file.Add(code.Line())
		}

		// 	code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("AddRemainingAccounts").Params(
		// 		Id("remainingAccounts").Index().Qual(PkgSolanaGo, "PublicKey"),
		// 	).Params(Op("*").Id(insExportedName)).
		// 		Block(
		// 			jen.Id("accounts").Op(":=").Lit(instruction.Accounts.NumAccounts()),
		// 			jen.For(
		// 				jen.Id("i").Op(":=").Range().Id("remainingAccounts"),
		// 			).Block(
		// 				jen.Id("index").Op(":=").Id("accounts").Op("+").Id("i"),
		// 				jen.Id("inst").Dot("AccountMetaSlice").Index(jen.Id("index")).Op("=").Qual(PkgSolanaGo, "Meta").Call(jen.Id("remainingAccounts").Index(jen.Id("i"))).Dot("WRITE").Call(),
		// 			),
		// 			Return(Id("inst")),
		// 		)
		// 	file.Add(code.Line())
		// }

		// SetAccounts 设置账户列表
		{
			code := Empty()
			code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("SetAccounts").
				Params(
					ListFunc(func(params *Group) {
						// 参数
						params.Id("accounts").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// 返回结果
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Id("inst").Dot("AccountMetaSlice").Op("=").Id("accounts")
					body.Return(Id("inst").Dot("Validate").Call())
				})
			file.Add(code.Line())
		}

		// SetRemainingAccounts 设置剩余账户
		{
			code := Empty()
			code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("SetRemainingAccounts").
				Params(
					ListFunc(func(params *Group) {
						// 参数
						params.Id("metas").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// 返回结果
						results.Op("*").Id(insExportedName)
					}),
				).
				BlockFunc(func(body *Group) {
					body.Id("inst").Dot("AccountMetaSlice").Op("=").Append(Id("inst").Dot("AccountMetaSlice").Index(Lit(0).Op(":").Lit(len(instruction.Accounts))), Id("metas").Op("..."))
					body.Return(Id("inst"))
				})
			file.Add(code.Line())
		}

		// GetRemainingAccounts 获取剩余账户
		{
			code := Empty()
			code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("GetRemainingAccounts").
				Params(
					ListFunc(func(params *Group) {
						// 无参数
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// 返回结果
						results.Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Return(Id("inst").Dot("AccountMetaSlice").Index(Lit(instruction.Accounts.NumAccounts()).Op(":")))
				})
			file.Add(code.Line())
		}

		//Build() 构建
		{
			// Declare `Build` method on instruction:
			code := Empty()

			code.Line().Line().Func().Params(Id("inst").Id(insExportedName)).Id("Build").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id("Instruction")
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					var typeIDCode Code

					GetConfig().TypeID.
						On(
							TypeIDNameSlice{
								TypeIDUvarint32,
							},
							func() {
								typeIDCode = Qual(PkgDfuseBinary, "TypeIDFromUvarint32").Call(Id("Instruction_" + insExportedName))
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDUint32,
							},
							func() {
								typeIDCode = Qual(PkgDfuseBinary, "TypeIDFromUint32").Call(Id("Instruction_"+insExportedName), Qual("encoding/binary", "LittleEndian"))
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDUint8,
							},
							func() {
								typeIDCode = Qual(PkgDfuseBinary, "TypeIDFromUint8").Call(Id("Instruction_" + insExportedName))
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDAnchor,
							},
							func() {
								typeIDCode = Id("Instruction_" + insExportedName)
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDNoType,
							},
							func() {
								// TODO
							},
						)

					body.Return().Op("&").Id("Instruction").Values(
						Dict{
							Id("BaseVariant"): Qual(PkgDfuseBinary, "BaseVariant").Values(
								Dict{
									Id("TypeID"): typeIDCode,
									Id("Impl"):   Id("inst"),
								},
							),
						},
					)
				})
			file.Add(code.Line())
		}

		//ValidateAndBuild() 验证并构建
		{
			// Declare `ValidateAndBuild` method on instruction:
			code := Empty()

			code.Line().Line().
				Comment("ValidateAndBuild validates the instruction parameters and accounts;").
				Line().
				Comment("if there is a validation error, it returns the error.").
				Line().
				Comment("Otherwise, it builds and returns the instruction.").
				Line().
				Func().Params(Id("inst").Id(insExportedName)).Id("ValidateAndBuild").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id("Instruction")
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.If(
						Err().Op(":=").Id("inst").Dot("Validate").Call(),
						Err().Op("!=").Nil(),
					).Block(
						Return(Nil(), Err()),
					)

					body.Return(Id("inst").Dot("Build").Call(), Nil())
				})
			file.Add(code.Line())
		}

		//Validate() 验证
		{
			// Declare `Validate` method on instruction:
			code := Empty()

			code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("Validate").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					if len(args) > 0 {
						body.Comment("Check whether all (required) parameters are set:")

						body.BlockFunc(func(paramVerifyBody *Group) {
							for _, arg := range args {
								exportedArgName := ToCamel(arg.Name)

								// Optional params can be empty.
								if arg.Type.IsIdlTypeOption() {
									continue
								}

								paramVerifyBody.If(Id("inst").Dot(exportedArgName).Op("==").Nil()).Block(
									Return(
										Qual("errors", "New").Call(Lit(Sf("%s parameter is not set", exportedArgName))),
									),
								)
							}
						})
						body.Line()
					}

					body.Comment("Check whether all (required) accounts are set:")
					body.BlockFunc(func(accountValidationBlock *Group) {
						instruction.Accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {
							exportedAccountName := ToCamel(filepath.Join(groupPath, ia.Name))

							if ia.Optional {
								accountValidationBlock.Line().Commentf(
									"[%v] = %s is optional",
									accountIndex,
									exportedAccountName,
								).Line()
							} else {
								accountValidationBlock.If(Id("inst").Dot("AccountMetaSlice").Index(Lit(accountIndex)).Op("==").Nil()).Block(
									Return(Qual("errors", "New").Call(Lit(Sf("accounts.%s is not set", exportedAccountName)))),
								)
							}

							return true
						})
					})

					body.Return(Nil())
				})
			file.Add(code.Line())
		}

		//EncodeToTree() 编码到树
		{
			// Declare `EncodeToTree(parent treeout.Branches)` method in instruction:
			code := Empty()

			code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("EncodeToTree").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("parent").Qual(PkgTreeout, "Branches")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					body.Id("parent").Dot("Child").Call(Qual(PkgFormat, "Program").Call(Id("ProgramName"), Id("ProgramID"))).Op(".").
						Line().Comment("").Line().
						Id("ParentFunc").Parens(Func().Parens(Id("programBranch").Qual(PkgTreeout, "Branches")).BlockFunc(
						func(programBranchGroup *Group) {
							programBranchGroup.Id("programBranch").Dot("Child").Call(Qual(PkgFormat, "Instruction").Call(Lit(insExportedName))).Op(".").
								Line().Comment("").Line().
								Id("ParentFunc").Parens(Func().Parens(Id("instructionBranch").Qual(PkgTreeout, "Branches")).BlockFunc(
								func(instructionBranchGroup *Group) {

									instructionBranchGroup.Line().Comment("Parameters of the instruction:")

									instructionBranchGroup.Id("instructionBranch").Dot("Child").Call(Lit(Sf("Params[len=%v]", len(args)))).Dot("ParentFunc").Parens(Func().Parens(Id("paramsBranch").Qual(PkgTreeout, "Branches")).BlockFunc(func(paramsBranchGroup *Group) {
										longest := treeFindLongestNameFromFields(args)
										for _, arg := range args {
											exportedArgName := ToCamel(arg.Name)
											paramsBranchGroup.Id("paramsBranch").Dot("Child").
												Call(
													Qual(PkgFormat, "Param").Call(
														Lit(strings.Repeat(" ", longest-len(exportedArgName))+exportedArgName+StringIf(arg.Type.IsIdlTypeOption(), " (OPT)")),
														Add(CodeIf(!arg.Type.IsIdlTypeOption() && !isComplexEnum(arg.Type), Op("*"))).Id("inst").Dot(exportedArgName),
													),
												)
										}
									}))

									instructionBranchGroup.Line().Comment("Accounts of the instruction:")

									instructionBranchGroup.Id("instructionBranch").Dot("Child").Call(Lit(Sf("Accounts[len=%v]", instruction.Accounts.NumAccounts()))).Dot("ParentFunc").Parens(
										Func().Parens(Id("accountsBranch").Qual(PkgTreeout, "Branches")).BlockFunc(func(accountsBranchGroup *Group) {

											longest := treeFindLongestNameFromAccounts(instruction.Accounts)
											instruction.Accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {

												cleanedName := treeFormatAccountName(ia.Name)

												exportedAccountName := filepath.Join(groupPath, cleanedName)

												access := Id("accountsBranch").Dot("Child").Call(Qual(PkgFormat, "Meta").Call(Lit(strings.Repeat(" ", longest-len(exportedAccountName))+exportedAccountName), Id("inst").Dot("AccountMetaSlice").Dot("Get").Call(Lit(accountIndex))))
												accountsBranchGroup.Add(access)
												return true
											})
										}))
								}))
						}))
				})
			file.Add(code.Line())
		}

		//MarshalWithEncoder() 使用编码器编码
		{
			// Declare `MarshalWithEncoder(encoder *bin.Encoder) error` method on instruction:
			file.Add(
				genMarshalWithEncoder_struct(
					&idl,
					false,
					insExportedName,
					"",
					args,
					true,
				),
			)
		}

		//UnmarshalWithDecoder() 使用解码器解码
		{
			// Declare `UnmarshalWithDecoder(decoder *bin.Decoder) error` method on instruction:
			file.Add(
				genUnmarshalWithDecoder_struct(
					&idl,
					false,
					insExportedName,
					"",
					args,
					bin.TypeID{},
				))
		}

		// Declare instruction initializer func:
		{
			paramNames := []string{}
			for _, arg := range args {
				paramNames = append(paramNames, arg.Name)
			}
			code := Empty()
			name := "New" + insExportedName + "Instruction"
			code.Commentf("%s declares a new %s instruction with the provided parameters and accounts.", name, insExportedName)
			code.Line()
			code.Func().Id(name).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						{
							for argIndex, arg := range args {
								paramNames = append(paramNames, arg.Name)
								params.Add(func() Code {
									if argIndex == 0 {
										return Line().Comment("Parameters:")
									}
									return Empty()
								}()).Line().Id(arg.Name).Add(genTypeName(arg.Type))
							}
						}
						{
							instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
								// skip sysvars:
								if isSysVar(account.Name) || account.Name == "program" {
									return true
								}
								if account.Address != "" {
									if getDefProgram(ToLowerCamel(account.Name), account.Address) != "" {
										//在instructionBuilder中已经设置的账户
										return true
									}
								}
								// if account.PDA != nil && account.PDA.Program == nil {
								// 	isOnlyConst := true
								// 	for _, seed := range account.PDA.Seeds {
								// 		if seed.Kind == "arg" { //Kind == "arg"
								// 			continue
								// 		}
								// 		if seed.Value == nil { //Kind == "account"
								// 			isOnlyConst = false
								// 			break
								// 		}
								// 	}
								// 	if !isOnlyConst {
								// 		return true
								// 	}
								// }
								var accountName string
								if parentGroupPath == "" {
									accountName = ToLowerCamel(account.Name)
								} else {
									accountName = ToLowerCamel(parentGroupPath + "/" + ToLowerCamel(account.Name))
								}

								if SliceContains(paramNames, accountName) {
									accountName = accountName + "Account"
								}

								params.Add(func() Code {
									if index == 0 {
										return Line().Comment("Accounts:").Line()
									}
									return Line()
								}()).Id(accountName).Qual(PkgSolanaGo, "PublicKey")
								return true
							})
						}
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id(insExportedName)
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					builder := body.Return().Id(formatBuilderFuncName(insExportedName)).Call()
					{
						for _, arg := range args {
							exportedArgName := ToCamel(arg.Name)
							builder.Op(".").Line().Id("Set" + exportedArgName).Call(Id(arg.Name))
						}
					}

					{
						declaredReceivers := []string{}
						instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
							// skip sysvars:
							if isSysVar(account.Name) || account.Name == "program" {
								return true
							}
							if account.Address != "" {
								if getDefProgram(ToLowerCamel(account.Name), account.Address) != "" {
									//在instructionBuilder中已经设置的账户
									return true
								}
							}
							// if account.PDA != nil && account.PDA.Program == nil {
							// 	isOnlyConst := true
							// 	for _, seed := range account.PDA.Seeds {
							// 		if seed.Kind == "arg" { //Kind == "arg"
							// 			continue
							// 		}
							// 		if seed.Value == nil { //Kind == "account"
							// 			isOnlyConst = false
							// 			break
							// 		}
							// 	}
							// 	if !isOnlyConst {
							// 		return true
							// 	}
							// }
							var accountName string
							if parentGroupPath == "" {
								accountName = ToLowerCamel(account.Name)
							} else {
								// TODO
							}

							if SliceContains(paramNames, accountName) {
								accountName = accountName + "Account"
							}

							builderStructName := insExportedName + ToCamel(parentGroupPath) + "AccountsBuilder"
							hasNestedParent := parentGroupPath != ""
							isDeclaredReceiver := SliceContains(declaredReceivers, parentGroupPath)

							if hasNestedParent && !isDeclaredReceiver {
								declaredReceivers = append(declaredReceivers, parentGroupPath)
								builder.Op(".").Line().Id("Set" + ToCamel(parentGroup.Name) + "AccountsFromBuilder").Call(
									Line().Id("New" + builderStructName).Call().
										Add(
											DoGroup(func(gr *Group) {
												// Body:
												for subIndex, subAccount := range parentGroup.Accounts {
													if subAccount.IdlAccount != nil {
														exportedAccountName := ToCamel(subAccount.IdlAccount.Name)
														accountName = ToLowerCamel(parentGroupPath + "/" + ToLowerCamel(exportedAccountName))

														gr.Op(".").Add(func() Code {
															if subIndex == 0 {
																return Line().Line()
															}
															return Line()
														}()).Id(formatAccountAccessorName("Set", exportedAccountName)).Call(Id(accountName))

														if subIndex == len(parentGroup.Accounts)-1 {
															gr.Op(",").Line()
														}
													}
												}
											}),
										),
								)
							}

							if !hasNestedParent {
								builder.Op(".").Line().Id(formatAccountAccessorName("Set", ToCamel(account.Name))).Call(Id(accountName))
							}

							return true
						})
					}

					// builder.Op(".").Line().Id("Build").Call()
				})

			file.Add(code.Line())
		}

		// 新增初始化函数
		{
			createdAccounts := make(map[string]string) //记录已经创建的账户
			paramNames := []string{}
			for _, arg := range args {
				paramNames = append(paramNames, arg.Name)
			}
			code := Empty()
			name := "NewSimple" + insExportedName + "Instruction"
			code.Commentf("%s declares a new %s instruction with the provided parameters and accounts.", name, insExportedName)
			code.Line()
			code.Func().Id(name).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						{
							for argIndex, arg := range args {
								paramNames = append(paramNames, arg.Name)
								params.Add(func() Code {
									if argIndex == 0 {
										return Line().Comment("Parameters:")
									}
									return Empty()
								}()).Line().Id(arg.Name).Add(genTypeName(arg.Type))
							}
						}
						//pda的有指定program的账户直接采用参数传递，没用指定program的账户，看种子是不是包含的有账户
						{
							instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
								// if account.Address != "" || account.PDA != nil {
								// 	return true
								// }
								if account.Name == "program" {
									return true
								}
								accountLower := ToLowerCamel(account.Name)
								if account.Address != "" {
									if getDefProgram(accountLower, account.Address) != "" {
										//在instructionBuilder中已经设置的账户
										return true
									}
								}
								if account.PDA != nil {
									if account.PDA.Program != nil {
										params.Add(Line().Id(accountLower).Qual(PkgSolanaGo, "PublicKey"))
										createdAccounts[account.Name] = account.Name
									} else {
										seeds := account.PDA.Seeds
										// if account.PDA.Program == nil { //pda的program没有
										isOnlyConst := true
										// var constSeeds [][]byte
										for _, seed := range seeds { //看这个种子的参数是否都是const,如果都是const,放入addresses,在instruction.go中直接定义变量
											if seed.Value == nil {
												isOnlyConst = false
												break
											}
											// constSeeds = append(constSeeds, seed.Value)
										}
										if isOnlyConst { //这个PDA是一个唯一账户，没有使用其他交易账户
											// address, _, _ := solana.FindProgramAddress(constSeeds, solana.MustPublicKeyFromBase58(idl.Address))
											// addresses[accountLower] = address.String() //小写的，大写开头的驼峰会和types冲突
											createdAccounts[accountLower] = account.Name
											return true
										}
										//这个种子的参数有交易账户，下面就直接调用constant.go的MustFind函数
									}
									return true
								}

								var accountName string
								if parentGroupPath == "" {
									accountName = accountLower
								} else {
									accountName = ToLowerCamel(parentGroupPath + "/" + accountLower)
								}

								if SliceContains(paramNames, accountName) {
									accountName = accountName + "Account"
								}
								if _, isExist := createdAccounts[accountLower]; isExist {
									return true
								}
								params.Add(func() Code {
									if index == 0 {
										return Line().Comment("Accounts:").Line()
									}
									return Line()
								}()).Id(accountName).Qual(PkgSolanaGo, "PublicKey")
								createdAccounts[accountLower] = accountLower
								return true
							})
						}
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id(insExportedName)
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					{
						instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
							if account.Address != "" {
								if getDefProgram(ToCamel(account.Name), account.Address) != "" {
									//InstructionBuilder里面已经set过了
									return true
								}
								body.Id(ToLowerCamel(account.Name)).Op(":=").Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(account.Address))
								createdAccounts[ToLowerCamel(account.Name)] = account.Address
							}
							return true
						})
					}
					for i := 0; i < len(instruction.Accounts)-len(createdAccounts); i++ {
						instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
							if _, isExist := createdAccounts[ToLowerCamel(account.Name)]; isExist {
								return true
							}
							if account.PDA != nil {
								if account.PDA.Program != nil { //说明pda的program是其他程序，已经作为参数传进来
									return true
								}
								seeds := account.PDA.Seeds
								params := make([]jen.Code, len(seeds))
								for _, seed := range seeds {
									if seed.Value == nil { //Kind == "account"
										if seed.Kind == "arg" {
											params = append(params, Id(ToLowerCamel(seed.Path)))
											continue
										}
										if _, isExist := createdAccounts[ToLowerCamel(seed.Path)]; !isExist {
											return true
										}
										// 如果种子引用是账户
										params = append(params, Id(ToLowerCamel(seed.Path)))
									}
								}
								if _, isExist := addresses[ToLowerCamel(account.Name)]; !isExist {
									address := getProgramAccount(account.Name)
									if address.Equals(solana.PublicKey{}) {
										exportedAccountName := ToCamel(account.Name)
										accessorName := strings.TrimSuffix(formatAccountAccessorName("MustFind", exportedAccountName), "Account") + "Address"
										body.Id(ToLowerCamel(account.Name)).Op(":=").Id(accessorName).Call(params...)
									} else {
										//solana.MustPublicKeyFromBase58(address)
										body.Id(ToLowerCamel(account.Name)).Op(":=").Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(address))
									}
								}
								createdAccounts[ToLowerCamel(account.Name)] = account.Name
							}
							return true
						})
						if len(createdAccounts) == len(instruction.Accounts) {
							break
						}
					}

					builder := body.Return().Id(formatBuilderFuncName(insExportedName)).Call()
					{
						for _, arg := range args {
							exportedArgName := ToCamel(arg.Name)
							builder.Op(".").Line().Id("Set" + exportedArgName).Call(Id(arg.Name))
						}
					}

					{
						declaredReceivers := []string{}

						instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
							// if account.Address != "" || account.PDA != nil {
							// 	return true
							// }
							if account.Name == "program" {
								return true
							}
							if account.Address != "" {
								if getDefProgram(ToCamel(account.Name), account.Address) != "" {
									//InstructionBuilder里面已经set过了
									return true
								}
							}
							if account.PDA != nil && account.PDA.Program == nil {
								isOnlyConst := true
								for _, seed := range account.PDA.Seeds {
									if seed.Value == nil { //Kind == "account"
										isOnlyConst = false
										break
									}
								}
								if isOnlyConst {
									return true
								}
							}
							var accountName string
							if parentGroupPath == "" {
								accountName = ToLowerCamel(account.Name)
							}

							if SliceContains(paramNames, accountName) {
								accountName = accountName + "Account"
							}

							builderStructName := insExportedName + ToCamel(parentGroupPath) + "AccountsBuilder"
							hasNestedParent := parentGroupPath != ""
							isDeclaredReceiver := SliceContains(declaredReceivers, parentGroupPath)

							if hasNestedParent && !isDeclaredReceiver {
								declaredReceivers = append(declaredReceivers, parentGroupPath)
								builder.Op(".").Line().Id("Set" + ToCamel(parentGroup.Name) + "AccountsFromBuilder").Call(
									Line().Id("New" + builderStructName).Call().
										Add(
											DoGroup(func(gr *Group) {
												// Body:
												for subIndex, subAccount := range parentGroup.Accounts {
													if subAccount.IdlAccount != nil {
														exportedAccountName := ToCamel(subAccount.IdlAccount.Name)
														accountName = ToLowerCamel(parentGroupPath + "/" + ToLowerCamel(exportedAccountName))

														gr.Op(".").Add(func() Code {
															if subIndex == 0 {
																return Line().Line()
															}
															return Line()
														}()).Id(formatAccountAccessorName("Set", exportedAccountName)).Call(Id(accountName))

														if subIndex == len(parentGroup.Accounts)-1 {
															gr.Op(",").Line()
														}
													}
												}
											}),
										),
								)
							}
							if !hasNestedParent {
								builder.Op(".").Line().Id(formatAccountAccessorName("Set", ToCamel(account.Name))).Call(Id(accountName))
							}

							return true
						})
					}

					// builder.Op(".").Line().Id("Build").Call()
				})

			file.Add(code.Line())
		}

		files = append(files, &FileWrapper{
			Name: strings.ToLower(insExportedName),
			File: file,
		})
	}
	for _, codes := range constantsCodeMap {
		for _, code := range codes {
			constantsCode.Func().Add(code)
			constantsCode.Line().Line()
		}
	}
	constantsFile.Add(constantsCode)

	files = append(files, &FileWrapper{
		Name: "constants",
		File: constantsFile,
	})

	// 生成instructions.go
	{

		file, err := genProgramBoilerplate(idl, addresses)
		if err != nil {
			return nil, err
		}
		files = append(files, &FileWrapper{
			Name: "instructions",
			File: file,
		})
	}

	{
		testFiles, err := genTestingFuncs(idl)
		if err != nil {
			return nil, err
		}
		files = append(files, testFiles...)
	}

	return files, nil
}

func genAccountGettersSetters(
	receiverTypeName string,
	account *IdlAccount,
	index int,
	exportedAccountName string,
	lowerAccountName string,
	accounts []IdlAccountItem,
	// addresses map[string]string,
	args []IdlField,
	// constantsCode *Statement,
	constantsCodeMap map[string][]*Statement,
	idl *IDL,
) Code {
	code := Empty()

	{
		code.Line().Line()
		name := formatAccountAccessorName("Set", exportedAccountName)
		code.Commentf("%s sets the %q account.", name, account.Name).Line()
		for _, doc := range account.Docs {
			code.Comment(doc).Line()
		}

		// Create account setters:
		code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id(name).
			Params(
				ListFunc(func(params *Group) {
					// Parameters:
					params.Id(lowerAccountName).Qual(PkgSolanaGo, "PublicKey")
				}),
			).
			Params(
				ListFunc(func(results *Group) {
					// Results:
					results.Op("*").Id(receiverTypeName)
				}),
			).
			BlockFunc(func(body *Group) {
				// Body:
				def := Id("inst").Dot("AccountMetaSlice").Index(Lit(index)).
					Op("=").Qual(PkgSolanaGo, "Meta").Call(Id(lowerAccountName))
				if account.Writable {
					def.Dot("WRITE").Call()
				}
				if account.Signer {
					def.Dot("SIGNER").Call()
				}
				body.Add(def)

				body.Return().Id("inst")
			})
	}

	{
		if account.PDA != nil {
			code.Line().Line()
			accessorName := strings.TrimSuffix(formatAccountAccessorName("Find", exportedAccountName), "Account") + "Address"
			seedValues := make([][]byte, len(account.PDA.Seeds))
			seedRefs := make([]string, len(account.PDA.Seeds))
			seedTypes := make([]IdlType, len(account.PDA.Seeds))
			var seedProgramValue *[]byte
			var seedProgramPath string //暂时没用，因为constants.go里面的PDA都是程序的唯一PDA
			if account.PDA.Program != nil {
				if account.PDA.Program.Kind == "account" {
					seedProgramPath = account.PDA.Program.Path
				} else if account.PDA.Program.Value == nil {
					panic("cannot handle non-const type program value in PDA seeds" + account.Address)
				}
				seedProgramValue = &account.PDA.Program.Value
			}
		OUTER:
			for i, seedDef := range account.PDA.Seeds {
				if seedDef.Value != nil { // type: const
					seedValues[i] = seedDef.Value
				} else {
					// First check if it's an account reference
					for _, acc := range accounts {
						if acc.IdlAccount.Name == seedDef.Path {
							seedRefs[i] = ToLowerCamel(acc.IdlAccount.Name)
							continue OUTER
						}
					}

					for _, argv := range args {
						argvName := strings.TrimPrefix(argv.Name, "_")
						if argvName == seedDef.Path {

							seedRefs[i] = ToLowerCamel(argv.Name)
							seedTypes[i] = argv.Type
							continue OUTER
						}
					}

					// Then check if it's an argument field reference
					parts := strings.Split(seedDef.Path, ".")
					if len(parts) == 2 {
						// Find the argument type
						var argType IdlType
						for _, arg := range args {
							if arg.Name == parts[0] {
								// Found the argument, now need to find the field type
								if arg.Type.IsIdlTypeDefined() {
									// Look up the defined type
									definedType := idl.Types.GetByName(arg.Type.GetIdlTypeDefined().Defined.Name)
									if definedType != nil && definedType.Type.Fields != nil {
										// Find the field
										for _, field := range *definedType.Type.Fields {
											if field.Name == parts[1] {
												argType = field.Type
												break
											}
										}
									}
								}
								break
							}
						}
						for _, typ := range idl.Types {
							if typ.Name == seedDef.Account {
								for _, field := range *typ.Type.Fields {
									if field.Name == parts[1] {
										argType = field.Type
										break
									}
								}
							}
						}

						paramName := ToLowerCamel(parts[0] + "_" + parts[1])

						seedTypes[i] = argType

						// Update the function signature to use the correct type
						seedRefs[i] = paramName
						continue OUTER
					}

					panic(fmt.Sprintf("cannot find related account or argument path %q", seedDef.Path))
				}
			}
			internalAccessorName := accessorName
			_, isExist := constantsCodeMap[internalAccessorName]
			if !isExist && account.PDA.Program == nil { //constants.go 只要程序ID为自己的程序ID的PDA,也就是没有program的PDA
				var kindIsAccount bool
				for _, seed := range account.PDA.Seeds {
					if seed.Value == nil { //如果kind是account/arg
						kindIsAccount = true
					}
				}
				if kindIsAccount { //如果种子里面有交易账户，才去生成对应的MustFind函数。否则，直接在instruction.go中定义变量
					constantsCodeMap[internalAccessorName] = append(constantsCodeMap[internalAccessorName],
						Id(internalAccessorName).
							Params(
								ListFunc(func(params *Group) {
									// Parameters:
									for i, seedRef := range seedRefs {
										if seedRef != "" {
											if seedTypes[i].IsArray() && seedTypes[i].GetArray().Elem.GetString() == "u8" {
												params.Id(seedRef).Index(Lit(32)).Byte()
												continue
											}
											switch seedTypes[i].asString {
											case IdlTypeBool, IdlTypeI8, IdlTypeI16, IdlTypeI32, IdlTypeI64, IdlTypeU8, IdlTypeU16, IdlTypeU32, IdlTypeU64, IdlTypeF32, IdlTypeF64, IdlTypeString:
												//是什么类型就是什么类型
												params.Id(seedRef).Add(genTypeName(seedTypes[i]))
											case IdlTypeI128, IdlTypeU128: //[16]byte
												params.Id(seedRef).Index(Lit(16)).Byte()
											default:
												params.Id(seedRef).Qual(PkgSolanaGo, "PublicKey")
											}

										}
									}
									if seedProgramPath != "" {
										params.Id(seedProgramPath).Qual(PkgSolanaGo, "PublicKey")
									}
								}),
							).
							Params(
								ListFunc(func(results *Group) {
									// Results:
									results.Id("pda").Qual(PkgSolanaGo, "PublicKey")
									results.Id("bumpSeed").Uint8()
									results.Id("err").Error()
								}),
							).
							BlockFunc(func(body *Group) {
								// Body:
								body.Add(Var().Id("seeds").Index().Index().Byte())

								for i, seedValue := range seedValues {
									if seedValue != nil {
										//body.Commentf("const: %s", string(seedValue))
										body.Commentf("const: 0x%s", hex.EncodeToString(seedValue))
										body.Add(Id("seeds").Op("=").Append(Id("seeds"), Index().Byte().ValuesFunc(func(group *Group) {
											for _, v := range seedValue {
												group.LitByte(v)
											}
										})))
									} else {
										seedRef := seedRefs[i]
										body.Commentf("path: %s", seedRef)
										if seedTypes[i].IsArray() && seedTypes[i].GetArray().Elem.GetString() == "u8" {
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id(seedRef).Index(Op(":")))) // Just pass the byte array directly
											continue
										}
										switch seedTypes[i].asString {
										case IdlTypeBool:
											body.Add(If(Id(seedRef).Block(
												Id("argBytes").Op(":=").Index().Byte().Values(Lit(1)),
											).Else().Block(
												Id("argBytes").Op(":=").Index().Byte().Values(Lit(0)),
											)))
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id("argBytes")))
										case IdlTypeI8, IdlTypeU8:
											body.Add(Id("argBytes").Op(":=").Index().Byte().Values(Byte().Call(Id(seedRef)))) //[]byte{byte(arg)}
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id("argBytes")))
										case IdlTypeI16, IdlTypeU16:
											body.Add(Id("argBytes").Op(":=").Make(Index().Byte(), Lit(2)))
											body.Add(Qual(PkgBinary, "LittleEndian").Dot("PutUint16").Call(Id("argBytes"), Uint16().Call(Id(seedRef))))
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id("argBytes")))
										case IdlTypeI32, IdlTypeU32:
											body.Add(Id("argBytes").Op(":=").Make(Index().Byte(), Lit(4)))
											body.Add(Qual(PkgBinary, "LittleEndian").Dot("PutUint32").Call(Id("argBytes"), Uint32().Call(Id(seedRef))))
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id("argBytes")))
										case IdlTypeI64, IdlTypeU64:
											body.Add(Id("argBytes").Op(":=").Make(Index().Byte(), Lit(8)))
											body.Add(Qual(PkgBinary, "LittleEndian").Dot("PutUint64").Call(Id("argBytes"), Uint64().Call(Id(seedRef))))
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id("argBytes")))
										case IdlTypeF32, IdlTypeF64:
											body.Add(Id("argBytes").Op(":=").Make(Index().Byte(), Lit(8)))
											body.Add(Qual(PkgBinary, "LittleEndian").Dot("PutUint64").Call(Id("argBytes"), Uint64().Call(Qual(PkgMath, "Float64bits").Call(Float64().Call(Id(seedRef))))))
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id("argBytes")))
										case IdlTypeI128, IdlTypeU128:
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id(seedRef).Index(Op(":"))))
										case IdlTypeString:
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id(seedRef)))
										default:
											body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id(seedRef).Dot("Bytes").Call()))
										}
									}
								}

								body.Line()

								seedProgramRef := Id("ProgramID")
								if seedProgramValue != nil {
									seedProgramRef = Id(getDefProgram("", solana.PublicKeyFromBytes(*seedProgramValue).String()))
								}

								body.Line()
								if seedProgramPath != "" {
									body.Add(
										List(Id("pda"), Id("bumpSeed"), Id("err")).Op("=").Add(Qual(PkgSolanaGo, "FindProgramAddress").Call(Id("seeds"), Id(seedProgramPath))),
									)
								} else {
									body.Add(
										List(Id("pda"), Id("bumpSeed"), Id("err")).Op("=").Add(Qual(PkgSolanaGo, "FindProgramAddress").Call(Id("seeds"), seedProgramRef)),
									)
								}

								body.Return()
							}))
					constantsCodeMap[internalAccessorName] = append(constantsCodeMap[internalAccessorName],
						Id("Must"+internalAccessorName).
							Params(
								ListFunc(func(params *Group) {
									for i, seedRef := range seedRefs {
										if seedRef != "" {
											if seedTypes[i].IsArray() && seedTypes[i].GetArray().Elem.GetString() == "u8" {
												params.Id(seedRef).Index(Lit(32)).Byte()
												continue
											}
											switch seedTypes[i].asString {
											case IdlTypeI8, IdlTypeI16, IdlTypeI32, IdlTypeI64, IdlTypeU8, IdlTypeU16, IdlTypeU32, IdlTypeU64, IdlTypeString:
												//是什么类型就是什么类型
												params.Id(seedRef).Add(genTypeName(seedTypes[i]))
											case IdlTypeI128, IdlTypeU128:
												params.Id(seedRef).Index(Lit(16)).Byte()
											default:
												params.Id(seedRef).Qual(PkgSolanaGo, "PublicKey")
											}

										}
									}
									if seedProgramPath != "" {
										params.Id(seedProgramPath).Qual(PkgSolanaGo, "PublicKey")
									}
								}),
							).
							Params(
								ListFunc(func(results *Group) {
									// Results:
									results.Id("pda").Qual(PkgSolanaGo, "PublicKey")
								}),
							).
							BlockFunc(func(body *Group) {
								// Body:
								body.Add(
									List(Id("pda"), Id("_"), Id("_")).Op("=").Add(Id(internalAccessorName).CallFunc(func(group *Group) {
										for i, seedRef := range seedRefs {
											if seedRef != "" {
												if seedTypes[i].IsArray() && seedTypes[i].GetArray().Elem.GetString() == "u8" {
													group.Id(seedRef)
												} else if seedTypes[i].asString == "i64" {
													group.Id(seedRef)
												} else {
													group.Id(seedRef)
												}
											}

										}
										if seedProgramPath != "" {
											group.Id(seedProgramPath)
										}
									})),
								)
								body.Return()
							}))
				}

			}
		}
	}

	{ // Create account getters:
		code.Line().Line()
		name := formatAccountAccessorName("Get", exportedAccountName)
		if account.Optional {
			code.Commentf("%s gets the %q account (optional).", name, account.Name).Line()
		} else {
			code.Commentf("%s gets the %q account.", name, account.Name).Line()
		}
		for _, doc := range account.Docs {
			code.Comment(doc).Line()
		}
		code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id(name).
			Params(
				ListFunc(func(params *Group) {
					// Parameters:
				}),
			).
			Params(
				ListFunc(func(results *Group) {
					// Results:
					results.Op("*").Qual(PkgSolanaGo, "AccountMeta")
				}),
			).
			BlockFunc(func(body *Group) {
				// Body:
				body.Return(Id("inst").Dot("AccountMetaSlice").Dot("Get").Call(Lit(index)))
			})
	}

	return code
}

func genProgramBoilerplate(idl IDL, addresses map[string]string) (*File, error) {
	file := NewGoFile(idl.Metadata.Name, true)
	for _, programDoc := range idl.Docs {
		file.HeaderComment(programDoc)
	}

	{
		// ProgramName variable:
		code := Empty()
		programName := ToCamel(idl.Metadata.Name)
		code.Const().Id("ProgramName").Op("=").Lit(programName)
		file.Add(code.Line())
	}

	// ProgramID variable:
	{
		code := Empty()
		code.Var().Id("ProgramID").Qual(PkgSolanaGo, "PublicKey").
			Add(
				func() Code {
					if idl.Address != "" {
						return Op("=").Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(idl.Address))
					}
					return nil
				}(),
			)
		file.Add(code.Line())
	}

	// `SetProgramID` func:
	{
		code := Empty()
		code.Func().Id("SetProgramID").Params(Id("PublicKey").Qual(PkgSolanaGo, "PublicKey")).Block(
			Id("ProgramID").Op("=").Id("PublicKey"),
			Qual(PkgSolanaGo, "RegisterInstructionDecoder").Call(Id("ProgramID"), Id("registryDecodeInstruction")),
		)
		file.Add(code.Line())
	}

	// register decoder:
	{
		code := Empty()
		code.Func().Id("init").Call().Block(
			If(
				Op("!").Id("ProgramID").Dot("IsZero").Call(),
			).Block(
				Qual(PkgSolanaGo, "RegisterInstructionDecoder").Call(Id("ProgramID"), Id("registryDecodeInstruction")),
			),
		)
		file.Add(code.Line())
	}

	// 注册所有的地址

	{
		code := Empty()
		code.Var().Parens(
			DoGroup(func(gr *Group) {
				var addrs [][]string
				for name, id := range addresses {
					addrs = append(addrs, []string{name, id})
				}
				sort.Slice(addrs, func(i, j int) bool {
					return addrs[i][0] < addrs[j][0]
				})
				for _, addr := range addrs {
					ins := Empty().Line()
					ins.Id(addr[0]).Op("=").Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(addr[1]))
					gr.Add(ins.Line().Line())
				}
			}),
		)
		file.Add(code.Line())
	}

	// Instruction ID enum:
	{
		GetConfig().TypeID.
			On(
				TypeIDNameSlice{
					TypeIDUvarint32,
					TypeIDUint32,
					TypeIDUint8,
				},
				func() {

					code := Empty()
					code.Const().Parens(
						DoGroup(func(gr *Group) {
							for instructionIndex, instruction := range idl.Instructions {
								insExportedName := ToCamel(instruction.Name)

								ins := Empty().Line()
								for _, doc := range instruction.Docs {
									ins.Comment(doc).Line()
								}
								ins.Id("Instruction_" + insExportedName)

								if instructionIndex == 0 {
									switch GetConfig().TypeID {
									case TypeIDUvarint32, TypeIDUint32:
										ins.Uint32().Op("=").Iota().Line()
									case TypeIDUint8:
										ins.Uint8().Op("=").Iota().Line()
									}
								}
								gr.Add(ins.Line().Line())
							}
						}),
					)
					file.Add(code.Line())

				},
			).
			On(
				TypeIDNameSlice{
					TypeIDAnchor,
				},
				func() {
					code := Empty()
					code.Var().Parens(
						DoGroup(func(gr *Group) {
							for _, instruction := range idl.Instructions {
								insExportedName := ToCamel(instruction.Name)

								ins := Empty().Line()
								for _, doc := range instruction.Docs {
									ins.Comment(doc).Line()
								}
								toBeHashed := sighash.ToSnakeForSighash(instruction.Name)
								if GetConfig().Debug {
									ins.Comment(Sf(`hash("%s:%s")`, bin.SIGHASH_GLOBAL_NAMESPACE, toBeHashed)).Line()
								}
								ins.Id("Instruction_" + insExportedName)

								ins.Op("=").Qual(PkgDfuseBinary, "TypeID").Call(
									Index(Lit(8)).Byte().Op("{").ListFunc(func(byteGroup *Group) {
										sighash := bin.SighashTypeID(bin.SIGHASH_GLOBAL_NAMESPACE, toBeHashed)
										if instruction.Discriminator != nil {
											sighash = *instruction.Discriminator
										}
										for _, byteVal := range sighash[:] {
											byteGroup.Lit(int(byteVal))
										}
									}).Op("}"),
								)
								gr.Add(ins.Line().Line())
							}
						}),
					)
					file.Add(code.Line())
				},
			).
			On(
				TypeIDNameSlice{
					TypeIDNoType,
				},
				func() {
					// TODO
				},
			)
	}

	// Declare `InstructionIDToName` function:
	{
		GetConfig().TypeID.
			On(
				TypeIDNameSlice{
					TypeIDUvarint32,
					TypeIDUint32,
					TypeIDUint8,
				},
				func() {
					code := Empty()
					code.Comment("InstructionIDToName returns the name of the instruction given its ID.").Line()
					code.Func().Id("InstructionIDToName").
						Params(
							func() Code {
								switch GetConfig().TypeID {
								case TypeIDUvarint32, TypeIDUint32:
									return Id("id").Uint32()
								case TypeIDUint8:
									return Id("id").Uint8()
								}
								return nil
							}(),
						).
						Params(String()).
						BlockFunc(func(body *Group) {
							body.Switch(Id("id")).BlockFunc(func(switchBlock *Group) {
								for _, instruction := range idl.Instructions {
									insExportedName := ToCamel(instruction.Name)
									switchBlock.Case(Id("Instruction_" + insExportedName)).Line().Return(Lit(insExportedName))
								}
								switchBlock.Default().Line().Return(Lit(""))
							})

						})

					file.Add(code.Line())
				},
			).
			On(
				TypeIDNameSlice{
					TypeIDAnchor,
				},
				func() {
					code := Empty()
					code.Comment("InstructionIDToName returns the name of the instruction given its ID.").Line()
					code.Func().Id("InstructionIDToName").
						Params(Id("id").Qual(PkgDfuseBinary, "TypeID")).
						Params(String()).
						BlockFunc(func(body *Group) {
							body.Switch(Id("id")).BlockFunc(func(switchBlock *Group) {
								for _, instruction := range idl.Instructions {
									insExportedName := ToCamel(instruction.Name)
									switchBlock.Case(Id("Instruction_" + insExportedName)).Line().Return(Lit(insExportedName))
								}
								switchBlock.Default().Line().Return(Lit(""))
							})

						})
					file.Add(code.Line())
				},
			).
			On(
				TypeIDNameSlice{
					TypeIDNoType,
				},
				func() {
					// TODO
				},
			)
	}
	// Base Instruction struct:
	{
		{
			code := Empty()
			code.Type().Id("Instruction").Struct(
				Qual(PkgDfuseBinary, "BaseVariant"),
			)
			file.Add(code.Line())
		}
		{
			// `EncodeToTree(parent treeout.Branches)` method
			code := Empty()
			code.Func().Parens(Id("inst").Op("*").Id("Instruction")).Id("EncodeToTree").
				Params(Id("parent").Qual(PkgTreeout, "Branches")).
				Params().
				BlockFunc(func(body *Group) {
					body.If(
						List(Id("enToTree"), Id("ok")).Op(":=").Id("inst").Dot("Impl").Op(".").Parens(Qual(PkgSolanaGoText, "EncodableToTree")).
							Op(";").
							Id("ok"),
					).Block(
						Id("enToTree").Dot("EncodeToTree").Call(Id("parent")),
					).Else().Block(
						Id("parent").Dot("Child").Call(Qual("github.com/davecgh/go-spew/spew", "Sdump").Call(Id("inst"))),
					)
				})
			file.Add(code.Line())
		}
		{
			// variant definitions for the decoder:
			GetConfig().TypeID.
				On(
					TypeIDNameSlice{
						TypeIDUvarint32,
						TypeIDUint32,
						TypeIDUint8,
					},
					func() {

						code := Empty()
						code.Var().Id("InstructionImplDef").Op("=").Qual(PkgDfuseBinary, "NewVariantDefinition").
							Parens(DoGroup(func(call *Group) {
								call.Line()

								switch GetConfig().TypeID {
								case TypeIDUvarint32:
									call.Qual(PkgDfuseBinary, "Uvarint32TypeIDEncoding").Op(",").Line()
								case TypeIDUint32:
									call.Qual(PkgDfuseBinary, "Uint32TypeIDEncoding").Op(",").Line()
								case TypeIDUint8:
									call.Qual(PkgDfuseBinary, "Uint8TypeIDEncoding").Op(",").Line()
								}

								call.Index().Qual(PkgDfuseBinary, "VariantType").
									BlockFunc(func(variantBlock *Group) {
										for _, instruction := range idl.Instructions {
											// NOTE: using `ToCamel` here:
											insName := ToCamel(instruction.Name)
											insExportedName := ToCamel(instruction.Name)
											variantBlock.Block(
												List(Lit(insName), Parens(Op("*").Id(insExportedName)).Parens(Nil())).Op(","),
											).Op(",")
										}
									}).Op(",").Line()
							}))
						file.Add(code.Line())

					},
				).
				On(
					TypeIDNameSlice{
						TypeIDAnchor,
					},
					func() {
						code := Empty()
						code.Var().Id("InstructionImplDef").Op("=").Qual(PkgDfuseBinary, "NewVariantDefinition").
							Parens(DoGroup(func(call *Group) {
								call.Line()
								call.Qual(PkgDfuseBinary, "AnchorTypeIDEncoding").Op(",").Line()

								call.Index().Qual(PkgDfuseBinary, "VariantType").
									BlockFunc(func(variantBlock *Group) {
										for _, instruction := range idl.Instructions {
											// NOTE: using `ToSnakeForSighash` here (necessary for sighash computing from instruction name)
											insName := sighash.ToSnakeForSighash(instruction.Name)
											insExportedName := ToCamel(instruction.Name)
											variantBlock.Block(
												List(Id("Name").Op(":").Lit(insName), Id("Type").Op(":").Parens(Op("*").Id(insExportedName)).Parens(Nil())).Op(","),
											).Op(",")
										}
									}).Op(",").Line()
							}))
						file.Add(code.Line())
					},
				).
				On(
					TypeIDNameSlice{
						TypeIDNoType,
					},
					func() {
						// TODO
					},
				)

		}
		{
			// method to return programID:
			code := Empty()
			code.Func().Parens(Id("inst").Op("*").Id("Instruction")).Id("ProgramID").Params().
				Parens(Qual(PkgSolanaGo, "PublicKey")).
				BlockFunc(func(body *Group) {
					body.Return(
						Id("ProgramID"),
					)
				})
			file.Add(code.Line())
		}
		{
			// method to return accounts:
			code := Empty()
			code.Func().Parens(Id("inst").Op("*").Id("Instruction")).Id("Accounts").Params().
				Parens(Id("out").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")).
				BlockFunc(func(body *Group) {
					body.Return(
						Id("inst").Dot("Impl").Op(".").Parens(Qual(PkgSolanaGo, "AccountsGettable")).Dot("GetAccounts").Call(),
					)
				})
			file.Add(code.Line())
		}
		{
			// `Data() ([]byte, error)` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("Data").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Index().Byte()
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Id("buf").Op(":=").New(Qual("bytes", "Buffer"))

					body.If(
						Err().Op(":=").Qual(PkgDfuseBinary, GetConfig().Encoding._NewEncoder()).Call(Id("buf")).Dot("Encode").Call(Id("inst")).
							Op(";").
							Err().Op("!=").Nil(),
					).Block(
						Return(List(Nil(), Qual("fmt", "Errorf").Call(Lit("unable to encode instruction: %w"), Err()))),
					)
					body.Return(Id("buf").Dot("Bytes").Call(), Nil())
				})
			file.Add(code.Line())
		}
		{
			// `TextEncode(encoder *text.Encoder, option *text.Option) error` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("TextEncode").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("encoder").Op("*").Qual(PkgSolanaGoText, "Encoder")
						params.Id("option").Op("*").Qual(PkgSolanaGoText, "Option")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Return(Id("encoder").Dot("Encode").Call(Id("inst").Dot("Impl"), Id("option")))
				})
			file.Add(code.Line())
		}
		{
			// `UnmarshalWithDecoder(decoder *bin.Decoder) error` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("UnmarshalWithDecoder").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("decoder").Op("*").Qual(PkgDfuseBinary, "Decoder")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Return(Id("inst").Dot("BaseVariant").Dot("UnmarshalBinaryVariant").Call(Id("decoder"), Id("InstructionImplDef")))
				})
			file.Add(code.Line())
		}
		{
			// `MarshalWithEncoder(encoder *bin.Encoder) error ` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("MarshalWithEncoder").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("encoder").Op("*").Qual(PkgDfuseBinary, "Encoder")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					GetConfig().TypeID.
						On(
							TypeIDNameSlice{
								TypeIDUvarint32,
								TypeIDUint32,
								TypeIDUint8,
							},
							func() {

								switch GetConfig().TypeID {
								case TypeIDUvarint32:
									body.Err().Op(":=").Id("encoder").Dot("WriteUVarInt").Call(Id("inst").Dot("TypeID").Dot("Uvarint32").Call())
								case TypeIDUint32:
									body.Err().Op(":=").Id("encoder").Dot("WriteUint32").Call(Id("inst").Dot("TypeID").Dot("Uint32").Call(), Qual("encoding/binary", "LittleEndian"))
								case TypeIDUint8:
									body.Err().Op(":=").Id("encoder").Dot("WriteUint8").Call(Id("inst").Dot("TypeID").Dot("Uint8").Call())
								}

							},
						).
						On(
							TypeIDNameSlice{
								TypeIDAnchor,
							},
							func() {
								body.Err().Op(":=").Id("encoder").Dot("WriteBytes").Call(Id("inst").Dot("TypeID").Dot("Bytes").Call(), False())
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDNoType,
							},
							func() {
								// TODO
							},
						)

					body.If(
						Err().Op("!=").Nil(),
					).Block(
						Return(Qual("fmt", "Errorf").Call(Lit("unable to write variant type: %w"), Err())),
					)
					body.Return(Id("encoder").Dot("Encode").Call(Id("inst").Dot("Impl")))
				})
			file.Add(code.Line())
		}
		{
			// `registryDecodeInstruction` func:
			code := Empty()
			code.Func().Id("registryDecodeInstruction").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("accounts").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")
						params.Id("data").Index().Byte()
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Interface()
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.List(Id("inst"), Err()).Op(":=").Id("decodeInstruction").Call(Id("accounts"), Id("data"))

					body.If(
						Err().Op("!=").Nil(),
					).Block(
						Return(Nil(), Err()),
					)
					body.Return(Id("inst"), Nil())
				})
			file.Add(code.Line())
		}
		{
			// `DecodeInstruction` func:
			code := Empty()
			code.Func().Id("decodeInstruction").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("accounts").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")
						params.Id("data").Index().Byte()
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id("Instruction")
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					body.Id("inst").Op(":=").New(Id("Instruction"))

					body.If(
						Err().Op(":=").Qual(PkgDfuseBinary, GetConfig().Encoding._NewDecoder()).Call(Id("data")).Dot("Decode").Call(Id("inst")).
							Op(";").
							Err().Op("!=").Nil(),
					).Block(
						Return(
							Nil(),
							Qual("fmt", "Errorf").Call(Lit("unable to decode instruction: %w"), Err()),
						),
					)

					body.If(

						List(Id("v"), Id("ok")).Op(":=").Id("inst").Dot("Impl").Op(".").Parens(Qual(PkgSolanaGo, "AccountsSettable")).
							Op(";").
							Id("ok"),
					).BlockFunc(func(gr *Group) {
						gr.Err().Op(":=").Id("v").Dot("SetAccounts").Call(Id("accounts"))
						gr.If(Err().Op("!=").Nil()).Block(
							Return(
								Nil(),
								Qual("fmt", "Errorf").Call(Lit("unable to set accounts for instruction: %w"), Err()),
							),
						)
					})

					body.Return(Id("inst"), Nil())
				})
			file.Add(code.Line())
		}
	}

	return file, nil
}

func formatAccountAccessorName(prefix, name string) string {
	endsWithAccount := strings.HasSuffix(strings.ToLower(name), "account")
	if !conf.RemoveAccountSuffix || !endsWithAccount {
		return prefix + name + "Account"
	}
	return prefix + name
}

func treeFindLongestNameFromFields(fields []IdlField) (ln int) {
	for _, v := range fields {
		if len(v.Name) > ln {
			ln = len(v.Name)
		}
	}
	return
}

func treeFindLongestNameFromAccounts(accounts IdlAccountItemSlice) (ln int) {
	accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {

		cleanedName := treeFormatAccountName(ia.Name)

		exportedAccountName := filepath.Join(groupPath, cleanedName)
		if len(exportedAccountName) > ln {
			ln = len(exportedAccountName)
		}

		return true
	})
	return
}

func treeFormatAccountName(name string) string {
	cleanedName := name
	if isSysVar(name) {
		cleanedName = strings.TrimSuffix(getSysVarName(name), "PublicKey")
	}
	if len(cleanedName) > len("account") {
		if strings.HasSuffix(cleanedName, "account") {
			cleanedName = strings.TrimSuffix(cleanedName, "account")
		} else if strings.HasSuffix(cleanedName, "Account") {
			cleanedName = strings.TrimSuffix(cleanedName, "Account")
		}
	}
	return cleanedName
}
