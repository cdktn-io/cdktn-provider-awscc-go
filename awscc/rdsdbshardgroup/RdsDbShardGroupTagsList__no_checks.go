// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdsdbshardgroup

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDbShardGroupTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsDbShardGroupTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDbShardGroupTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDbShardGroupTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsDbShardGroupTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDbShardGroupTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDbShardGroupTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDbShardGroupTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

