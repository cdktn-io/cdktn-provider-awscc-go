// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluemltransform


type GlueMlTransformTransformParametersFindMatchesParameters struct {
	// The value for accuracy and cost tradeoff. A value of 0.5 means balance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_ml_transform#accuracy_cost_tradeoff GlueMlTransform#accuracy_cost_tradeoff}
	AccuracyCostTradeoff *float64 `field:"optional" json:"accuracyCostTradeoff" yaml:"accuracyCostTradeoff"`
	// If true, forces the output to match the provided labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_ml_transform#enforce_provided_labels GlueMlTransform#enforce_provided_labels}
	EnforceProvidedLabels interface{} `field:"optional" json:"enforceProvidedLabels" yaml:"enforceProvidedLabels"`
	// The value for precision and recall tradeoff. A value of 0.5 means no preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_ml_transform#precision_recall_tradeoff GlueMlTransform#precision_recall_tradeoff}
	PrecisionRecallTradeoff *float64 `field:"optional" json:"precisionRecallTradeoff" yaml:"precisionRecallTradeoff"`
	// The name of a column that uniquely identifies rows in the source table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_ml_transform#primary_key_column_name GlueMlTransform#primary_key_column_name}
	PrimaryKeyColumnName *string `field:"optional" json:"primaryKeyColumnName" yaml:"primaryKeyColumnName"`
}

