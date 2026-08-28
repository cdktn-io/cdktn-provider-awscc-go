// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupOnlineStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#enable_online_store SagemakerFeatureGroup#enable_online_store}.
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#security_config SagemakerFeatureGroup#security_config}.
	SecurityConfig *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#storage_type SagemakerFeatureGroup#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// TTL configuration of the feature group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_feature_group#ttl_duration SagemakerFeatureGroup#ttl_duration}
	TtlDuration *SagemakerFeatureGroupOnlineStoreConfigTtlDuration `field:"optional" json:"ttlDuration" yaml:"ttlDuration"`
}

