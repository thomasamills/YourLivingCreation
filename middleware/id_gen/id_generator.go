package id_gen

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

type ULIDGenerator interface {
	GetULIDfromtime() string
}

type ULIDGeneratorImpl struct{}

func NewULIDGenerator() ULIDGenerator {
	return &ULIDGeneratorImpl{}
}

func (ig *ULIDGeneratorImpl) GetULIDfromtime() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
