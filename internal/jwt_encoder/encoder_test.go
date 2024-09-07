package jwtencoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var key = []byte("11111111111111111111111111111111")

var encoder = NewJwtEncoder(*NewJwtHS256Signer(key))

func TestHS512(t *testing.T) {
  payload := map[string]any {
    "foo": "bar",
    "apache": "kafka",
  }

  res, err := encoder.Encode(payload)

  assert.Nil(t, err)
  assert.Equal(t, res, "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcGFjaGUiOiJrYWZrYSIsImZvbyI6ImJhciJ9.PENqb3G3J-e15p0pFQVF2m5F3j_QK_iZZj4ilDvEssy44SYmKVs2VoazdHPbMuHXaXXb4oj9Qbah6Ke9_0UtjA")
}
