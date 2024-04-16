package assessments

import (
	"context"
	"test/e2e/consts"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

func BucketAccessCreated(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	permissions := consts.GetValue(ctx, t, consts.Permissions, consts.PermissionsRW)

	_ = permissions // FIXME: placeholder

	return ctx
}
