package jwtdecoder

type JwtVerifier interface {
	Verify(data []byte, signature []byte) (bool, error)
}
