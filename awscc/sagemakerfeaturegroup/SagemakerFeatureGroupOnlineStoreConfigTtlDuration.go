// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupOnlineStoreConfigTtlDuration struct {
	// Unit of ttl configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_feature_group#unit SagemakerFeatureGroup#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Value of ttl configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_feature_group#value SagemakerFeatureGroup#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

