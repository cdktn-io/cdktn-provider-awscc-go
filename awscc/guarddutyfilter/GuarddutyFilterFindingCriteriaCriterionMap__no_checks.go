// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package guarddutyfilter

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionMap) validateGetParameters(key *string) error {
	return nil
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewGuarddutyFilterFindingCriteriaCriterionMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

