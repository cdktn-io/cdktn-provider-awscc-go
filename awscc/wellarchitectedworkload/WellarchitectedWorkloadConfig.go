// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedworkload

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedWorkloadConfig struct {
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
	// The description for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#description WellarchitectedWorkload#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The environment for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#environment WellarchitectedWorkload#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The list of lenses associated with the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#lenses WellarchitectedWorkload#lenses}
	Lenses *[]*string `field:"required" json:"lenses" yaml:"lenses"`
	// The name of the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#workload_name WellarchitectedWorkload#workload_name}
	WorkloadName *string `field:"required" json:"workloadName" yaml:"workloadName"`
	// The list of Amazon Web Services account IDs associated with the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#account_ids WellarchitectedWorkload#account_ids}
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// The URL of the architectural design for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#architectural_design WellarchitectedWorkload#architectural_design}
	ArchitecturalDesign *string `field:"optional" json:"architecturalDesign" yaml:"architecturalDesign"`
	// The list of Amazon Web Services Regions associated with the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#aws_regions WellarchitectedWorkload#aws_regions}
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// Discovery configuration associated to the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#discovery_config WellarchitectedWorkload#discovery_config}
	DiscoveryConfig *WellarchitectedWorkloadDiscoveryConfig `field:"optional" json:"discoveryConfig" yaml:"discoveryConfig"`
	// The industry for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#industry WellarchitectedWorkload#industry}
	Industry *string `field:"optional" json:"industry" yaml:"industry"`
	// The industry type for the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#industry_type WellarchitectedWorkload#industry_type}
	IndustryType *string `field:"optional" json:"industryType" yaml:"industryType"`
	// The list of non-Amazon Web Services Regions associated with the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#non_aws_regions WellarchitectedWorkload#non_aws_regions}
	NonAwsRegions *[]*string `field:"optional" json:"nonAwsRegions" yaml:"nonAwsRegions"`
	// The notes associated with the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#notes WellarchitectedWorkload#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The review owner of the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#review_owner WellarchitectedWorkload#review_owner}
	ReviewOwner *string `field:"optional" json:"reviewOwner" yaml:"reviewOwner"`
	// The tags associated with the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wellarchitected_workload#tags WellarchitectedWorkload#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

