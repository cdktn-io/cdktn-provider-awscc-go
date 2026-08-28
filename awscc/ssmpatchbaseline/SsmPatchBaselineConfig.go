// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmpatchbaseline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmPatchBaselineConfig struct {
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
	// The name of the patch baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#name SsmPatchBaseline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of rules defining the approval rules for a patch baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#approval_rules SsmPatchBaseline#approval_rules}
	ApprovalRules *SsmPatchBaselineApprovalRules `field:"optional" json:"approvalRules" yaml:"approvalRules"`
	// A list of explicitly approved patches for the baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#approved_patches SsmPatchBaseline#approved_patches}
	ApprovedPatches *[]*string `field:"optional" json:"approvedPatches" yaml:"approvedPatches"`
	// Defines the compliance level for approved patches.
	//
	// This means that if an approved patch is reported as missing, this is the severity of the compliance violation. The default value is UNSPECIFIED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#approved_patches_compliance_level SsmPatchBaseline#approved_patches_compliance_level}
	ApprovedPatchesComplianceLevel *string `field:"optional" json:"approvedPatchesComplianceLevel" yaml:"approvedPatchesComplianceLevel"`
	// Indicates whether the list of approved patches includes non-security updates that should be applied to the instances.
	//
	// The default value is 'false'. Applies to Linux instances only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#approved_patches_enable_non_security SsmPatchBaseline#approved_patches_enable_non_security}
	ApprovedPatchesEnableNonSecurity interface{} `field:"optional" json:"approvedPatchesEnableNonSecurity" yaml:"approvedPatchesEnableNonSecurity"`
	// The compliance status for vendor recommended security updates that are not approved by this patch baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#available_security_updates_compliance_status SsmPatchBaseline#available_security_updates_compliance_status}
	AvailableSecurityUpdatesComplianceStatus *string `field:"optional" json:"availableSecurityUpdatesComplianceStatus" yaml:"availableSecurityUpdatesComplianceStatus"`
	// Set the baseline as default baseline. Only registering to default patch baseline is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#default_baseline SsmPatchBaseline#default_baseline}
	DefaultBaseline interface{} `field:"optional" json:"defaultBaseline" yaml:"defaultBaseline"`
	// The description of the patch baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#description SsmPatchBaseline#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A set of global filters used to include patches in the baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#global_filters SsmPatchBaseline#global_filters}
	GlobalFilters *SsmPatchBaselineGlobalFilters `field:"optional" json:"globalFilters" yaml:"globalFilters"`
	// Defines the operating system the patch baseline applies to. The Default value is WINDOWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#operating_system SsmPatchBaseline#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// PatchGroups is used to associate instances with a specific patch baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#patch_groups SsmPatchBaseline#patch_groups}
	PatchGroups *[]*string `field:"optional" json:"patchGroups" yaml:"patchGroups"`
	// A list of explicitly rejected patches for the baseline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#rejected_patches SsmPatchBaseline#rejected_patches}
	RejectedPatches *[]*string `field:"optional" json:"rejectedPatches" yaml:"rejectedPatches"`
	// The action for Patch Manager to take on patches included in the RejectedPackages list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#rejected_patches_action SsmPatchBaseline#rejected_patches_action}
	RejectedPatchesAction *string `field:"optional" json:"rejectedPatchesAction" yaml:"rejectedPatchesAction"`
	// Information about the patches to use to update the instances, including target operating systems and source repository.
	//
	// Applies to Linux instances only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#sources SsmPatchBaseline#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ssm_patch_baseline#tags SsmPatchBaseline#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

