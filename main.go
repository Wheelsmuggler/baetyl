package main

import (
	_ "github.com/baetyl/baetyl/ami/kube"
	_ "github.com/baetyl/baetyl/ami/native"
	"github.com/baetyl/baetyl/cmd"
	_ "github.com/baetyl/baetyl/plugin/httplink"
)

func main() {
	cmd.Execute()
}
