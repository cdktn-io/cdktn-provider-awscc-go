// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package quicksightagent

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QuicksightAgentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (q *jsiiProxy_QuicksightAgentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QuicksightAgentTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QuicksightAgentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QuicksightAgentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QuicksightAgentTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QuicksightAgentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQuicksightAgentTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

