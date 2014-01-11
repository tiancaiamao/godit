package main

import (
	"github.com/nsf/termbox-go"
)

type jump_inner_line_mode struct {
	stub_overlay_mode
	godit   *godit
	forward bool
}

func init_jump_inner_line_mode(godit *godit, _forward bool) *jump_inner_line_mode {
	k := new(jump_inner_line_mode)
	k.godit = godit
	k.forward = _forward
	if _forward {
		k.godit.set_status("->")
	} else {
		k.godit.set_status("<-")
	}
	return k
}

func (k *jump_inner_line_mode) on_key(ev *termbox.Event) {
	if ev.Mod != 0 {
		return
	}

	v := k.godit.views.leaf
	if ev.Ch != 0 {
		if k.forward {
			v.on_vcommand(vcommand_move_cursor_to_rune_forward, ev.Ch)
		} else {
			v.on_vcommand(vcommand_move_cursor_to_rune_backward, ev.Ch)
		}
	}
	k.godit.set_overlay_mode(nil)
}
