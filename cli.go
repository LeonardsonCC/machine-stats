package main

import (
	"errors"
	"flag"
)

type Args struct {
	Memory bool
	Cpu    bool
}

func ParseArgs() (*Args, error) {
  memory := flag.Bool("m", false, "should show memory usage")
  cpu := flag.Bool("c", false, "should show cpu usage")
  flag.Parse()

	args := &Args{
    Memory: *memory,
		Cpu: *cpu,
  }

	if !args.Memory && !args.Cpu {
		return nil, errors.New("Nothing to show")
	}

	return args, nil
}
