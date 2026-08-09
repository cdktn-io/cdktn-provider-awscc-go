// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableLocalSecondaryIndexesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableLocalSecondaryIndexesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

