// Copyright 2019 The Terraformer Authors.
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

package ns1

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	ns1 "github.com/ns1/ns1-go"
)

type ZoneGenerator struct {
	Ns1Service
}

func (g *ZoneGenerator) createZoneResources(client *ns1.APIClient) error {
	zones, err := client.GetZones()
	if err != nil {
		return err
	}

	for _, zone := range zones {
		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			zone.Id,
			zone.Id,
			"ns1_zone",
			"ns1",
			[]string{}))
	}

	return nil
}

func (g *ZoneGenerator) InitResources() error {
	client := ns1.New(g.Args["api_key"].(string))

	if err := g.createZoneResources(client); err != nil {
		return err
	}

	return nil
}
