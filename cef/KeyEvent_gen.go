// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// KeyEvent (cef_key_event_t from include/internal/cef_types.h)
// Structure representing keyboard event information.
type KeyEvent struct {
	// Type (_type)
	// The type of keyboard event.
	Type KeyEventType
	// Modifiers (modifiers)
	// Bit flags describing any pressed modifier keys. See
	// cef_event_flags_t for values.
	Modifiers uint32
	// WindowsKeyCode (windows_key_code)
	// The Windows key code for the key event. This value is used by the DOM
	// specification. Sometimes it comes directly from the event (i.e. on
	// Windows) and sometimes it's determined using a mapping function. See
	// WebCore/platform/chromium/KeyboardCodes.h for the list of values.
	WindowsKeyCode int32
	// NativeKeyCode (native_key_code)
	// The actual key code genenerated by the platform.
	NativeKeyCode int32
	// IsSystemKey (is_system_key)
	// Indicates whether the event is considered a "system key" event (see
	// http://msdn.microsoft.com/en-us/library/ms646286(VS.85).aspx for details).
	// This value will always be false on non-Windows platforms.
	IsSystemKey int32
	// Character (character)
	// The character generated by the keystroke.
	Character int16
	// UnmodifiedCharacter (unmodified_character)
	// Same as |character| but unmodified by any concurrently-held modifiers
	// (except shift). This is useful for working out shortcut keys.
	UnmodifiedCharacter int16
	// FocusOnEditableField (focus_on_editable_field)
	// True if the focus is currently on an editable field on the page. This is
	// useful for determining if standard key events should be intercepted.
	FocusOnEditableField int32
}

// NewKeyEvent creates a new KeyEvent.
func NewKeyEvent() *KeyEvent {
	return &KeyEvent{}
}

func (d *KeyEvent) toNative(native *C.cef_key_event_t) *C.cef_key_event_t {
	if d == nil {
		return nil
	}
	native._type = C.cef_key_event_type_t(d.Type)
	native.modifiers = C.uint32(d.Modifiers)
	native.windows_key_code = C.int(d.WindowsKeyCode)
	native.native_key_code = C.int(d.NativeKeyCode)
	native.is_system_key = C.int(d.IsSystemKey)
	native.character = C.char16(d.Character)
	native.unmodified_character = C.char16(d.UnmodifiedCharacter)
	native.focus_on_editable_field = C.int(d.FocusOnEditableField)
	return native
}

func (n *C.cef_key_event_t) toGo() *KeyEvent {
	if n == nil {
		return nil
	}
	var d KeyEvent
	n.intoGo(&d)
	return &d
}

func (n *C.cef_key_event_t) intoGo(d *KeyEvent) {
	d.Type = KeyEventType(n._type)
	d.Modifiers = uint32(n.modifiers)
	d.WindowsKeyCode = int32(n.windows_key_code)
	d.NativeKeyCode = int32(n.native_key_code)
	d.IsSystemKey = int32(n.is_system_key)
	d.Character = int16(n.character)
	d.UnmodifiedCharacter = int16(n.unmodified_character)
	d.FocusOnEditableField = int32(n.focus_on_editable_field)
}
