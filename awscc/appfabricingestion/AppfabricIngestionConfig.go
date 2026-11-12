// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppfabricIngestionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/appfabric_ingestion#app AppfabricIngestion#app}
	App *string `field:"required" json:"app" yaml:"app"`
	// The Amazon Resource Name (ARN) of the app bundle for the ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/appfabric_ingestion#app_bundle_arn AppfabricIngestion#app_bundle_arn}
	AppBundleArn *string `field:"required" json:"appBundleArn" yaml:"appBundleArn"`
	// The type of the ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/appfabric_ingestion#ingestion_type AppfabricIngestion#ingestion_type}
	IngestionType *string `field:"required" json:"ingestionType" yaml:"ingestionType"`
	// The ID of the application tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/appfabric_ingestion#tenant_id AppfabricIngestion#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/appfabric_ingestion#tags AppfabricIngestion#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

