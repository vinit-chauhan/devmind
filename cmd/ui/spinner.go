package ui

import (
	"context"
	"fmt"
	"time"
)

type Spinner struct {
	chars []rune
	stop  chan struct{}
	ctx   context.Context
}

func NewSpinner(ctx context.Context) *Spinner {
	return &Spinner{
		chars: []rune{'|', '/', '-', '\\'},
		stop:  make(chan struct{}),
		ctx:   ctx,
	}
}

func (s *Spinner) Start(msg string) {
	go func() {
		i := 0
		for {
			select {
			case <-s.stop:
				fmt.Print('\r')
			case <-s.ctx.Done():
				fmt.Print('\r')
			default:
				fmt.Printf("\r%c %s", s.chars[i%len(s.chars)], msg)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Stop(msg string) {
	s.stop <- struct{}{}
	fmt.Println(msg)
}
