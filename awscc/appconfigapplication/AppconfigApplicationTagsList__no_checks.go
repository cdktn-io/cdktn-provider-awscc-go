// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfigapplication

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppconfigApplicationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppconfigApplicationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppconfigApplicationTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppconfigApplicationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppconfigApplicationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppconfigApplicationTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppconfigApplicationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppconfigApplicationTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

