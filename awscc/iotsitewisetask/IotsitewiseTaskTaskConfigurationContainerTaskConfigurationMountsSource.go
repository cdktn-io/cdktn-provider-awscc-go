// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisetask


type IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMountsSource struct {
	// Configures a mount that reads from an Amazon S3 access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#s3_access_point IotsitewiseTask#s3_access_point}
	S3AccessPoint *IotsitewiseTaskTaskConfigurationContainerTaskConfigurationMountsSourceS3AccessPoint `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
}

