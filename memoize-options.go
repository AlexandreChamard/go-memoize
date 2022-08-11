package memoize

import "time"

type memoizationOption struct {
	maxMemoize uint
	expiration time.Duration
}

func NewMemoizeOption() *memoizationOption {
	return &memoizationOption{
		maxMemoize: 1 << 16,
		expiration: 10 * time.Second,
	}
}

func (this *memoizationOption) MaxMemoize(maxMemoize uint) *memoizationOption {
	this.maxMemoize = maxMemoize
	return this
}

func (this *memoizationOption) NoMemoizationLimit() *memoizationOption {
	return this.MaxMemoize(0)
}

func (this *memoizationOption) Expiration(expiration time.Duration) *memoizationOption {
	this.expiration = expiration
	return this
}

func (this *memoizationOption) NoExpiration() *memoizationOption {
	return this.Expiration(0)
}
