// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pcscluster

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PcsClusterErrorInfoList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PcsClusterErrorInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PcsClusterErrorInfoList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PcsClusterErrorInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PcsClusterErrorInfoList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PcsClusterErrorInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPcsClusterErrorInfoListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

