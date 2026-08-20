// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecrrepository


type EcrRepositoryImageScanningConfiguration struct {
	// The setting that determines whether images are scanned after being pushed to a repository.
	//
	// If set to ``true``, images will be scanned after being pushed. If this parameter is not specified, it will default to ``false`` and images will not be scanned unless a scan is manually started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecr_repository#scan_on_push EcrRepository#scan_on_push}
	ScanOnPush interface{} `field:"optional" json:"scanOnPush" yaml:"scanOnPush"`
}

