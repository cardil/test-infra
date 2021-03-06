/*
Copyright 2019 The Knative Authors

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

package ops

import (
	clm "knative.dev/pkg/testutils/clustermanager/e2e-tests"
)

// Get gets a GKE cluster
func (rw *RequestWrapper) Get() (*clm.GKECluster, error) {
	rw.Request.SkipCreation = true
	// Reuse `Create` for getting operation, so that we can reuse the same logic
	// such as protected project/cluster etc.
	return rw.Create()
}
