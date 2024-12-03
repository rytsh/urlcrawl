package main

import (
	"rytsh/urlcrawl/cmd/urlcrawl/args"

	"github.com/rakunlabs/into"
	"github.com/rakunlabs/logi"
)

var (
	version = "v0.0.0"
	commit  = "?"
	date    = ""
)

func main() {
	args.BuildVars.Version = version
	args.BuildVars.Date = date
	args.BuildVars.Commit = commit

	into.Init(
		args.Execute,
		into.WithLogger(logi.InitializeLog(logi.WithCaller(false))),
		into.WithMsgf("urlcrawl [%s]", version),
		into.WithStartFn(nil),
		into.WithStopFn(nil),
	)
}
