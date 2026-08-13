// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver


type AppsyncResolverRuntime struct {
	// The ``name`` of the runtime to use. Currently, the only allowed value is ``APPSYNC_JS``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_resolver#name AppsyncResolver#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ``version`` of the runtime to use. Currently, the only allowed version is ``1.0.0``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_resolver#runtime_version AppsyncResolver#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
}

