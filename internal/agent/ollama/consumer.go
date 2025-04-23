package ollama

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/vinit-chauhan/devmind/cmd/ui"
)

func Consume(ctx context.Context, tknCh <-chan string, doneCh chan<- struct{}, full *strings.Builder) {

	spinner := ctx.Value("spinner").(*ui.Spinner)

	for tok := range tknCh {

		if spinner != nil {
			spinner.Stop()
			spinner = nil
		}

		fmt.Fprint(os.Stdout, tok)
		full.WriteString(tok)
	}

	fmt.Fprintln(os.Stdout, "")
	doneCh <- struct{}{}
}
