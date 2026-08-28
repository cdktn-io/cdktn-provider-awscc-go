// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deadlinebudget

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeadlineBudgetActionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeadlineBudgetActionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeadlineBudgetActionsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetActionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetActionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetActionsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeadlineBudgetActionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeadlineBudgetActionsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

