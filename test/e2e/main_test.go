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
