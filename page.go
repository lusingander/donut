package main

import (
	"errors"
	"strings"

	"github.com/mattn/go-runewidth"
)

const (
	separator = "────────────────────────────────────────────────────────────────────────────────"
)

type pages struct {
	ps     []*page
	width  int
	height int
}

type page struct {
	number int
	lines  []string
}

func readPages(s string) (*pages, error) {
	ss := strings.Split(s, "\n")
	ps := make([]*page, 0)

	w := 0
	h := 0
	n := 1
	ls := make([]string, 0)
	for _, s := range ss {
		if s == separator {
			p := &page{
				number: n,
				lines:  ls,
			}
			ps = append(ps, p)

			hh := len(ls)
			if h < hh {
				h = hh
			}
			n += 1
			ls = make([]string, 0)
		} else {
			ww := runewidth.StringWidth(s)
			if w < ww {
				w = ww
			}
			ls = append(ls, s)
		}
	}

	if len(ps) == 0 {
		return nil, errors.New("invalid input")
	}

	return &pages{ps: ps, width: w, height: h}, nil
}
