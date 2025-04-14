package argumentParser;

import (
	"os"
	"strconv"
	"errors"
)

type ArgumentDef struct {
	RelativeLocation string;
	MaxCpuUsage int;
	MaxMemoryUsage int; 
	Cmd string;
}

func ParseArguments() (*ArgumentDef, error){
	arguments := os.Args;
	if(len(arguments)<5){
		return nil, errors.New("Incorrect number of arguments")
	}
	location := arguments[1];
	cpu, errCpu := strconv.Atoi(arguments[2]);
	memory, errMem := strconv.Atoi(arguments[3]);
	cmd := arguments[4];

	if(errCpu != nil){
		return nil, errCpu
	}
	if(errMem != nil){
		return nil, errMem
	}
	args := ArgumentDef{location, cpu, memory, cmd};
	return &args, nil;
}
