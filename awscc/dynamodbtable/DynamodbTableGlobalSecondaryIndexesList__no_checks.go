// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableGlobalSecondaryIndexesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

