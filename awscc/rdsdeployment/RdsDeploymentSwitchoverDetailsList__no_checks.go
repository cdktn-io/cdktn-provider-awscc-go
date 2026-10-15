// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdsdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDeploymentSwitchoverDetailsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsDeploymentSwitchoverDetailsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDeploymentSwitchoverDetailsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDeploymentSwitchoverDetailsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDeploymentSwitchoverDetailsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDeploymentSwitchoverDetailsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDeploymentSwitchoverDetailsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

