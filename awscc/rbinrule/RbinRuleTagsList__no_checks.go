// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rbinrule

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RbinRuleTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RbinRuleTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RbinRuleTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RbinRuleTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RbinRuleTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RbinRuleTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RbinRuleTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRbinRuleTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

