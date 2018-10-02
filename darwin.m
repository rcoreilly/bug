// +build darwin
// +build 386 amd64
// +build !ios

#include "_cgo_export.h"
#include <pthread.h>
#include <stdio.h>

#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>
#import <OpenGL/gl3.h>
#import <IOKit/graphics/IOGraphicsLib.h>

static NSMutableArray *pasteWriteItems = NULL;

// add text to the list of items to paste
void pasteWriteAddText(char* data, int len) {
    NSString *ns_clip;
    bool ret;

    if(pasteWriteItems == NULL) {
        pasteWriteItems = [NSMutableArray array];
    }
    
    ns_clip = [[NSString alloc] initWithBytes:data length:len encoding:NSUTF8StringEncoding];
    [pasteWriteItems addObject:ns_clip];
}	

void pasteWrite(NSPasteboard* pb) {
    if(pasteWriteItems == NULL) {
        return;
    }
    [pb writeObjects: pasteWriteItems];
    [pasteWriteItems removeAllObjects];
}	

void clipClear() {
    NSPasteboard *pb = [NSPasteboard generalPasteboard];
    if(pb == NULL) {
        return;
    }
    [pb clearContents];
    if(pasteWriteItems != NULL) {
        [pasteWriteItems removeAllObjects];
    }
}

void clipWrite() {
    NSPasteboard *pb = [NSPasteboard generalPasteboard];
    if(pb == NULL) {
        return;
    }
    pasteWrite(pb);
}

