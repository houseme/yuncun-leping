// Package env is a utility package for env.
package env

import (
	"context"
)

// Env .
func Env() *UtilEnv {
	return &UtilEnv{}
}

type UtilEnv struct {
}

const (
	// env
	dev  = "dev"
	prod = "prod"
	test = "test"
	// environment
	develop    = "develop"
	production = "production"
)

// Dev .
func (u *UtilEnv) Dev(_ context.Context) string {
	return dev
}

// Prod .
func (u *UtilEnv) Prod(_ context.Context) string {
	return prod
}

// Test .
func (u *UtilEnv) Test(_ context.Context) string {
	return test
}

// Develop .
func (u *UtilEnv) Develop(_ context.Context) string {
	return develop
}

// Production .
func (u *UtilEnv) Production(_ context.Context) string {
	return production
}
