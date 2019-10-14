// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cfgiface_test

import (
	"context"

	"github.com/googleapis/google-cloud-go-testing/runtimeconfig/cfgiface"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
)

func ExampleAdaptService() {
	ctx := context.Background()

	s, err := runtimeconfig.NewService(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	service := cfgiface.AdaptService(s)

	_, err = service.Projects().Configs().Variables().List("my-config").Do()
	if err != nil {
		// TODO: Handle error.
	}
}
