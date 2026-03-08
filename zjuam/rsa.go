package zjuam

import (
	"math/big"
)

// rsaEncrypt performs raw RSA encryption (m^e mod n) as required by ZJUAM CAS.
// The exponent and modulus are hex-encoded strings from the ZJUAM public key endpoint.
func rsaEncrypt(password, exponent, modulus string) string {
	n := new(big.Int)
	n.SetString(modulus, 16)

	e := new(big.Int)
	e.SetString(exponent, 16)

	m := new(big.Int)
	m.SetBytes([]byte(password))

	c := new(big.Int)
	c.Exp(m, e, n)

	return c.Text(16)
}
