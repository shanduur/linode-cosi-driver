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
