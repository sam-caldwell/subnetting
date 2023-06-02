package main

import (
	"fmt"
	"github.com/sam-caldwell/subnetting/v2/subnetting"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(subnetting.ErrMissingArguments)
		os.Exit(subnetting.ExitMissingArgs)
	}
	parentCIDR := os.Args[1]
	subnetSize := func() int {
		var err error
		var n int64
		s := os.Args[2]
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			fmt.Println(err)
			os.Exit(subnetting.ExitSubnettingError)
		}
		return int(n)
	}()
	if subnets, err := subnetting.CalculateSubnets(parentCIDR, subnetSize); err != nil {
		fmt.Printf(subnetting.ErrGeneral, err)
	} else {
		for _, network := range subnets {
			fmt.Printf("%s", network)
		}
	}
	os.Exit(subnetting.ExitSuccess)
}
