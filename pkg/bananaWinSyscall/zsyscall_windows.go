// Code generated by 'go generate'; DO NOT EDIT.

package bananawinsyscall

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"

	bananaphone "github.com/C-Sto/BananaPhone/pkg/BananaPhone"
)

var _ unsafe.Pointer

var (
	bpGlobal, bperr = bananaphone.NewBananaPhone(bananaphone.AutoBananaPhoneMode)
)

func NtAllocateVirtualMemory(processHandle windows.Handle, baseAddress *uintptr, zeroBits uintptr, regionSize *uintptr, allocationType uint64, protect uint64) (err error) {
	if bpGlobal == nil {
		err = fmt.Errorf("BananaPhone uninitialised: %s", bperr.Error())
		return
	}

	sysid, e := bpGlobal.GetSysID("NtAllocateVirtualMemory")
	if e != nil {
		err = e
		return
	}
	r1, _ := bananaphone.Syscall(sysid, uintptr(processHandle), uintptr(unsafe.Pointer(baseAddress)), uintptr(zeroBits), uintptr(unsafe.Pointer(regionSize)), uintptr(allocationType), uintptr(protect))
	if r1 != 0 {
		err = fmt.Errorf("error code: %x", r1)
	}
	return
}

func NtOpenProcess(processHandle *windows.Handle, desiredAccess windows.ACCESS_MASK, objectAttributes *windows.OBJECT_ATTRIBUTES, clientID *ClientID) (err error) {
	if bpGlobal == nil {
		err = fmt.Errorf("BananaPhone uninitialised: %s", bperr.Error())
		return
	}

	sysid, e := bpGlobal.GetSysID("NtOpenProcess")
	if e != nil {
		err = e
		return
	}
	r1, _ := bananaphone.Syscall(sysid, uintptr(unsafe.Pointer(processHandle)), uintptr(desiredAccess), uintptr(unsafe.Pointer(objectAttributes)), uintptr(unsafe.Pointer(clientID)))
	if r1 != 0 {
		err = fmt.Errorf("error code: %x", r1)
	}
	return
}

func NtProtectVirtualMemory(processHandle windows.Handle, baseAddress *uintptr, numberOfBytesToProtect *int64, newAccessProtection int64, OldAccessProtection *int64) (err error) {
	if bpGlobal == nil {
		err = fmt.Errorf("BananaPhone uninitialised: %s", bperr.Error())
		return
	}

	sysid, e := bpGlobal.GetSysID("NtProtectVirtualMemory")
	if e != nil {
		err = e
		return
	}
	r1, _ := bananaphone.Syscall(sysid, uintptr(processHandle), uintptr(unsafe.Pointer(baseAddress)), uintptr(unsafe.Pointer(numberOfBytesToProtect)), uintptr(newAccessProtection), uintptr(unsafe.Pointer(OldAccessProtection)))
	if r1 != 0 {
		err = fmt.Errorf("error code: %x", r1)
	}
	return
}
