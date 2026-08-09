// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ec2securitygroup

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2SecurityGroupSecurityGroupEgressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2SecurityGroupSecurityGroupEgressListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

