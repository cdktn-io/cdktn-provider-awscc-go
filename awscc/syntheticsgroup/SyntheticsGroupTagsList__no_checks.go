// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticsgroup

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsGroupTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsGroupTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsGroupTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsGroupTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsGroupTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsGroupTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsGroupTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsGroupTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

