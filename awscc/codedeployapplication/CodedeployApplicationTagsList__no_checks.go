// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package codedeployapplication

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodedeployApplicationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CodedeployApplicationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodedeployApplicationTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodedeployApplicationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodedeployApplicationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodedeployApplicationTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodedeployApplicationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodedeployApplicationTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

