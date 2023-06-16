// Code generated by 'go generate'; DO NOT EDIT.

package windows

import (
	"syscall"
	"unsafe"

	"golang.org/x/net v0.7.0x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modpsapi    = windows.NewLazySystemDLL("psapi.dll")
	modntdll    = windows.NewLazySystemDLL("ntdll.dll")
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	procGlobalMemoryStatusEx             = modkernel32.NewProc("GlobalMemoryStatusEx")
	procGetLogicalDriveStringsW          = modkernel32.NewProc("GetLogicalDriveStringsW")
	procGetProcessMemoryInfo             = modpsapi.NewProc("GetProcessMemoryInfo")
	procGetProcessImageFileNameW         = modpsapi.NewProc("GetProcessImageFileNameW")
	procGetSystemTimes                   = modkernel32.NewProc("GetSystemTimes")
	procGetDriveTypeW                    = modkernel32.NewProc("GetDriveTypeW")
	procEnumProcesses                    = modpsapi.NewProc("EnumProcesses")
	procGetDiskFreeSpaceExW              = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procProcess32FirstW                  = modkernel32.NewProc("Process32FirstW")
	procProcess32NextW                   = modkernel32.NewProc("Process32NextW")
	procCreateToolhelp32Snapshot         = modkernel32.NewProc("CreateToolhelp32Snapshot")
	procNtQuerySystemInformation         = modntdll.NewProc("NtQuerySystemInformation")
	procNtQueryInformationProcess        = modntdll.NewProc("NtQueryInformationProcess")
	procLookupPrivilegeNameW             = modadvapi32.NewProc("LookupPrivilegeNameW")
	procLookupPrivilegeValueW            = modadvapi32.NewProc("LookupPrivilegeValueW")
	procAdjustTokenPrivileges            = modadvapi32.NewProc("AdjustTokenPrivileges")
	procFindFirstVolumeW                 = modkernel32.NewProc("FindFirstVolumeW")
	procFindNextVolumeW                  = modkernel32.NewProc("FindNextVolumeW")
	procFindVolumeClose                  = modkernel32.NewProc("FindVolumeClose")
	procGetVolumePathNamesForVolumeNameW = modkernel32.NewProc("GetVolumePathNamesForVolumeNameW")
	procReadProcessMemory                = modkernel32.NewProc("ReadProcessMemory")
	procGetTickCount64                   = modkernel32.NewProc("GetTickCount64")
)

func _GlobalMemoryStatusEx(buffer *MemoryStatusEx) (err error) {
	r1, _, e1 := syscall.Syscall(procGlobalMemoryStatusEx.Addr(), 1, uintptr(unsafe.Pointer(buffer)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetLogicalDriveStringsW(bufferLength uint32, buffer *uint16) (length uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetLogicalDriveStringsW.Addr(), 2, uintptr(bufferLength), uintptr(unsafe.Pointer(buffer)), 0)
	length = uint32(r0)
	if length == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetProcessMemoryInfo(handle syscall.Handle, psmemCounters *ProcessMemoryCountersEx, cb uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procGetProcessMemoryInfo.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(psmemCounters)), uintptr(cb))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetProcessImageFileName(handle syscall.Handle, outImageFileName *uint16, size uint32) (length uint32, err error) {
	r0, _, e1 := syscall.Syscall(procGetProcessImageFileNameW.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(outImageFileName)), uintptr(size))
	length = uint32(r0)
	if length == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetSystemTimes(idleTime *syscall.Filetime, kernelTime *syscall.Filetime, userTime *syscall.Filetime) (err error) {
	r1, _, e1 := syscall.Syscall(procGetSystemTimes.Addr(), 3, uintptr(unsafe.Pointer(idleTime)), uintptr(unsafe.Pointer(kernelTime)), uintptr(unsafe.Pointer(userTime)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetDriveType(rootPathName *uint16) (dt DriveType, err error) {
	r0, _, e1 := syscall.Syscall(procGetDriveTypeW.Addr(), 1, uintptr(unsafe.Pointer(rootPathName)), 0, 0)
	dt = DriveType(r0)
	if dt == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _EnumProcesses(processIds *uint32, sizeBytes uint32, bytesReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumProcesses.Addr(), 3, uintptr(unsafe.Pointer(processIds)), uintptr(sizeBytes), uintptr(unsafe.Pointer(bytesReturned)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetDiskFreeSpaceEx(directoryName *uint16, freeBytesAvailable *uint64, totalNumberOfBytes *uint64, totalNumberOfFreeBytes *uint64) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetDiskFreeSpaceExW.Addr(), 4, uintptr(unsafe.Pointer(directoryName)), uintptr(unsafe.Pointer(freeBytesAvailable)), uintptr(unsafe.Pointer(totalNumberOfBytes)), uintptr(unsafe.Pointer(totalNumberOfFreeBytes)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _Process32First(handle syscall.Handle, processEntry32 *ProcessEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcess32FirstW.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(processEntry32)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _Process32Next(handle syscall.Handle, processEntry32 *ProcessEntry32) (err error) {
	r1, _, e1 := syscall.Syscall(procProcess32NextW.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(processEntry32)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _CreateToolhelp32Snapshot(flags uint32, processID uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateToolhelp32Snapshot.Addr(), 2, uintptr(flags), uintptr(processID), 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _NtQuerySystemInformation(systemInformationClass uint32, systemInformation *byte, systemInformationLength uint32, returnLength *uint32) (ntstatus uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procNtQuerySystemInformation.Addr(), 4, uintptr(systemInformationClass), uintptr(unsafe.Pointer(systemInformation)), uintptr(systemInformationLength), uintptr(unsafe.Pointer(returnLength)), 0, 0)
	ntstatus = uint32(r0)
	if ntstatus == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _NtQueryInformationProcess(processHandle syscall.Handle, processInformationClass uint32, processInformation *byte, processInformationLength uint32, returnLength *uint32) (ntstatus uint32, err error) {
	r0, _, e1 := syscall.Syscall6(procNtQueryInformationProcess.Addr(), 5, uintptr(processHandle), uintptr(processInformationClass), uintptr(unsafe.Pointer(processInformation)), uintptr(processInformationLength), uintptr(unsafe.Pointer(returnLength)), 0)
	ntstatus = uint32(r0)
	if ntstatus == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _LookupPrivilegeName(systemName string, luid *int64, buffer *uint16, size *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return
	}
	return __LookupPrivilegeName(_p0, luid, buffer, size)
}

func __LookupPrivilegeName(systemName *uint16, luid *int64, buffer *uint16, size *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procLookupPrivilegeNameW.Addr(), 4, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(luid)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _LookupPrivilegeValue(systemName string, name string, luid *int64) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return
	}
	var _p1 *uint16
	_p1, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return __LookupPrivilegeValue(_p0, _p1, luid)
}

func __LookupPrivilegeValue(systemName *uint16, name *uint16, luid *int64) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _AdjustTokenPrivileges(token syscall.Token, releaseAll bool, input *byte, outputSize uint32, output *byte, requiredSize *uint32) (success bool, err error) {
	var _p0 uint32
	if releaseAll {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(input)), uintptr(outputSize), uintptr(unsafe.Pointer(output)), uintptr(unsafe.Pointer(requiredSize)))
	success = r0 != 0
	if true {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _FindFirstVolume(volumeName *uint16, size uint32) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procFindFirstVolumeW.Addr(), 2, uintptr(unsafe.Pointer(volumeName)), uintptr(size), 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _FindNextVolume(handle syscall.Handle, volumeName *uint16, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procFindNextVolumeW.Addr(), 3, uintptr(handle), uintptr(unsafe.Pointer(volumeName)), uintptr(size))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _FindVolumeClose(handle syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFindVolumeClose.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetVolumePathNamesForVolumeName(volumeName string, buffer *uint16, bufferSize uint32, length *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(volumeName)
	if err != nil {
		return
	}
	return __GetVolumePathNamesForVolumeName(_p0, buffer, bufferSize, length)
}

func __GetVolumePathNamesForVolumeName(volumeName *uint16, buffer *uint16, bufferSize uint32, length *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetVolumePathNamesForVolumeNameW.Addr(), 4, uintptr(unsafe.Pointer(volumeName)), uintptr(unsafe.Pointer(buffer)), uintptr(bufferSize), uintptr(unsafe.Pointer(length)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _ReadProcessMemory(handle syscall.Handle, baseAddress uintptr, buffer uintptr, size uintptr, numRead *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procReadProcessMemory.Addr(), 5, uintptr(handle), uintptr(baseAddress), uintptr(buffer), uintptr(size), uintptr(unsafe.Pointer(numRead)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetTickCount64() (uptime uint64, err error) {
	r0, _, e1 := syscall.Syscall(procGetTickCount64.Addr(), 0, 0, 0, 0)
	uptime = uint64(r0)
	if uptime == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
