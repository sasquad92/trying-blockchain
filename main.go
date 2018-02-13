package main

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BC to Kevin")
	bc.AddBlock("Send 1 BC to Eva")

	// for _, block := range bc.blocks {
	// 	pow := NewProofOfWork(block)

	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }
}
