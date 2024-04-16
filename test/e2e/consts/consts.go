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

package consts

import (
	"context"
	"testing"

	"github.com/linode/linode-cosi-driver/pkg/servers/provisioner"
	"github.com/linode/linodego"
)

const (
	Bucket           string = "bucket"
	BucketGreenfield string = "greenfield"
	BucketBrownfield string = "brownfield"

	ACL                  string = provisioner.ParamACL
	ACLPrivate                  = string(linodego.ACLPrivate)
	ACLAuthenticatedRead        = string(linodego.ACLAuthenticatedRead)
	ACLPublicRead               = string(linodego.ACLPublicRead)
	ACLPublicReadWrite          = string(linodego.ACLPublicReadWrite)

	CORS         string = provisioner.ParamCORS
	CORSEnabled         = string(provisioner.ParamCORSValueEnabled)
	CORSDisabled        = string(provisioner.ParamCORSValueDisabled)

	Permissions   string = provisioner.ParamPermissions
	PermissionsRO        = string(provisioner.ParamPermissionsValueReadOnly)
	PermissionsRW        = string(provisioner.ParamPermissionsValueReadWrite)
)

func GetValue[T any](ctx context.Context, t *testing.T, key string, defaultValue T) T {
	value, ok := ctx.Value(key).(T)
	if !ok {
		t.Logf("value %q not set, using default %#+v",
			key, defaultValue)
		value = defaultValue
	}

	return value
}
