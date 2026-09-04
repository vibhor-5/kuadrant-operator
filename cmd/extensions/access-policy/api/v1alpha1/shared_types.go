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

// Alias shared HTTP types directly from sigs.k8s.io/kube-agentic-networking/api/v1alpha1

type HTTPPathMatch = agenticv1alpha1.HTTPPathMatch
type PathMatchType = agenticv1alpha1.PathMatchType
type HTTPHeaderMatch = agenticv1alpha1.HTTPHeaderMatch
type HeaderMatchType = agenticv1alpha1.HeaderMatchType
type HTTPHeaderName = agenticv1alpha1.HTTPHeaderName
type HTTPMethod = agenticv1alpha1.HTTPMethod
type Hostname = agenticv1alpha1.Hostname
type PortNumber = agenticv1alpha1.PortNumber

const (
	PathMatchExact               = agenticv1alpha1.PathMatchExact
	PathMatchPathPrefix          = agenticv1alpha1.PathMatchPathPrefix
	PathMatchRegularExpression   = agenticv1alpha1.PathMatchRegularExpression
	HeaderMatchExact             = agenticv1alpha1.HeaderMatchExact
	HeaderMatchRegularExpression = agenticv1alpha1.HeaderMatchRegularExpression
)
