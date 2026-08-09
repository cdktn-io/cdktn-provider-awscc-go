// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package devopsagentservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DevopsagentServiceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DevopsagentServiceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DevopsagentServiceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DevopsagentServiceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DevopsagentServiceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DevopsagentServiceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DevopsagentServiceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDevopsagentServiceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

