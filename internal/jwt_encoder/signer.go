package jwtencoder

type JwtSigner interface {
	Sign(data []byte) ([]byte, error)
	GetAlgorithm() string
}
