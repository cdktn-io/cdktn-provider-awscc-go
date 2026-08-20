// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package accountaccessapplication

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountaccessApplicationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccountaccessApplicationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccountaccessApplicationTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccountaccessApplicationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccountaccessApplicationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccountaccessApplicationTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccountaccessApplicationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccountaccessApplicationTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

