// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package snstopic

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsTopicSubscriptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SnsTopicSubscriptionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SnsTopicSubscriptionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SnsTopicSubscriptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSnsTopicSubscriptionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

