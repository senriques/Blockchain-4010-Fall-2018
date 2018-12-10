// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 9 Zero Knowledge Proof

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)
// From: https://arrow.dit.ie/cgi/viewcontent.cgi?article=1031&context=itbj
// Id-in-crypto.pdf -- See p. 38, 40, 41 ((By the page numbers in the documetn))
// p40 (page number) is page 30 in my PDF viewer.  Look for section 7.9.

func main() {
	// From p40
	p     := big.NewInt(88667)
	q     := big.NewInt(1031)
	alpha := big.NewInt(70322)
	a     := big.NewInt(755)
	t1    := big.NewInt(0)
	v     := big.NewInt(0)
	r     := big.NewInt(543) 					// Should be random, but for this example
	x     := big.NewInt(0)
	e     := big.NewInt(1000)					// Bob's challenge e-value
	t2     := big.NewInt(0)
	t3     := big.NewInt(0)
	t4     := big.NewInt(0)
	z      := big.NewInt(0)						// Bob's verification value

	t1.Sub(q, a)
	v.Exp(alpha, t1, p)								// v = (alpha ^ ( q-a )) % p
	fmt.Printf("Setup Complete :       v= %s\n", v)

	// Alice is Client:  Msg 1 - Client to Server, Alice Chooses, and send to Bob
	M := big.NewInt(9999) 					// Generate crypto random value
	r, err := rand.Int(rand.Reader, M)
	if err != nil {
		fmt.Printf("Failed to generate random value, %s\n", err)
		os.Exit(1)
	}
	x.Exp(alpha, r, p)	 							// x=(alpha^r) % p
	fmt.Printf("Send To Bob :          x= %s\n", x)
	// Response to Message 1, Server back to client
	// Bob is the Server:
	// Bob sends the challenge 'e' back to Alice e to do the computation
	fmt.Printf("Bob Challenges Alice : e= %s\n", e)
	y := big.NewInt(0)								// Alice now computes: y = a*e % q
	y.Mul(a, e)
	y.Mod(y, q)
	y.Add(y, r)
	// fmt.Printf("a= %s\nr= %s\nq= %s\n", a, r, q)  		// debug
	fmt.Printf("Alice responds :       y= %s\n", y) // Prints 851

	// Message 2 - Client (Alice) with response to challenge.
	// Bob (server) verifies: x == z == (a^y) * (v^e) % p
	// Bob (server) verifies: x == z == ((a^y)%p)*((v^e)%p) % p
  t2.Exp(alpha, y, p)
	t3.Exp(v, e, p)
	t4.Mul(t2, t3)
	z.Mod(t4, p)
	//	fmt.Printf("x= %s\nalpha= %s\ny= %s\nv= %s\ne= %s\np= %s\n", x, alpha, y, v, e, p) 	// debug
	//	fmt.Printf("t2= %s\nt3= %s\nt4= %s\nz= %s\n", t2, t3, t4, z)	// debug
	fmt.Printf("Bob Verifies :         z= %s\n", z) // Prints on default 851

	// Response 2 - Success/Fail message from server back to client
	if x.Cmp(z) == 0 {
		fmt.Printf("Authorized! Yea\n")
	} else {
		fmt.Printf("Nope nope nope\n")
	}
}
