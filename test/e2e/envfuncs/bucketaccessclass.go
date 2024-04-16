package envfuncs

import (
	"context"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

func SetupBucketAccessClasses(ctx context.Context, cfg *envconf.Config, t *testing.T) (context.Context, error) {
	return ctx, nil
}

func TeardownBucketAccessClasses(ctx context.Context, cfg *envconf.Config, t *testing.T) (context.Context, error) {
	return ctx, nil
}
