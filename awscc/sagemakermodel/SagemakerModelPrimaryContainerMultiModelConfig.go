// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelPrimaryContainerMultiModelConfig struct {
	// Whether to cache models for a multi-model endpoint.
	//
	// By default, multi-model endpoints cache models so that a model does not have to be loaded into memory each time it is invoked. Some use cases do not benefit from model caching. For example, if an endpoint hosts a large number of models that are each invoked infrequently, the endpoint might perform better if you disable model caching. To disable model caching, set the value of this parameter to `Disabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_model#model_cache_setting SagemakerModel#model_cache_setting}
	ModelCacheSetting *string `field:"optional" json:"modelCacheSetting" yaml:"modelCacheSetting"`
}

