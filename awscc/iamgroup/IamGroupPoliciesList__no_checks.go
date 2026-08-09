// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iamgroup

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamGroupPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IamGroupPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IamGroupPoliciesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IamGroupPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IamGroupPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IamGroupPoliciesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IamGroupPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIamGroupPoliciesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

