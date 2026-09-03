// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package m2deployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type M2DeploymentConfig struct {
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
	// The application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/m2_deployment#application_id M2Deployment#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The version number of the application to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/m2_deployment#application_version M2Deployment#application_version}
	ApplicationVersion *float64 `field:"required" json:"applicationVersion" yaml:"applicationVersion"`
	// The environment ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/m2_deployment#environment_id M2Deployment#environment_id}
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}

