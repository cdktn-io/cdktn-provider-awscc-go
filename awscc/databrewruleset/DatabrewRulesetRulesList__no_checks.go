// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databrewruleset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabrewRulesetRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabrewRulesetRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabrewRulesetRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabrewRulesetRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabrewRulesetRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabrewRulesetRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabrewRulesetRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabrewRulesetRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

