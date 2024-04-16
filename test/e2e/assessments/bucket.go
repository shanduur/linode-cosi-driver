package assessments

import (
	"context"
	"test/e2e/consts"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

func BucketCreated(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	acl := consts.GetValue(ctx, t, consts.ACL, consts.ACLPrivate)
	cors := consts.GetValue(ctx, t, consts.CORS, consts.CORSDisabled)

	_, _ = acl, cors // FIXME: placeholder

	return ctx
}
