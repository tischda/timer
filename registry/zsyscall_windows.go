//go:build windows

package registry

import (
	"syscall"
	"unsafe"
)

// Code Copyright 2015 The Go Authors extracted from:
// https://github.com/golang/sys/blob/master/windows/registry/zsyscall_windows.go

var (
	// Advanced Services (advapi32.dll) provide access to the Windows registry
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procRegCreateKeyExW = modadvapi32.NewProc("RegCreateKeyExW")
	procRegDeleteKeyW   = modadvapi32.NewProc("RegDeleteKeyW")
	procRegSetValueExW  = modadvapi32.NewProc("RegSetValueExW")
	procRegEnumValueW   = modadvapi32.NewProc("RegEnumValueW")
	procRegDeleteValueW = modadvapi32.NewProc("RegDeleteValueW")
)

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724844(v=vs.85).aspx
func regCreateKeyEx(key syscall.Handle, subkey *uint16, reserved uint32, class *uint16, options uint32, desired uint32, sa *syscall.SecurityAttributes, result *syscall.Handle, disposition *uint32) (regerrno error) {
	ret, _, _ := procRegCreateKeyExW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(subkey)),
		uintptr(reserved),
		uintptr(unsafe.Pointer(class)),
		uintptr(options),
		uintptr(desired),
		uintptr(unsafe.Pointer(sa)),
		uintptr(unsafe.Pointer(result)),
		uintptr(unsafe.Pointer(disposition)))
	if ret != 0 {
		regerrno = syscall.Errno(ret)
	}
	return
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724845(v=vs.85).aspx
func regDeleteKey(key syscall.Handle, subkey *uint16) (regerrno error) {
	ret, _, _ := procRegDeleteKeyW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(subkey)))
	if ret != 0 {
		regerrno = syscall.Errno(ret)
	}
	return
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724923(v=vs.85).aspx
func regSetValueEx(hKey syscall.Handle, lpValueName *uint16, Reserved uint32, dwType uint32, lpData *byte, cbData uint32) (regerrno error) {
	ret, _, _ := procRegSetValueExW.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpValueName)),
		uintptr(Reserved),
		uintptr(dwType),
		uintptr(unsafe.Pointer(lpData)),
		uintptr(cbData))

	// If the function fails, the return value is a nonzero error code defined in Winerror.h
	if ret != 0 {
		regerrno = syscall.Errno(ret)
	}
	return
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724851(v=vs.85).aspx
func regDeleteValue(hKey syscall.Handle, lpValueName *uint16) (regerrno error) {
	ret, _, _ := procRegDeleteValueW.Call(
		uintptr(hKey),
		uintptr(unsafe.Pointer(lpValueName)))

	// If the function fails, the return value is a nonzero error code defined in Winerror.h
	if ret != 0 {
		regerrno = syscall.Errno(ret)
	}
	return
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724865(v=vs.85).aspx
func regEnumValue(key syscall.Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error) {
	ret, _, _ := procRegEnumValueW.Call(
		uintptr(key),
		uintptr(index),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(nameLen)),
		uintptr(unsafe.Pointer(reserved)),
		uintptr(unsafe.Pointer(valtype)),
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(buflen)))
	if ret != 0 {
		regerrno = syscall.Errno(ret)
	}
	return
}
