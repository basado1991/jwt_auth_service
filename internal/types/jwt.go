package types

const JWT_TYPE = "JWT"

type JwtHeader struct {
  Algorithm string `json:"alg"`
  Type string `json:"typ"`
}
