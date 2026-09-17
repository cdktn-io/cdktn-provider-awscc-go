// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsAutoPatchConfigPatchSchedule struct {
	// The date and time of the next scheduled patch, set by the system when a patch AMI is detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#next_patch_date SagemakerCluster#next_patch_date}
	NextPatchDate *string `field:"optional" json:"nextPatchDate" yaml:"nextPatchDate"`
}

