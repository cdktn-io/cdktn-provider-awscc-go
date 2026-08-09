// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apprunnerservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApprunnerServiceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApprunnerServiceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApprunnerServiceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApprunnerServiceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApprunnerServiceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

