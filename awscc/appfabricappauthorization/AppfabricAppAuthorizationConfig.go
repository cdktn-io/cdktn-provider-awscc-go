// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppfabricAppAuthorizationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#app AppfabricAppAuthorization#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle to use for the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#app_bundle_identifier AppfabricAppAuthorization#app_bundle_identifier}
	AppBundleIdentifier *string `field:"required" json:"appBundleIdentifier" yaml:"appBundleIdentifier"`
	// The authorization type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#auth_type AppfabricAppAuthorization#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Contains information about an application tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#tenant AppfabricAppAuthorization#tenant}
	Tenant *AppfabricAppAuthorizationTenant `field:"required" json:"tenant" yaml:"tenant"`
	// Contains credentials for the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#credential AppfabricAppAuthorization#credential}
	Credential *AppfabricAppAuthorizationCredential `field:"optional" json:"credential" yaml:"credential"`
	// Tags to apply to the app authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/appfabric_app_authorization#tags AppfabricAppAuthorization#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

