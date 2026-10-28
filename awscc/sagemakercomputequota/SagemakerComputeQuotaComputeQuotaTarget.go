// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota


type SagemakerComputeQuotaComputeQuotaTarget struct {
	// The name of the team.
	//
	// Materialized as Kubernetes namespace 'hyperpod-ns-<TeamName>' on the EKS cluster, so the value must satisfy RFC 1123 namespace constraints. Reassignment is not supported (createOnly).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#team_name SagemakerComputeQuota#team_name}
	TeamName *string `field:"required" json:"teamName" yaml:"teamName"`
	// Relative fair-share weight when fair-share scheduling is enabled on the parent ClusterSchedulerConfig. Defaults service-side to 0 when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#fair_share_weight SagemakerComputeQuota#fair_share_weight}
	FairShareWeight *float64 `field:"optional" json:"fairShareWeight" yaml:"fairShareWeight"`
}

