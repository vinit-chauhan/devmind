package consumer

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/vinit-chauhan/devmind/cmd/ui"
	"github.com/vinit-chauhan/devmind/internal/logger"
)

func Consume(ctx context.Context, tknCh <-chan string, doneCh chan<- struct{}, full *strings.Builder) {
	spinner, _ := ctx.Value("spinner").(*ui.Spinner)

	defer func() {
		logger.Debug("Closing the consumer")
		if spinner != nil {
			spinner.Stop()
		}
		fmt.Fprintln(os.Stdout) // ensure newline
		doneCh <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Fprint(os.Stdout, "\nRequest cancelled")
			return

		case tok, ok := <-tknCh:
			if !ok {
				logger.Debug("Channel closed")
				return
			}
			if spinner != nil {
				spinner.Stop()
				spinner = nil
			}
			fmt.Fprint(os.Stdout, tok)
			full.WriteString(tok)
		}
	}
}
