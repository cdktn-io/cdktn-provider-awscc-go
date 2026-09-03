// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2ServiceConfig struct {
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
	// The name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#name Resiliencehubv2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// AWS regions for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#regions Resiliencehubv2Service#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Assertions associated with this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#assertions Resiliencehubv2Service#assertions}
	Assertions interface{} `field:"optional" json:"assertions" yaml:"assertions"`
	// Systems associated with this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#associated_systems Resiliencehubv2Service#associated_systems}
	AssociatedSystems interface{} `field:"optional" json:"associatedSystems" yaml:"associatedSystems"`
	// Dependency discovery state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#dependency_discovery Resiliencehubv2Service#dependency_discovery}
	DependencyDiscovery *string `field:"optional" json:"dependencyDiscovery" yaml:"dependencyDiscovery"`
	// The description of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#description Resiliencehubv2Service#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Input sources for this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#input_sources Resiliencehubv2Service#input_sources}
	InputSources interface{} `field:"optional" json:"inputSources" yaml:"inputSources"`
	// The KMS key ID for encrypting service data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#kms_key_id Resiliencehubv2Service#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#permission_model Resiliencehubv2Service#permission_model}.
	PermissionModel *Resiliencehubv2ServicePermissionModel `field:"optional" json:"permissionModel" yaml:"permissionModel"`
	// The ARN of the resilience policy to associate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#policy_arn Resiliencehubv2Service#policy_arn}
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
	// Configuration for automatic report generation on a Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#report_configuration Resiliencehubv2Service#report_configuration}
	ReportConfiguration *Resiliencehubv2ServiceReportConfiguration `field:"optional" json:"reportConfiguration" yaml:"reportConfiguration"`
	// Tags assigned to the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#tags Resiliencehubv2Service#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

