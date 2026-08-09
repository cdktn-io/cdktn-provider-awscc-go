// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ec2instance

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2InstanceLicenseSpecificationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2InstanceLicenseSpecificationsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2InstanceLicenseSpecificationsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceLicenseSpecificationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceLicenseSpecificationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceLicenseSpecificationsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2InstanceLicenseSpecificationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2InstanceLicenseSpecificationsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

