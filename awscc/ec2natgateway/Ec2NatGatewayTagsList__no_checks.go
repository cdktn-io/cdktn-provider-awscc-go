// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ec2natgateway

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2NatGatewayTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2NatGatewayTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2NatGatewayTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2NatGatewayTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2NatGatewayTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2NatGatewayTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2NatGatewayTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2NatGatewayTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

