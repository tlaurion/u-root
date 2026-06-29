// Copyright 2024 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (tinygo || tinygo.enable) && linux && (amd64 || 386)

package memio

import (
	"log"
	"os"
	"sync"
	"unsafe"
)

// TinyGo-compatible port I/O via /dev/port file.
// Direct in/out instructions via CGo are not available under gobusybox
// (go list cannot process import "C" in the package discovery phase).
// The iopl() syscall is already called by ArchPort.In/Out in archport_linux.go.
//
// /dev/port handle is lazily opened and cached to avoid repeated open/close.

var (
	portOnce sync.Once
	portFile *os.File
)

func getPort() *os.File {
	portOnce.Do(func() {
		f, err := os.OpenFile("/dev/port", os.O_RDWR, 0)
		if err != nil {
			f, err = os.OpenFile("/dev/port", os.O_RDONLY, 0)
			if err != nil {
				portFile = nil
				return
			}
		}
		portFile = f
	})
	return portFile
}

func archInl(port uint16) uint32 {
	f := getPort()
	if f == nil {
		return 0
	}
	var data uint32
	buf := (*[4]byte)(unsafe.Pointer(&data))[:]
	n, err := f.ReadAt(buf, int64(port))
	if err != nil || n != len(buf) {
		log.Printf("archInl(0x%x): read %d/%d bytes: %v", port, n, len(buf), err)
		return 0
	}
	return data
}

func archInw(port uint16) uint16 {
	f := getPort()
	if f == nil {
		return 0
	}
	var data uint16
	buf := (*[2]byte)(unsafe.Pointer(&data))[:]
	n, err := f.ReadAt(buf, int64(port))
	if err != nil || n != len(buf) {
		log.Printf("archInw(0x%x): read %d/%d bytes: %v", port, n, len(buf), err)
		return 0
	}
	return data
}

func archInb(port uint16) uint8 {
	f := getPort()
	if f == nil {
		return 0
	}
	var data uint8
	buf := (*[1]byte)(unsafe.Pointer(&data))[:]
	n, err := f.ReadAt(buf, int64(port))
	if err != nil || n != len(buf) {
		log.Printf("archInb(0x%x): read %d/%d bytes: %v", port, n, len(buf), err)
		return 0
	}
	return data
}

func archOutl(port uint16, val uint32) {
	f := getPort()
	if f == nil {
		return
	}
	buf := (*[4]byte)(unsafe.Pointer(&val))[:]
	n, err := f.WriteAt(buf, int64(port))
	if err != nil || n != len(buf) {
		log.Printf("archOutl(0x%x): wrote %d/%d bytes: %v", port, n, len(buf), err)
	}
}

func archOutw(port uint16, val uint16) {
	f := getPort()
	if f == nil {
		return
	}
	buf := (*[2]byte)(unsafe.Pointer(&val))[:]
	n, err := f.WriteAt(buf, int64(port))
	if err != nil || n != len(buf) {
		log.Printf("archOutw(0x%x): wrote %d/%d bytes: %v", port, n, len(buf), err)
	}
}

func archOutb(port uint16, val uint8) {
	f := getPort()
	if f == nil {
		return
	}
	buf := (*[1]byte)(unsafe.Pointer(&val))[:]
	n, err := f.WriteAt(buf, int64(port))
	if err != nil || n != len(buf) {
		log.Printf("archOutb(0x%x): wrote %d/%d bytes: %v", port, n, len(buf), err)
	}
}
