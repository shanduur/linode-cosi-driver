// Copyright 2024 Akamai Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e

package e2e_test

import (
	"os"
	"test/e2e/envfuncs"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/env"
)

var (
	testenv env.Environment
)

// TestMain runs all test cases in this package and exits with the right status code. It sets up environment before
// running tests, tears down after each test. Sets of Setup/Teardown functions are registered to run on every Test, so
// that no data remains lingering from earlier runs or other leftovers if they were present due before each test.
func TestMain(m *testing.M) {
	testenv = env.New()

	testenv.BeforeEachTest(
		envfuncs.SetupBucketAccessClasses,
		envfuncs.TeardownBucketAccessClasses,
	)

	testenv.AfterEachTest(
		envfuncs.TeardownBucketAccessClasses,
		envfuncs.TeardownBucketClasses,
	)

	// launch package tests
	os.Exit(testenv.Run(m))
}
