// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cassandrakeyspace

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CassandraKeyspaceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CassandraKeyspaceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CassandraKeyspaceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CassandraKeyspaceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CassandraKeyspaceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CassandraKeyspaceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CassandraKeyspaceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCassandraKeyspaceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

