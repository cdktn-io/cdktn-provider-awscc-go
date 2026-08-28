// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputBatchTransformInputDatasetFormat struct {
	// The CSV format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_explainability_job_definition#csv SagemakerModelExplainabilityJobDefinition#csv}
	Csv *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputBatchTransformInputDatasetFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// The Json format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_explainability_job_definition#json SagemakerModelExplainabilityJobDefinition#json}
	Json *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputBatchTransformInputDatasetFormatJson `field:"optional" json:"json" yaml:"json"`
	// A flag indicating if the dataset format is Parquet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_explainability_job_definition#parquet SagemakerModelExplainabilityJobDefinition#parquet}
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

