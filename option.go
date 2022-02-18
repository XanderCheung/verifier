package verifier

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// An Option configures a verifier.
type Option interface {
	Apply(*Verifier)
}

// OptionFunc is a function that configures a verifier.
type OptionFunc func(*Verifier)

// Apply calls f(verifier)
func (f OptionFunc) Apply(verifier *Verifier) {
	f(verifier)
}

// WithJwtTimeFunc can be used to set jwtTimeFunc.
func WithJwtTimeFunc(timeFunc func() time.Time) Option {
	return OptionFunc(func(v *Verifier) {
		v.jwtTimeFunc = timeFunc
		jwt.TimeFunc = timeFunc
	})
}

// WithTokenExpireDuration can be used to set tokenExpireDuration.
func WithTokenExpireDuration(duration time.Duration) Option {
	return OptionFunc(func(v *Verifier) {
		v.tokenExpireDuration = duration
	})
}

// WithAuthExpireDuration can be used to set authExpireDuration.
func WithAuthExpireDuration(duration time.Duration) Option {
	return OptionFunc(func(v *Verifier) {
		v.authExpireDuration = duration
	})
}

// WithTempTokenExpireDuration can be used to set tempTokenExpireDuration.
func WithTempTokenExpireDuration(duration time.Duration) Option {
	return OptionFunc(func(v *Verifier) {
		v.tempTokenExpireDuration = duration
	})
}