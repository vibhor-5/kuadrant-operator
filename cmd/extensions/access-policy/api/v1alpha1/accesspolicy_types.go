/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	agenticv1alpha1 "sigs.k8s.io/kube-agentic-networking/api/v1alpha1"
)

// Alias API types directly from sigs.k8s.io/kube-agentic-networking/api/v1alpha1

type AccessPolicy = agenticv1alpha1.XAccessPolicy
type AccessPolicyList = agenticv1alpha1.XAccessPolicyList
type AccessPolicySpec = agenticv1alpha1.AccessPolicySpec
type AccessRule = agenticv1alpha1.AccessRule
type AccessPolicyActionType = agenticv1alpha1.AccessPolicyActionType
type AccessRuleSource = agenticv1alpha1.AccessRuleSource
type AuthorizationSourceType = agenticv1alpha1.AuthorizationSourceType
type AuthorizationSourceSPIFFE = agenticv1alpha1.AuthorizationSourceSPIFFE
type AuthorizationSourceServiceAccount = agenticv1alpha1.AuthorizationSourceServiceAccount
type AuthorizationRule = agenticv1alpha1.AuthorizationRule
type AuthorizationRuleType = agenticv1alpha1.AuthorizationRuleType
type AccessPolicyCELRule = agenticv1alpha1.AccessPolicyCELRule
type MCPAttributes = agenticv1alpha1.MCPAttributes
type MCPMethod = agenticv1alpha1.MCPMethod
type MCPMethodName = agenticv1alpha1.MCPMethodName
type MCPMethodParam = agenticv1alpha1.MCPMethodParam
type MCPBaseProtocolMethodsOption = agenticv1alpha1.MCPBaseProtocolMethodsOption
type AccessPolicyStatus = agenticv1alpha1.AccessPolicyStatus

const (
	ActionTypeAllow                       = agenticv1alpha1.ActionTypeAllow
	ActionTypeExternalAuth                = agenticv1alpha1.ActionTypeExternalAuth
	AuthorizationSourceTypeSPIFFE         = agenticv1alpha1.AuthorizationSourceTypeSPIFFE
	AuthorizationSourceTypeServiceAccount = agenticv1alpha1.AuthorizationSourceTypeServiceAccount
	AuthorizationRuleTypeInline           = agenticv1alpha1.AuthorizationRuleTypeInline
	AuthorizationRuleTypeCEL              = agenticv1alpha1.AuthorizationRuleTypeCEL
	MCPBaseProtocolMethodsOptionSkip      = agenticv1alpha1.MCPBaseProtocolMethodsOptionSkip
	MCPBaseProtocolMethodsOptionMatch     = agenticv1alpha1.MCPBaseProtocolMethodsOptionMatch
	PolicyConditionAccepted               = agenticv1alpha1.PolicyConditionAccepted
	PolicyReasonAccepted                  = agenticv1alpha1.PolicyReasonAccepted
	PolicyLimitPerTargetExceeded          = agenticv1alpha1.PolicyLimitPerTargetExceeded
	PolicyReasonInvalidCEL                = agenticv1alpha1.PolicyReasonInvalidCEL
)

func init() {
	SchemeBuilder.Register(&agenticv1alpha1.XAccessPolicy{}, &agenticv1alpha1.XAccessPolicyList{})
}
