// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/krigga/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

//go:build darwin || windows
// +build darwin windows

package robotgo

// GetBounds get the window bounds
func GetBounds(pid int32, args ...int) (int, int, int, int) {
	var hwnd int
	if len(args) > 0 {
		hwnd = args[0]
	}

	return internalGetBounds(pid, hwnd)
}

// internalGetTitle get the window title
func internalGetTitle(pid int32, args ...int32) string {
	var isHwnd int32
	if len(args) > 0 {
		isHwnd = args[0]
	}
	gtitle := cgetTitle(pid, isHwnd)

	return gtitle
}

// ActivePID active the window by PID,
//
// If args[0] > 0 on the Windows platform via a window handle to active
//
// Examples:
//	ids, _ := robotgo.FindIds()
//	robotgo.ActivePID(ids[0])
func ActivePID(pid int32, args ...int) error {
	var hwnd int
	if len(args) > 0 {
		hwnd = args[0]
	}

	internalActive(pid, hwnd)
	return nil
}
