// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package launchwizarddeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LaunchwizardDeploymentConfig struct {
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
	// Workload deployment pattern name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/launchwizard_deployment#deployment_pattern_name LaunchwizardDeployment#deployment_pattern_name}
	DeploymentPatternName *string `field:"required" json:"deploymentPatternName" yaml:"deploymentPatternName"`
	// Name of LaunchWizard deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/launchwizard_deployment#name LaunchwizardDeployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Workload Name for LaunchWizard deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/launchwizard_deployment#workload_name LaunchwizardDeployment#workload_name}
	WorkloadName *string `field:"required" json:"workloadName" yaml:"workloadName"`
	// LaunchWizard deployment specifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/launchwizard_deployment#specifications LaunchwizardDeployment#specifications}
	Specifications *map[string]*string `field:"optional" json:"specifications" yaml:"specifications"`
	// Tags for LaunchWizard deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/launchwizard_deployment#tags LaunchwizardDeployment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

