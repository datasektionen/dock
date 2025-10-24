package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/datasektionen/dock/pkg/config"
	"github.com/datasektionen/dock/pkg/dao"
	"github.com/datasektionen/dock/pkg/rfinger"
	// "github.com/datasektionen/dock/pkg/spam"
	"github.com/datasektionen/dock/pkg/ston"
	"golang.org/x/term"
)

func main() {

	cfg := config.GetConfig()

	dao := dao.New(cfg)

	go rfinger.Listen(cfg, dao)
	go ston.Listen(cfg, dao)

	if term.IsTerminal(int(os.Stdin.Fd())) {
		stdin := bufio.NewReader(os.Stdin)
		for {
			_, err := stdin.ReadString('\n')
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				panic(err)
			}
		}
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		wg.Wait()
	}
}
