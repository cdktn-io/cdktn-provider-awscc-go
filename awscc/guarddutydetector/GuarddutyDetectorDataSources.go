// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutydetector


type GuarddutyDetectorDataSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/guardduty_detector#kubernetes GuarddutyDetector#kubernetes}.
	Kubernetes *GuarddutyDetectorDataSourcesKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/guardduty_detector#malware_protection GuarddutyDetector#malware_protection}.
	MalwareProtection *GuarddutyDetectorDataSourcesMalwareProtection `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/guardduty_detector#s3_logs GuarddutyDetector#s3_logs}.
	S3Logs *GuarddutyDetectorDataSourcesS3Logs `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

