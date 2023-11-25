package main

import (
	"fmt"
	"strings"
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

func (p *page) print() {
	fmt.Printf("[page: %d]\n", p.number)
	for _, line := range p.lines {
		fmt.Printf("%s\n", line)
	}
}

func readPages(s string) *pages {
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
			ww := len(s)
			if w < ww {
				w = ww
			}
			ls = append(ls, s)
		}
	}

	return &pages{ps: ps, width: w, height: h}
}
