// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigShapBaselineConfig struct {
	// The MIME type of the baseline data. Choose from 'text/csv' or 'application/jsonlines'. Defaults to 'text/csv'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#mime_type SagemakerEndpointConfigA#mime_type}
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// The inline SHAP baseline data in string format.
	//
	// ShapBaseline can have one or multiple records to be used as the baseline dataset. The format of the SHAP baseline file should be the same format as the training dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#shap_baseline SagemakerEndpointConfigA#shap_baseline}
	ShapBaseline *string `field:"optional" json:"shapBaseline" yaml:"shapBaseline"`
	// The uniform resource identifier (URI) of the S3 bucket where the SHAP baseline file is stored.
	//
	// The format of the SHAP baseline file should be the same format as the format of the training dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_endpoint_config#shap_baseline_uri SagemakerEndpointConfigA#shap_baseline_uri}
	ShapBaselineUri *string `field:"optional" json:"shapBaselineUri" yaml:"shapBaselineUri"`
}

