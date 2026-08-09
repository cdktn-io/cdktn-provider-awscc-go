// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package casesfield

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CasesFieldTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CasesFieldTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CasesFieldTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CasesFieldTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CasesFieldTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CasesFieldTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CasesFieldTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCasesFieldTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

