//go:build !no_gotd_slices
// +build !no_gotd_slices

// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// InputQuickReplyShortcutClassArray is adapter for slice of InputQuickReplyShortcutClass.
type InputQuickReplyShortcutClassArray []InputQuickReplyShortcutClass

// Sort sorts slice of InputQuickReplyShortcutClass.
func (s InputQuickReplyShortcutClassArray) Sort(less func(a, b InputQuickReplyShortcutClass) bool) InputQuickReplyShortcutClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputQuickReplyShortcutClass.
func (s InputQuickReplyShortcutClassArray) SortStable(less func(a, b InputQuickReplyShortcutClass) bool) InputQuickReplyShortcutClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputQuickReplyShortcutClass.
func (s InputQuickReplyShortcutClassArray) Retain(keep func(x InputQuickReplyShortcutClass) bool) InputQuickReplyShortcutClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputQuickReplyShortcutClassArray) First() (v InputQuickReplyShortcutClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputQuickReplyShortcutClassArray) Last() (v InputQuickReplyShortcutClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputQuickReplyShortcutClassArray) PopFirst() (v InputQuickReplyShortcutClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputQuickReplyShortcutClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputQuickReplyShortcutClassArray) Pop() (v InputQuickReplyShortcutClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputQuickReplyShortcut returns copy with only InputQuickReplyShortcut constructors.
func (s InputQuickReplyShortcutClassArray) AsInputQuickReplyShortcut() (to InputQuickReplyShortcutArray) {
	for _, elem := range s {
		value, ok := elem.(*InputQuickReplyShortcut)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputQuickReplyShortcutID returns copy with only InputQuickReplyShortcutID constructors.
func (s InputQuickReplyShortcutClassArray) AsInputQuickReplyShortcutID() (to InputQuickReplyShortcutIDArray) {
	for _, elem := range s {
		value, ok := elem.(*InputQuickReplyShortcutID)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputQuickReplyShortcutArray is adapter for slice of InputQuickReplyShortcut.
type InputQuickReplyShortcutArray []InputQuickReplyShortcut

// Sort sorts slice of InputQuickReplyShortcut.
func (s InputQuickReplyShortcutArray) Sort(less func(a, b InputQuickReplyShortcut) bool) InputQuickReplyShortcutArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputQuickReplyShortcut.
func (s InputQuickReplyShortcutArray) SortStable(less func(a, b InputQuickReplyShortcut) bool) InputQuickReplyShortcutArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputQuickReplyShortcut.
func (s InputQuickReplyShortcutArray) Retain(keep func(x InputQuickReplyShortcut) bool) InputQuickReplyShortcutArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputQuickReplyShortcutArray) First() (v InputQuickReplyShortcut, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputQuickReplyShortcutArray) Last() (v InputQuickReplyShortcut, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputQuickReplyShortcutArray) PopFirst() (v InputQuickReplyShortcut, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputQuickReplyShortcut
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputQuickReplyShortcutArray) Pop() (v InputQuickReplyShortcut, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputQuickReplyShortcutIDArray is adapter for slice of InputQuickReplyShortcutID.
type InputQuickReplyShortcutIDArray []InputQuickReplyShortcutID

// Sort sorts slice of InputQuickReplyShortcutID.
func (s InputQuickReplyShortcutIDArray) Sort(less func(a, b InputQuickReplyShortcutID) bool) InputQuickReplyShortcutIDArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputQuickReplyShortcutID.
func (s InputQuickReplyShortcutIDArray) SortStable(less func(a, b InputQuickReplyShortcutID) bool) InputQuickReplyShortcutIDArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputQuickReplyShortcutID.
func (s InputQuickReplyShortcutIDArray) Retain(keep func(x InputQuickReplyShortcutID) bool) InputQuickReplyShortcutIDArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s InputQuickReplyShortcutIDArray) First() (v InputQuickReplyShortcutID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputQuickReplyShortcutIDArray) Last() (v InputQuickReplyShortcutID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputQuickReplyShortcutIDArray) PopFirst() (v InputQuickReplyShortcutID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputQuickReplyShortcutID
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputQuickReplyShortcutIDArray) Pop() (v InputQuickReplyShortcutID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}