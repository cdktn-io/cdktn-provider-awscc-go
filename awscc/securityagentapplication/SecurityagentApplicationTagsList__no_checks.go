// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package securityagentapplication

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityagentApplicationTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecurityagentApplicationTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityagentApplicationTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityagentApplicationTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityagentApplicationTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityagentApplicationTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityagentApplicationTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityagentApplicationTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

