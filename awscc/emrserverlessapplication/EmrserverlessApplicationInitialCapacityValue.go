// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationInitialCapacityValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emrserverless_application#worker_configuration EmrserverlessApplication#worker_configuration}.
	WorkerConfiguration *EmrserverlessApplicationInitialCapacityValueWorkerConfiguration `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
	// Initial count of workers to be initialized when an Application is started.
	//
	// This count will be continued to be maintained until the Application is stopped
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emrserverless_application#worker_count EmrserverlessApplication#worker_count}
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
}

