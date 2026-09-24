// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment


type Greengrassv2DeploymentIotJobConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/greengrassv2_deployment#abort_config Greengrassv2Deployment#abort_config}.
	AbortConfig *Greengrassv2DeploymentIotJobConfigurationAbortConfig `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/greengrassv2_deployment#job_executions_rollout_config Greengrassv2Deployment#job_executions_rollout_config}.
	JobExecutionsRolloutConfig *Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfig `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/greengrassv2_deployment#timeout_config Greengrassv2Deployment#timeout_config}.
	TimeoutConfig *Greengrassv2DeploymentIotJobConfigurationTimeoutConfig `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

