// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cassandratable

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CassandraTablePartitionKeyColumnsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCassandraTablePartitionKeyColumnsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

