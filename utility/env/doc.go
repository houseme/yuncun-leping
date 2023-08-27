/*
 * Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

// Package env is a utility package for env.
package env

import (
	"context"
)

// Env .
func Env() *UtilEnv {
	return &UtilEnv{}
}

// UtilEnv .
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
