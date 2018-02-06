package main

import(
	"fmt"
	"strings"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func main() {

}

func run(){

	var chainparams = &chaincfg.TestNet3Params
	s = "a"
	cp2Addr, err := btcutil.DecodeAddress(s,chainparams)
	