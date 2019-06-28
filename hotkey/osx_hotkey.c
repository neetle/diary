
#include <Carbon/Carbon.h>

extern void run_callback(int id);

OSStatus OnHotKeyEvent(EventHandlerCallRef nextHandler, EventRef theEvent, void * data) {
	int x = data;
	run_callback(x);
	return noErr;
}

OSStatus InstallEv(EventTargetRef ref, EventTypeSpec eventType, int id) {
    void * idAddr = &id;
    return InstallEventHandler(ref, &OnHotKeyEvent, 1, &eventType, id, NULL);
}

OSStatus SetupHotkey(EventTargetRef ref, int id, UInt32 key, UInt32 modifiers) {
    // run_callback(id);
	EventHotKeyRef gMyHotKeyRef;

    EventHotKeyID gMyHotKeyID;
    gMyHotKeyID.signature=typeEventHotKeyID;
    // gMyHotKeyID.signature='htk1';
    gMyHotKeyID.id=id;

    return RegisterEventHotKey(
        kVK_Space, 
        optionKey, 
        gMyHotKeyID, 
        ref, 
        kEventHotKeyExclusive, 
        &gMyHotKeyRef
    );  
}