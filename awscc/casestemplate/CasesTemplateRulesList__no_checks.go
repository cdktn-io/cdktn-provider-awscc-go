// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package casestemplate

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CasesTemplateRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CasesTemplateRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CasesTemplateRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CasesTemplateRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CasesTemplateRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CasesTemplateRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CasesTemplateRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCasesTemplateRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

