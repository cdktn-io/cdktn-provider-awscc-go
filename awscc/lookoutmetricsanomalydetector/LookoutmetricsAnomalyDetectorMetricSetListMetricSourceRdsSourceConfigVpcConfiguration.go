// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceRdsSourceConfigVpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lookoutmetrics_anomaly_detector#security_group_id_list LookoutmetricsAnomalyDetector#security_group_id_list}.
	SecurityGroupIdList *[]*string `field:"optional" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/lookoutmetrics_anomaly_detector#subnet_id_list LookoutmetricsAnomalyDetector#subnet_id_list}.
	SubnetIdList *[]*string `field:"optional" json:"subnetIdList" yaml:"subnetIdList"`
}

