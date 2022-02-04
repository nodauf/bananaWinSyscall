package bananawinsyscall

import "golang.org/x/sys/windows"

type ClientID struct {
	UniqueProcess windows.Handle
	UniqueThread  windows.Handle
}

/*
__kernel_entry NTSYSCALLAPI NTSTATUS NtAllocateVirtualMemory(
  HANDLE    ProcessHandle,
  PVOID     *BaseAddress,
  ULONG_PTR ZeroBits,
  PSIZE_T   RegionSize,
  ULONG     AllocationType,
  ULONG     Protect
);
*/
//dsys NtAllocateVirtualMemory(processHandle windows.Handle, baseAddress *uintptr, zeroBits uintptr, regionSize *uintptr, allocationType uint64, protect uint64) (err error)

/*__kernel_entry NTSYSCALLAPI NTSTATUS NtOpenProcess(
  [out]          PHANDLE            ProcessHandle,
  [in]           ACCESS_MASK        DesiredAccess,
  [in]           POBJECT_ATTRIBUTES ObjectAttributes,
  [in, optional] PCLIENT_ID         ClientId
);*/
//dsys NtOpenProcess(processHandle *windows.Handle, desiredAccess windows.ACCESS_MASK, objectAttributes *windows.OBJECT_ATTRIBUTES, clientID *ClientID) (err error)

/*
__kernel_entry NTSYSCALLAPI NTSTATUS NtProtectVirtualMemory(
  IN HANDLE               ProcessHandle,
  IN OUT PVOID            *BaseAddress,
  IN OUT PULONG           NumberOfBytesToProtect,
  IN ULONG                NewAccessProtection,
  OUT PULONG              OldAccessProtection );*/
//dsys NtProtectVirtualMemory(processHandle windows.Handle, baseAddress *uintptr, numberOfBytesToProtect *int64, newAccessProtection int64, OldAccessProtection *int64) (err error)

//go:generate go run github.com/nodauf/bananaWinSyscall/mkdirectwinsyscall -output zsyscall_windows.go syscall.go
