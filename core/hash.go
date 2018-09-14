package core
import "crypto/sha256"



type Hash []byte



func CalculateHash(data [] byte) Hash{
		sum := sha256.Sum256(data)
		return Hash(sum[:])
}