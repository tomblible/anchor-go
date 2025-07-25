// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package vr

import ag_solanago "github.com/gagliardetto/solana-go"

func FindPoolAddress(owner ag_solanago.PublicKey, mintA ag_solanago.PublicKey, mintB ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: 0x706f6f6c
	seeds = append(seeds, []byte{byte(0x70), byte(0x6f), byte(0x6f), byte(0x6c)})
	// path: owner
	seeds = append(seeds, owner.Bytes())
	// path: mintA
	seeds = append(seeds, mintA.Bytes())
	// path: mintB
	seeds = append(seeds, mintB.Bytes())

	pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	return
}

func MustFindPoolAddress(owner ag_solanago.PublicKey, mintA ag_solanago.PublicKey, mintB ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, _ = FindPoolAddress(owner, mintA, mintB)
	return
}

func FindVaultAAddress(pool ag_solanago.PublicKey, mintA ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: pool
	seeds = append(seeds, pool.Bytes())
	// path: mintA
	seeds = append(seeds, mintA.Bytes())

	pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	return
}

func MustFindVaultAAddress(pool ag_solanago.PublicKey, mintA ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, _ = FindVaultAAddress(pool, mintA)
	return
}

func FindVaultBAddress(pool ag_solanago.PublicKey, mintB ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: pool
	seeds = append(seeds, pool.Bytes())
	// path: mintB
	seeds = append(seeds, mintB.Bytes())

	pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	return
}

func MustFindVaultBAddress(pool ag_solanago.PublicKey, mintB ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, _ = FindVaultBAddress(pool, mintB)
	return
}
