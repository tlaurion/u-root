// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (tinygo || tinygo.enable) && linux && amd64

package trampoline

import "errors"

// Stub for TinyGo: trampoline booting uses x86 real-mode assembly via CGo,
// which is not available under gobusybox's package discovery (go list cannot
// process import "C").

// Setup creates the trampoline binary. Stub for TinyGo.
func Setup(path string, magic, infoAddr, entryPoint uintptr) ([]byte, error) {
	return nil, errors.New("trampoline not supported under TinyGo")
}

func jumpToAddr(addr uintptr) { panic("unimplemented: tinygo trampoline") }
func GdtFlush(gdtPtr uintptr) { panic("unimplemented: tinygo gdt flush") }
func TssFlush()               { panic("unimplemented: tinygo tss flush") }
func InitSmp(...uint32)       { panic("unimplemented: tinygo smp init") }
func LoadSegmentRegisters()   { panic("unimplemented: tinygo load segment registers") }

func JumpToKernel(kernEntry uintptr, magic, addr uint32, info uintptr) error {
	return errors.New("JumpToKernel not supported under TinyGo")
}
func JumpToKernel16(kernEntry uintptr, info uintptr) error {
	return errors.New("JumpToKernel16 not supported under TinyGo")
}
func JumpToKernel32(kernEntry uintptr, info uintptr) error {
	return errors.New("JumpToKernel32 not supported under TinyGo")
}
func WriteTableEntry(e820 *[4096]uint8, data []byte, entrySize int) int { return 0 }
func A20Enable() bool { return false }
