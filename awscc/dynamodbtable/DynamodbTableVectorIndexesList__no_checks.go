// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableVectorIndexesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableVectorIndexesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableVectorIndexesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableVectorIndexesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableVectorIndexesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableVectorIndexesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableVectorIndexesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableVectorIndexesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

