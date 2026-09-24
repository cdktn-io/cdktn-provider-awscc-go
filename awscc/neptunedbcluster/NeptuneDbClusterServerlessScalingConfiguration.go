// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package neptunedbcluster


type NeptuneDbClusterServerlessScalingConfiguration struct {
	// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster.
	//
	// You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptune_db_cluster#max_capacity NeptuneDbCluster#max_capacity}
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster.
	//
	// You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/neptune_db_cluster#min_capacity NeptuneDbCluster#min_capacity}
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

