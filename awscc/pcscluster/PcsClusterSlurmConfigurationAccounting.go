// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster


type PcsClusterSlurmConfigurationAccounting struct {
	// The default value for all purge settings for `slurmdbd.conf`. For more information, see the [slurmdbd.conf documentation at SchedMD](https://slurm.schedmd.com/slurmdbd.conf.html). The default value is `-1`. A value of `-1` means there is no purge time and records persist as long as the cluster exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_cluster#default_purge_time_in_days PcsCluster#default_purge_time_in_days}
	DefaultPurgeTimeInDays *float64 `field:"optional" json:"defaultPurgeTimeInDays" yaml:"defaultPurgeTimeInDays"`
	// The default value is `NONE`. A value of `STANDARD` means that Slurm accounting is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_cluster#mode PcsCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

