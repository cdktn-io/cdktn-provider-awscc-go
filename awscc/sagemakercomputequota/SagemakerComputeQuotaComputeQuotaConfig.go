// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota


type SagemakerComputeQuotaComputeQuotaConfig struct {
	// Per-instance-type compute quotas for the team.
	//
	// When VCpu / MemoryInGiB / Accelerators are partially specified, SageMaker auto-fills the remaining fields ratio-proportionally to the underlying instance ? Read may return values the customer did not supply (drift may flag this).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#compute_quota_resources SagemakerComputeQuota#compute_quota_resources}
	ComputeQuotaResources interface{} `field:"optional" json:"computeQuotaResources" yaml:"computeQuotaResources"`
	// Team-task preemption policy. Defaults service-side to 'LowerPriority' when omitted on Create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#preempt_team_tasks SagemakerComputeQuota#preempt_team_tasks}
	PreemptTeamTasks *string `field:"optional" json:"preemptTeamTasks" yaml:"preemptTeamTasks"`
	// Idle-resource-sharing strategy for this team's quota. When omitted, the service applies defaults (Strategy=LendAndBorrow, BorrowLimit=50).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#resource_sharing_config SagemakerComputeQuota#resource_sharing_config}
	ResourceSharingConfig *SagemakerComputeQuotaComputeQuotaConfigResourceSharingConfig `field:"optional" json:"resourceSharingConfig" yaml:"resourceSharingConfig"`
}

