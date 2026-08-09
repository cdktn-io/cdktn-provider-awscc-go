// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rbinrule

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RbinRuleResourceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RbinRuleResourceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RbinRuleResourceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RbinRuleResourceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RbinRuleResourceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RbinRuleResourceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RbinRuleResourceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRbinRuleResourceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

