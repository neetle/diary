#ifndef CLIBRARY_H
#define CLIBRARY_H
#include <Carbon/Carbon.h>
OSStatus InstallEv(EventTargetRef ref, EventTypeSpec eventType, int id);
OSStatus SetupHotkey(EventTargetRef ref, int id, UInt32 key, UInt32 modifiers);
#endif