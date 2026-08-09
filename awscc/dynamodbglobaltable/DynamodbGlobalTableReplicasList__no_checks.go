// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamodbglobaltable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbGlobalTableReplicasList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbGlobalTableReplicasList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicasList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbGlobalTableReplicasListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

