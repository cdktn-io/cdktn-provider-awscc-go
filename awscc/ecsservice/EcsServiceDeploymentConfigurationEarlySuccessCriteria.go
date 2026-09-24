// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceDeploymentConfigurationEarlySuccessCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#enable EcsService#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#healthy_percent EcsService#healthy_percent}.
	HealthyPercent *float64 `field:"optional" json:"healthyPercent" yaml:"healthyPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_service#source_service_revision_cleanup EcsService#source_service_revision_cleanup}.
	SourceServiceRevisionCleanup *string `field:"optional" json:"sourceServiceRevisionCleanup" yaml:"sourceServiceRevisionCleanup"`
}

