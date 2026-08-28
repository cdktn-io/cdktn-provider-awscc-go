// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package codeconnectionshost

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeconnectionsHostTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CodeconnectionsHostTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodeconnectionsHostTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodeconnectionsHostTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodeconnectionsHostTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodeconnectionsHostTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodeconnectionsHostTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodeconnectionsHostTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

