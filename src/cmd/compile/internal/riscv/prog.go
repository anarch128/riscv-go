package riscv

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/riscv"
)

// This table gives the basic information about instruction generated by the
// compiler and processed in the optimizer.  Instructions not generated may be
// omitted.
//
// NOTE(prattmic): I believe that the gc.Size flags are used only for non-SSA
// peephole optimizations, and can thus be omitted for RISCV.
var progmap = map[obj.As]obj.ProgInfo{
	obj.ATYPE:     {Flags: gc.Pseudo | gc.Skip},
	obj.ATEXT:     {Flags: gc.Pseudo},
	obj.AFUNCDATA: {Flags: gc.Pseudo},
	obj.APCDATA:   {Flags: gc.Pseudo},
	obj.AUNDEF:    {Flags: gc.Break},
	obj.AUSEFIELD: {Flags: gc.OK},
	obj.ACHECKNIL: {Flags: gc.LeftRead},
	obj.AVARDEF:   {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARKILL:  {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARLIVE:  {Flags: gc.Pseudo | gc.LeftRead},
	obj.ARET:      {Flags: gc.Break},
	obj.AJMP:      {Flags: gc.Jump | gc.Break | gc.KillCarry},
	obj.ACALL:     {Flags: gc.RightAddr | gc.Call | gc.KillCarry},

	// NOP is an internal no-op that also stands for USED and SET
	// annotations.
	obj.ANOP: {Flags: gc.LeftRead | gc.RightWrite},

	// RISCV simple three operand instructions
	riscv.AADD:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AAND:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AMUL:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AMULW:    {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AMULH:    {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AMULHU:   {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ADIV:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ADIVU:    {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ADIVW:    {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ADIVUW:   {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AREM:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AREMU:    {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AREMW:    {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AREMUW:   {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AOR:      {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ASLL:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ASLT:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ASLTU:    {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ASRA:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ASRL:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.ASUB:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AXOR:     {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AFADDS:   {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AFSUBS:   {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AFMULS:   {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AFDIVS:   {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AFSGNJS:  {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AFSGNJNS: {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},
	riscv.AFSGNJXS: {Flags: gc.LeftRead | gc.RegRead | gc.RightWrite},

	// RISC-V register-immediate instructions
	riscv.AADDI:  {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AANDI:  {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AORI:   {Flags: gc.LeftRead | gc.RightWrite},
	riscv.ASLLI:  {Flags: gc.LeftRead | gc.RightWrite},
	riscv.ASLTI:  {Flags: gc.LeftRead | gc.RightWrite},
	riscv.ASLTIU: {Flags: gc.LeftRead | gc.RightWrite},
	riscv.ASRLI:  {Flags: gc.LeftRead | gc.RightWrite},
	riscv.ASRAI:  {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AXORI:  {Flags: gc.LeftRead | gc.RightWrite},

	// RISCV moves, loads, and stores
	riscv.ALD:    {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.ASD:    {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.AMOV:   {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.AMOVB:  {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.AMOVBU: {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.AMOVH:  {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.AMOVHU: {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.AMOVW:  {Flags: gc.LeftRead | gc.RightWrite | gc.Move},
	riscv.AMOVWU: {Flags: gc.LeftRead | gc.RightWrite | gc.Move},

	// Other RISC-V instructions
	riscv.ASEQZ:   {Flags: gc.LeftRead | gc.RightWrite},
	riscv.ASNEZ:   {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AFSQRTS: {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AFNEGS:  {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AFCVTSW: {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AFCVTSL: {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AFCVTWS: {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AFCVTLS: {Flags: gc.LeftRead | gc.RightWrite},
	riscv.AECALL:  {Flags: gc.OK},

	// RISCV conditional branches
	riscv.ABEQ:  {Flags: gc.Cjmp | gc.LeftRead | gc.RegRead},
	riscv.ABNE:  {Flags: gc.Cjmp | gc.LeftRead | gc.RegRead},
	riscv.ABGE:  {Flags: gc.Cjmp | gc.LeftRead | gc.RegRead},
	riscv.ABGEU: {Flags: gc.Cjmp | gc.LeftRead | gc.RegRead},
	riscv.ABLT:  {Flags: gc.Cjmp | gc.LeftRead | gc.RegRead},
	riscv.ABLTU: {Flags: gc.Cjmp | gc.LeftRead | gc.RegRead},
}

func proginfo(p *obj.Prog) {
	info, ok := progmap[p.As]
	if !ok {
		p.Ctxt.Diag("proginfo missing prog %v", p.As)
		return
	}

	p.Info = info
}
