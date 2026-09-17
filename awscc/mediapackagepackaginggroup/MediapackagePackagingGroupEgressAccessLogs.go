// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagepackaginggroup


type MediapackagePackagingGroupEgressAccessLogs struct {
	// Sets a custom AWS CloudWatch log group name for egress logs.
	//
	// If a log group name isn't specified, the default name is used: /aws/MediaPackage/VodEgressAccessLogs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediapackage_packaging_group#log_group_name MediapackagePackagingGroup#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

