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
NTSTATUS NtWriteVirtualMemory(
  IN HANDLE               ProcessHandle,
  IN PVOID                BaseAddress,
  IN PVOID                Buffer,
  IN ULONG                NumberOfBytesToWrite,
  OUT PULONG              NumberOfBytesWritten OPTIONAL );*/
//dsys NtWriteVirtualMemory(processHandle windows.Handle, baseAddress uintptr, buffer *byte, numberOfBytesToWrite uintptr, numberOfBytesWritten *uint32) (err error)

/*
NTSTATUS NtResumeThread(
  IN  HANDLE ThreadHandle,
  IN  PULONG PreviousSuspendCount
);*/
//dsys NtResumeThread(threadHandle windows.Handle, previousSuspendCount *uint32) (err error)

/*
NTSTATUS NtCreateProcess(
  OUT PHANDLE           ProcessHandle,
  IN ACCESS_MASK        DesiredAccess,
  IN POBJECT_ATTRIBUTES ObjectAttributes OPTIONAL,
  IN HANDLE             ParentProcess,
  IN BOOLEAN            InheritObjectTable,
  IN HANDLE             SectionHandle OPTIONAL,
  IN HANDLE             DebugPort OPTIONAL,
  IN HANDLE             ExceptionPort OPTIONAL
);*/
//dsys NtCreateProcess(processHandle *windows.Handle, desiredAccess windows.ACCESS_MASK, objectAttributes *windows.OBJECT_ATTRIBUTES, parentProcess windows.Handle, inheritObjectTable bool, sectionHandle windows.Handle, debugPort windows.Handle, exceptionPort windows.Handle) (err error)

/*
__kernel_entry NTSYSCALLAPI NTSTATUS NtOpenFile(
  [out] PHANDLE            FileHandle,
  [in]  ACCESS_MASK        DesiredAccess,
  [in]  POBJECT_ATTRIBUTES ObjectAttributes,
  [out] PIO_STATUS_BLOCK   IoStatusBlock,
  [in]  ULONG              ShareAccess,
  [in]  ULONG              OpenOptions
);*/
//dsys NtOpenFile(fileHandle *windows.Handle, desiredAccess windows.ACCESS_MASK, objectAttributes *windows.OBJECT_ATTRIBUTES, ioStatusBlock *windows.IO_STATUS_BLOCK, shareAccess uint64, openOptions uint64) (err error)

/*
NTSTATUS NtCreateSection (
  _Out_ PHANDLE SectionHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
  _In_opt_ PLARGE_INTEGER MaximumSize,
  _In_ ULONG SectionPageProtection,
  _In_ ULONG AllocationAttributes,
  _In_opt_ HANDLE FileHandle
);*/
//dsys NtCreateSection(sectionHandle *windows.Handle, desiredAccess windows.ACCESS_MASK, objectAttributes *windows.OBJECT_ATTRIBUTES, maximumSize *int64, sectionPageProtection uint32, allocationAttributes uint32, fileHandle windows.Handle) (err error)

/*
NTSTATUS NtClose(
      _In_ HANDLE Handle
);*/
//dsys NtClose(handle *windows.Handle) (err error)

//go:generate go run github.com/nodauf/bananaWinSyscall/mkdirectwinsyscall -output zsyscall_windows.go syscall.go
