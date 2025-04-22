package ui

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Spinner struct {
	chars []rune
	stop  chan struct{}
	done  chan struct{}
	msg   string

	ctx context.Context
}

func NewSpinner(ctx context.Context) *Spinner {
	return &Spinner{
		chars: []rune{'|', '/', '-', '\\'},
		stop:  make(chan struct{}),
		done:  make(chan struct{}),
		ctx:   ctx,
	}
}

func (s *Spinner) Start(msg string) {
	s.msg = msg
	go func() {
		i := 0
		for {
			select {
			case <-s.stop:
				s.clearLine()
				s.done <- struct{}{}
				return
			case <-s.ctx.Done():
				s.clearLine()
				s.done <- struct{}{}
				return
			default:
				fmt.Printf("\r%c %s", s.chars[i%len(s.chars)], msg)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	s.stop <- struct{}{}
	<-s.done
}

func (s *Spinner) clearLine() {
	clearLen := len(s.msg) + 4 // char + space + msg
	fmt.Print("\r" + strings.Repeat(" ", clearLen) + "\r")
}
