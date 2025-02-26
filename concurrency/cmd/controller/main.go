/*
Copyright 2021 The Tekton Authors
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

package main

import (
	"github.com/tektoncd/experimental/concurrency/pkg/apis/concurrency/v1alpha1"
	"github.com/tektoncd/experimental/concurrency/pkg/reconciler/concurrency"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	filteredinformerfactory "knative.dev/pkg/client/injection/kube/informers/factory/filtered"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/signals"
)

func main() {
	ctx := filteredinformerfactory.WithSelectors(signals.NewContext(), v1alpha1.ManagedByLabelKey)
	opts := &pipeline.Options{}
	sharedmain.MainWithContext(ctx, concurrency.ControllerName, concurrency.NewController(opts))
}
