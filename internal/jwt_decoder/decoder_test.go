package jwtdecoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var key = []byte("11111111111111111111111111111111")

var encoder = NewJwtDecoder(*NewJwtHS256Verifier(key))

func TestHS512(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcGFjaGUiOiJrYWZrYSIsImZvbyI6ImJhciJ9.PENqb3G3J-e15p0pFQVF2m5F3j_QK_iZZj4ilDvEssy44SYmKVs2VoazdHPbMuHXaXXb4oj9Qbah6Ke9_0UtjA"

	payload, err := encoder.Decode(token)

	assert.Nil(t, err)
	assert.Contains(t, payload, "foo")
	assert.Contains(t, payload, "apache")
	assert.Equal(t, payload["foo"], "bar")
	assert.Equal(t, payload["apache"], "kafka")
}
