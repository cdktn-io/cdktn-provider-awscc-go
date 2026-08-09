// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package bedrockagentcoreharness

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) validateGetParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *map[string]*BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfig:
		val := val.(*map[string]*BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfig)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case map[string]*BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfig:
		val_ := val.(map[string]*BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfig)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *map[string]*BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfig; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBedrockagentcoreHarnessMemoryAgentCoreMemoryConfigurationRetrievalConfigMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

