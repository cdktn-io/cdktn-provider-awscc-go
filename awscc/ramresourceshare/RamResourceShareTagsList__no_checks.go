// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ramresourceshare

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RamResourceShareTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RamResourceShareTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RamResourceShareTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RamResourceShareTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RamResourceShareTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RamResourceShareTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RamResourceShareTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRamResourceShareTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

