// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ec2prefixlist

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2PrefixListTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_Ec2PrefixListTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2PrefixListTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2PrefixListTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2PrefixListTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2PrefixListTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2PrefixListTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2PrefixListTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

