// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdsdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDeploymentTasksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsDeploymentTasksList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDeploymentTasksList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDeploymentTasksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDeploymentTasksList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDeploymentTasksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDeploymentTasksListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

