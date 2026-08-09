// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package finspaceenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FinspaceEnvironmentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FinspaceEnvironmentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FinspaceEnvironmentTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FinspaceEnvironmentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FinspaceEnvironmentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FinspaceEnvironmentTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FinspaceEnvironmentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFinspaceEnvironmentTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

