// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformTransformParameters struct {
	// The type of machine learning transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_ml_transform#transform_type GlueMlTransform#transform_type}
	TransformType *string `field:"required" json:"transformType" yaml:"transformType"`
	// The parameters to configure the find matches transform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_ml_transform#find_matches_parameters GlueMlTransform#find_matches_parameters}
	FindMatchesParameters *GlueMlTransformTransformParametersFindMatchesParameters `field:"optional" json:"findMatchesParameters" yaml:"findMatchesParameters"`
}

