// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eventsrule

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventsRuleTargetsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EventsRuleTargetsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EventsRuleTargetsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EventsRuleTargetsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EventsRuleTargetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventsRuleTargetsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EventsRuleTargetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEventsRuleTargetsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

