// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deadlinebudget

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeadlineBudgetTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeadlineBudgetTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeadlineBudgetTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeadlineBudgetTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

