// Copyright 2024 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (tinygo || tinygo.enable) && amd64

package universalpayload

import "errors"

// Stub for TinyGo: UPP boot uses x86 firmware calls via CGo,
// which is not available under gobusybox.

func uartInit(port uint16)             {}
func uartPutchar(c byte)                {}
func uartGetchar() byte                 { return 0 }
func uartWasInput() bool                { return false }
func setupScreens()                     {}
func startCore(cpu uint, entry, dt uint64) {}
func addrOfStart() uintptr    { return 0 }
func addrOfStackTop() uintptr { return 0 }
func addrOfHobAddr() uintptr  { return 0 }
func getPhysicalAddressSizes() (uint8, error)  { return 0, errors.New("not implemented") }
func constructTrampoline(buf []uint8, addr uint64, entry uint64) []uint8 { return buf }
func archGetAcpiRsdpData() (uint64, []byte, error) { return 0, nil, errors.New("not implemented") }
func appendAddonMemMap(_ *EFIMemoryMapHOB) uint64  { return 0 }
func isMemReserved(memType string) bool             { return false }
