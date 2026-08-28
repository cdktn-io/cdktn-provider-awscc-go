// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionNetworkConfig struct {
	// Whether to encrypt all communications between distributed processing jobs.
	//
	// Choose True to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_bias_job_definition#enable_inter_container_traffic_encryption SagemakerModelBiasJobDefinition#enable_inter_container_traffic_encryption}
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Whether to allow inbound and outbound network calls to and from the containers used for the processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_bias_job_definition#enable_network_isolation SagemakerModelBiasJobDefinition#enable_network_isolation}
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies a VPC that your training jobs and hosted models have access to.
	//
	// Control access to and from your training and model containers by configuring the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_bias_job_definition#vpc_config SagemakerModelBiasJobDefinition#vpc_config}
	VpcConfig *SagemakerModelBiasJobDefinitionNetworkConfigVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

