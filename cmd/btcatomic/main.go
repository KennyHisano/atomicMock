package main

import(
	"fmt"
	"strings"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"strconv"
	"net/rpc"
)

func main() {

}

type command interface {
	runcommand(*rpc.Client) error
}

type initiateCmd struct {
	cp2Addr *btcutil.AddressPubKeyHash
	amount  btcutil.Amount
}



func run() (err error) {

	var cmd command
	//get address of testnet

	var s = "senderaddresshere"
	amountVal := float64(1.3)
	var chainparams = &chaincfg.TestNet3Params
	cp2Addr, err := btcutil.DecodeAddress(s,chainparams)

	amount, err := btcutil.NewAmount(amountVal)
	if err != nil {
			return fmt.Errorf("amount wrong")
	}

	cp2AddrP2PKH := cp2Addr.(*btcutil.AddressPubKeyHash)

	cmd = &initiateCmd{cp2Addr: cp2AddrP2PKH,amount: amount}

}
