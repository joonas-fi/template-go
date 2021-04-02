package main

import (
	"context"

	"github.com/function61/gokit/os/osutil"
)

func main() {
	osutil.ExitIfError(
		logic(osutil.CancelOnInterruptOrTerminate(nil)))
}

func logic(ctx context.Context)error{
	return nil
}

