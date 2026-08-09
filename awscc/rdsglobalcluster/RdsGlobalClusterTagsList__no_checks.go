// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdsglobalcluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsGlobalClusterTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsGlobalClusterTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsGlobalClusterTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsGlobalClusterTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsGlobalClusterTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsGlobalClusterTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsGlobalClusterTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsGlobalClusterTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

