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
