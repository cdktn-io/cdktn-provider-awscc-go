// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfiguration struct {
	// Contains additional notebook engine MAP<string, string> parameter mappings in the form of key-value pairs.
	//
	// To specify an Athena notebook that the Jupyter server will download and serve, specify a value for the StartSessionRequest$NotebookVersion field, and then add a key named NotebookId to AdditionalConfigs that has the value of the Athena notebook ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#additional_configs AthenaWorkGroup#additional_configs}
	AdditionalConfigs *map[string]*string `field:"optional" json:"additionalConfigs" yaml:"additionalConfigs"`
	// The configuration classifications that can be specified for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#classifications AthenaWorkGroup#classifications}
	Classifications interface{} `field:"optional" json:"classifications" yaml:"classifications"`
	// The number of DPUs to use for the coordinator.
	//
	// A coordinator is a special executor that orchestrates processing work and manages other executors in a notebook session. The default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#coordinator_dpu_size AthenaWorkGroup#coordinator_dpu_size}
	CoordinatorDpuSize *float64 `field:"optional" json:"coordinatorDpuSize" yaml:"coordinatorDpuSize"`
	// The default number of DPUs to use for executors.
	//
	// An executor is the smallest unit of compute that a notebook session can request from Athena. The default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#default_executor_dpu_size AthenaWorkGroup#default_executor_dpu_size}
	DefaultExecutorDpuSize *float64 `field:"optional" json:"defaultExecutorDpuSize" yaml:"defaultExecutorDpuSize"`
	// The maximum number of DPUs that can run concurrently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#max_concurrent_dpus AthenaWorkGroup#max_concurrent_dpus}
	MaxConcurrentDpus *float64 `field:"optional" json:"maxConcurrentDpus" yaml:"maxConcurrentDpus"`
	// Specifies custom jar files and Spark properties for use cases like cluster encryption, table formats, and general Spark tuning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#spark_properties AthenaWorkGroup#spark_properties}
	SparkProperties *map[string]*string `field:"optional" json:"sparkProperties" yaml:"sparkProperties"`
}

