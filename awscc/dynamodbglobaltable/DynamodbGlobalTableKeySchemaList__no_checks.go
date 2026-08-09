// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamodbglobaltable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbGlobalTableKeySchemaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbGlobalTableKeySchemaList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbGlobalTableKeySchemaList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableKeySchemaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableKeySchemaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableKeySchemaList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableKeySchemaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbGlobalTableKeySchemaListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

