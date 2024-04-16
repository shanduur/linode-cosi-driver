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
	"test/e2e/assessments"
	"test/e2e/consts"
	"test/e2e/featurebuilder"
	"testing"
)

// TestResourceProvisioning executest series of features that confirm valid execution of provisioner server.
func TestResourceProvisioning(t *testing.T) {
	matrix := featurebuilder.NewMatrix("",
		[]featurebuilder.LabelMatrix{
			{
				Key: consts.Bucket,
				ValueMatrix: []string{
					consts.BucketGreenfield,
					consts.BucketBrownfield,
				},
			},
			{
				Key: consts.ACL,
				ValueMatrix: []string{
					consts.ACLPrivate,
					consts.ACLAuthenticatedRead,
					consts.ACLPublicRead,
					consts.ACLPublicReadWrite,
				},
			},
			{
				Key: consts.CORS,
				ValueMatrix: []string{
					consts.CORSEnabled,
					consts.CORSDisabled,
				},
			},
			{
				Key: consts.Permissions,
				ValueMatrix: []string{
					consts.PermissionsRO,
					consts.PermissionsRW,
				},
			},
		},
		featurebuilder.CommonAssesments{
			"Bucket created":       assessments.BucketCreated,
			"BucketAccess created": assessments.BucketAccessCreated,
		},
	)

	for _, f := range matrix {
		testenv.Test(t, f.Feature())
	}
}
