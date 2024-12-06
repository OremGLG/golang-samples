// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snippets

// [START compute_consistency_group_list]
import (
	"context"
	"fmt"
	"io"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
)

// listConsistencyGroup get list of disks in consistency group for a project in a given region.
func listConsistencyGroup(w io.Writer, projectID, region, groupName string) error {
	// projectID := "your_project_id"
	// region := "europe-west4"
	// groupName := "your_group_name"

	ctx := context.Background()
	disksClient, err := compute.NewRegionDisksRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewResourcePoliciesRESTClient: %w", err)
	}
	defer disksClient.Close()

	req := &computepb.ListRegionDisksRequest{
		Project: projectID,
		Region:  region,
	}

	it := disksClient.List(ctx, req)
	for {
		disk, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		for _, diskPolicy := range disk.GetResourcePolicies() {
			if strings.Contains(diskPolicy, groupName) {
				fmt.Fprintf(w, "- %s\n", disk.GetName())
			}
		}
	}

	return nil
}

// [END compute_consistency_group_list]