package dto

type Verifiable interface {
	Verify() error
}
