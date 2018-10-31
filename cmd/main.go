package main

import "goblockchain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 2 BTC to Jacky")
	bc.Print()
}
