// Copyright 2025 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 && tinygo

package cpuid

// Stub implementation for TinyGo: return empty/zero values for everything.
// CPU feature detection is a nice-to-have; u-root runs fine without it.
// CGo-based CPUID (via import "C") is not available under gobusybox's
// package discovery phase (go list cannot process import "C" files).

func cpuid_low(arg1, arg2 uint32) (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}

func xgetbv_low(arg1 uint32) (eax, edx uint32) {
	return 0, 0
}
