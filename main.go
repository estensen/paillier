package main

import (
	"fmt"
	"math/big"

	"github.com/getamis/alice/crypto/homo/paillier"
)

func main() {
	p, _ := paillier.NewPaillier(2048)
	m1 := big.NewInt(100)
	m2 := big.NewInt(200)
	c1, _ := p.Encrypt(m1.Bytes())
	c2, _ := p.Encrypt(m2.Bytes())
	encryptedSum, _ := p.PublicKey.Add(c1, c2)
	decryptedSum, _ := p.Decrypt(encryptedSum)
	fmt.Println("The decryption of adding m1 and m2 to be", new(big.Int).SetBytes(decryptedSum))
	// The decryption of adding m1 and m2 to be 300
}
