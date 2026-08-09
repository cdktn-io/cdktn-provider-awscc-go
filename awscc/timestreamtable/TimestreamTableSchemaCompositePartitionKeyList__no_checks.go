// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package timestreamtable

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TimestreamTableSchemaCompositePartitionKeyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TimestreamTableSchemaCompositePartitionKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TimestreamTableSchemaCompositePartitionKeyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableSchemaCompositePartitionKeyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableSchemaCompositePartitionKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableSchemaCompositePartitionKeyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TimestreamTableSchemaCompositePartitionKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTimestreamTableSchemaCompositePartitionKeyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

