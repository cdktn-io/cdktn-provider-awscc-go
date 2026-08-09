// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster


type PcsClusterSlurmConfiguration struct {
	// The accounting configuration includes configurable settings for Slurm accounting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#accounting PcsCluster#accounting}
	Accounting *PcsClusterSlurmConfigurationAccounting `field:"optional" json:"accounting" yaml:"accounting"`
	// The shared Slurm key for authentication, also known as the cluster secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#auth_key PcsCluster#auth_key}
	AuthKey *PcsClusterSlurmConfigurationAuthKey `field:"optional" json:"authKey" yaml:"authKey"`
	// Additional cgroup-specific configuration that directly maps to cgroup.conf settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#cgroup_custom_settings PcsCluster#cgroup_custom_settings}
	CgroupCustomSettings interface{} `field:"optional" json:"cgroupCustomSettings" yaml:"cgroupCustomSettings"`
	// JWT authentication configuration for Slurm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#jwt_auth PcsCluster#jwt_auth}
	JwtAuth *PcsClusterSlurmConfigurationJwtAuth `field:"optional" json:"jwtAuth" yaml:"jwtAuth"`
	// The time before an idle node is scaled down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#scale_down_idle_time_in_seconds PcsCluster#scale_down_idle_time_in_seconds}
	ScaleDownIdleTimeInSeconds *float64 `field:"optional" json:"scaleDownIdleTimeInSeconds" yaml:"scaleDownIdleTimeInSeconds"`
	// Additional Slurm-specific configuration that directly maps to Slurm settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#slurm_custom_settings PcsCluster#slurm_custom_settings}
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
	// Additional slurmdbd-specific configuration that directly maps to slurmdbd.conf settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#slurmdbd_custom_settings PcsCluster#slurmdbd_custom_settings}
	SlurmdbdCustomSettings interface{} `field:"optional" json:"slurmdbdCustomSettings" yaml:"slurmdbdCustomSettings"`
	// The SlurmRest configuration includes configurable settings for Slurm Rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_cluster#slurm_rest PcsCluster#slurm_rest}
	SlurmRest *PcsClusterSlurmConfigurationSlurmRest `field:"optional" json:"slurmRest" yaml:"slurmRest"`
}

