// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package networkmanagersitetositevpnattachment

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRules:
		val := val.(*[]*NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRules)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRules:
		val_ := val.([]*NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRules)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRules; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewNetworkmanagerSiteToSiteVpnAttachmentNatGatewayRulesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

