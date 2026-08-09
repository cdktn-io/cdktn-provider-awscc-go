// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apptesttestcase

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApptestTestCaseStepsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApptestTestCaseStepsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApptestTestCaseStepsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApptestTestCaseStepsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApptestTestCaseStepsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApptestTestCaseStepsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApptestTestCaseStepsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApptestTestCaseStepsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

