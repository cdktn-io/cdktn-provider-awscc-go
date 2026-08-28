// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderinfrastructureconfiguration


type ImagebuilderInfrastructureConfigurationLogging struct {
	// The S3 path in which to store the logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_infrastructure_configuration#s3_logs ImagebuilderInfrastructureConfiguration#s3_logs}
	S3Logs *ImagebuilderInfrastructureConfigurationLoggingS3Logs `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}

