package envfuncs

import (
	"context"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

func SetupBucketClasses(ctx context.Context, cfg *envconf.Config, t *testing.T) (context.Context, error) {
	return ctx, nil
}

func TeardownBucketClasses(ctx context.Context, cfg *envconf.Config, t *testing.T) (context.Context, error) {
	return ctx, nil
}
