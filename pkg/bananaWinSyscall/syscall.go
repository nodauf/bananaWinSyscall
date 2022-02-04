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
//dsys NtProtectVirtualMemory(processHandle windows.Handle, baseAddress *uintptr, numberOfBytesToProtect *uintptr, newAccessProtection int64, OldAccessProtection *int64) (err error)

/*
__kernel_entry NTSYSCALLAPI NTSTATUS NtCreateThreadEx(
    OUT PHANDLE                 hThread,
    IN  ACCESS_MASK             DesiredAccess,
    IN  POBJECT_ATTRIBUTES      ObjectAttributes,
    IN  HANDLE                  ProcessHandle,
    IN  LPTHREAD_START_ROUTINE  lpStartAddress,
    IN  LPVOID                  lpParameter,
    IN  BOOL                    CreateSuspended,
    IN  DWORD                   StackZeroBits,
    IN  DWORD                   SizeOfStackCommit,
    IN  DWORD                   SizeOfStackReserve,
    OUT CREATE_THREAD_INFO      lpBytesBuffer
);*/
//dsys NtCreateThreadEx(threadHandle *windows.Handle, desiredAccess windows.ACCESS_MASK, objectAttributes *windows.OBJECT_ATTRIBUTES, processHandle windows.Handle, startAddress uintptr, parameter uintptr, createSuspended bool, stackZeroBits uint32, sizeOfStackCommit uint32, sizeOfStackReserve uint32, lpbytesbuffer uint32) (err error)

/*
NtWriteVirtualMemory(
  IN HANDLE               ProcessHandle,
  IN PVOID                BaseAddress,
  IN PVOID                Buffer,
  IN ULONG                NumberOfBytesToWrite,
  OUT PULONG              NumberOfBytesWritten OPTIONAL );*/
//dsys NtWriteVirtualMemory(processHandle windows.Handle, baseAddress uintptr, buffer *byte, numberOfBytesToWrite uint32, numberOfBytesWritten *uint32) (err error)

//go:generate go run github.com/nodauf/bananaWinSyscall/mkdirectwinsyscall -output zsyscall_windows.go syscall.go
