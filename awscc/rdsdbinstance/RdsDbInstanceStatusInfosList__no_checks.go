// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdsdbinstance

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsDbInstanceStatusInfosList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsDbInstanceStatusInfosList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsDbInstanceStatusInfosList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsDbInstanceStatusInfosList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsDbInstanceStatusInfosList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsDbInstanceStatusInfosList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsDbInstanceStatusInfosListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

