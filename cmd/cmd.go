package main

import (
	"os"

	"github.com/takutakahashi/pagekite-ingress-controller/pkg/pagekite"
)

func main() {
	pk := pagekite.NewPageKite()
	go pk.StartObserver()
	<-pk.Stop
	os.Exit(0)
}
