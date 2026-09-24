// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution


type PersonalizeSolutionSolutionConfigHpoConfigHpoResourceConfig struct {
	// The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#max_number_of_training_jobs PersonalizeSolution#max_number_of_training_jobs}
	MaxNumberOfTrainingJobs *string `field:"optional" json:"maxNumberOfTrainingJobs" yaml:"maxNumberOfTrainingJobs"`
	// The maximum number of parallel training jobs when you create a solution version.
	//
	// The maximum value for maxParallelTrainingJobs is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/personalize_solution#max_parallel_training_jobs PersonalizeSolution#max_parallel_training_jobs}
	MaxParallelTrainingJobs *string `field:"optional" json:"maxParallelTrainingJobs" yaml:"maxParallelTrainingJobs"`
}

