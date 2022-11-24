// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 The Ebitengine Authors

//go:build darwin || linux

package fakecgo

import (
	"syscall"
	"unsafe"
)

// setg_trampoline calls setg with the G provided
func setg_trampoline(setg uintptr, G uintptr)

//go:linkname memmove runtime.memmove
func memmove(to, from unsafe.Pointer, n uintptr)

// call5 takes fn the C function and 5 arguments and calls the function with those arguments
func call5(fn, a1, a2, a3, a4, a5 uintptr) uintptr

func malloc(size uintptr) unsafe.Pointer {
	ret := call5(mallocABI0, size, 0, 0, 0, 0)
	// this indirection is to avoid go vet complaining about possible misuse of unsafe.Pointer
	return *(*unsafe.Pointer)(unsafe.Pointer(&ret))
}

func free(ptr unsafe.Pointer) {
	call5(freeABI0, uintptr(ptr), 0, 0, 0, 0)
}

func setenv(name, value *byte, overwrite int32) int32 {
	return int32(call5(setenvABI0, uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value)), uintptr(overwrite), 0, 0))
}

func unsetenv(name *byte) int32 {
	return int32(call5(unsetenvABI0, uintptr(unsafe.Pointer(name)), 0, 0, 0, 0))
}

func pthread_attr_init(attr *pthread_attr_t) int32 {
	return int32(call5(pthread_attr_initABI0, uintptr(unsafe.Pointer(attr)), 0, 0, 0, 0))
}

func pthread_create(thread *pthread_t, attr *pthread_attr_t, start, arg unsafe.Pointer) int32 {
	return int32(call5(pthread_createABI0, uintptr(unsafe.Pointer(thread)), uintptr(unsafe.Pointer(attr)), uintptr(start), uintptr(arg), 0))
}

func pthread_detach(thread pthread_t) int32 {
	return int32(call5(pthread_detachABI0, uintptr(thread), 0, 0, 0, 0))
}

func pthread_sigmask(how sighow, ign, oset *sigset_t) int32 {
	return int32(call5(pthread_sigmaskABI0, uintptr(how), uintptr(unsafe.Pointer(ign)), uintptr(unsafe.Pointer(oset)), 0, 0))
}

func pthread_attr_getstacksize(attr *pthread_attr_t, stacksize *size_t) int32 {
	return int32(call5(pthread_attr_getstacksizeABI0, uintptr(unsafe.Pointer(attr)), uintptr(unsafe.Pointer(stacksize)), 0, 0, 0))
}

func pthread_attr_destroy(attr *pthread_attr_t) int32 {
	return int32(call5(pthread_attr_destroyABI0, uintptr(unsafe.Pointer(attr)), 0, 0, 0, 0))
}

func sigfillset(set *sigset_t) int32 {
	return int32(call5(sigfillsetABI0, uintptr(unsafe.Pointer(set)), 0, 0, 0, 0))
}

func nanosleep(ts, rem *syscall.Timespec) int32 {
	return int32(call5(nanosleepABI0, uintptr(unsafe.Pointer(ts)), uintptr(unsafe.Pointer(rem)), 0, 0, 0))
}

func abort() {
	call5(abortABI0, 0, 0, 0, 0, 0)
}

//go:linkname _malloc _malloc
var _malloc uintptr
var mallocABI0 = uintptr(unsafe.Pointer(&_malloc))

//go:linkname _free _free
var _free uintptr
var freeABI0 = uintptr(unsafe.Pointer(&_free))

//go:linkname _setenv _setenv
var _setenv uintptr
var setenvABI0 = uintptr(unsafe.Pointer(&_setenv))

//go:linkname _unsetenv _unsetenv
var _unsetenv uintptr
var unsetenvABI0 = uintptr(unsafe.Pointer(&_unsetenv))

//go:linkname _pthread_attr_init _pthread_attr_init
var _pthread_attr_init uintptr
var pthread_attr_initABI0 = uintptr(unsafe.Pointer(&_pthread_attr_init))

//go:linkname _pthread_create _pthread_create
var _pthread_create uintptr
var pthread_createABI0 = uintptr(unsafe.Pointer(&_pthread_create))

//go:linkname _pthread_detach _pthread_detach
var _pthread_detach uintptr
var pthread_detachABI0 = uintptr(unsafe.Pointer(&_pthread_detach))

//go:linkname _pthread_sigmask _pthread_sigmask
var _pthread_sigmask uintptr
var pthread_sigmaskABI0 = uintptr(unsafe.Pointer(&_pthread_sigmask))

//go:linkname _pthread_attr_getstacksize _pthread_attr_getstacksize
var _pthread_attr_getstacksize uintptr
var pthread_attr_getstacksizeABI0 = uintptr(unsafe.Pointer(&_pthread_attr_getstacksize))

//go:linkname _pthread_attr_destroy _pthread_attr_destroy
var _pthread_attr_destroy uintptr
var pthread_attr_destroyABI0 = uintptr(unsafe.Pointer(&_pthread_attr_destroy))

//go:linkname _sigfillset _sigfillset
var _sigfillset uintptr
var sigfillsetABI0 = uintptr(unsafe.Pointer(&_sigfillset))

//go:linkname _nanosleep _nanosleep
var _nanosleep uintptr
var nanosleepABI0 = uintptr(unsafe.Pointer(&_nanosleep))

//go:linkname _abort _abort
var _abort uintptr
var abortABI0 = uintptr(unsafe.Pointer(&_abort))
