// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainClusterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#cold_storage_options OpensearchserviceDomain#cold_storage_options}.
	ColdStorageOptions *OpensearchserviceDomainClusterConfigColdStorageOptions `field:"optional" json:"coldStorageOptions" yaml:"coldStorageOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#dedicated_master_count OpensearchserviceDomain#dedicated_master_count}.
	DedicatedMasterCount *float64 `field:"optional" json:"dedicatedMasterCount" yaml:"dedicatedMasterCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#dedicated_master_enabled OpensearchserviceDomain#dedicated_master_enabled}.
	DedicatedMasterEnabled interface{} `field:"optional" json:"dedicatedMasterEnabled" yaml:"dedicatedMasterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#dedicated_master_type OpensearchserviceDomain#dedicated_master_type}.
	DedicatedMasterType *string `field:"optional" json:"dedicatedMasterType" yaml:"dedicatedMasterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#instance_count OpensearchserviceDomain#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#instance_type OpensearchserviceDomain#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#multi_az_with_standby_enabled OpensearchserviceDomain#multi_az_with_standby_enabled}.
	MultiAzWithStandbyEnabled interface{} `field:"optional" json:"multiAzWithStandbyEnabled" yaml:"multiAzWithStandbyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#node_options OpensearchserviceDomain#node_options}.
	NodeOptions interface{} `field:"optional" json:"nodeOptions" yaml:"nodeOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#warm_count OpensearchserviceDomain#warm_count}.
	WarmCount *float64 `field:"optional" json:"warmCount" yaml:"warmCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#warm_enabled OpensearchserviceDomain#warm_enabled}.
	WarmEnabled interface{} `field:"optional" json:"warmEnabled" yaml:"warmEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#warm_type OpensearchserviceDomain#warm_type}.
	WarmType *string `field:"optional" json:"warmType" yaml:"warmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#zone_awareness_config OpensearchserviceDomain#zone_awareness_config}.
	ZoneAwarenessConfig *OpensearchserviceDomainClusterConfigZoneAwarenessConfig `field:"optional" json:"zoneAwarenessConfig" yaml:"zoneAwarenessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/opensearchservice_domain#zone_awareness_enabled OpensearchserviceDomain#zone_awareness_enabled}.
	ZoneAwarenessEnabled interface{} `field:"optional" json:"zoneAwarenessEnabled" yaml:"zoneAwarenessEnabled"`
}

