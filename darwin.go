// +build darwin
// +build 386 amd64

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework OpenGL -framework IOKit
#include <OpenGL/gl3.h>
#import <Carbon/Carbon.h> // for HIToolbox/Events.h
#import <Cocoa/Cocoa.h>
#include <pthread.h>
#include <stdint.h>
#include <stdlib.h>
void clipClear();
void pasteWriteAddText(char* data, int dlen);
void clipWrite();
*/
import "C"
import "unsafe"

func Clear() {
	C.clipClear()
}

func WriteText(b []byte) {
	Clear()
	sz := len(b)
	cdata := C.malloc(C.size_t(sz))
	copy((*[1 << 30]byte)(cdata)[0:sz], b)
	C.pasteWriteAddText((*C.char)(cdata), C.int(sz))
	C.free(unsafe.Pointer(cdata))
	C.clipWrite()
}
