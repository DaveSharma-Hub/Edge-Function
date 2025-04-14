package main

import (
	"fmt"
	"github.com/DaveSharma-Hub/Edge-Function/src/edge-function/argumentParser"
	"github.com/DaveSharma-Hub/Edge-Function/src/edge-function/server"
)

func main(){
	args, err := argumentParser.ParseArguments();
	if (err != nil) {
		fmt.Println(err);
		return;
	}

	fmt.Println("Main: ", args.RelativeLocation, args.MaxCpuUsage, args.MaxMemoryUsage, args.Cmd);
	server.Start();
}