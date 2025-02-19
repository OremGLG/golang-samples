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

// [START compute_disk_start_replication]
import (
	"context"
	"fmt"
	"io"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

// startReplication starts disk replication in a project for a given zone.
func startReplication(
	w io.Writer,
	projectID, zone, diskName, primaryDiskName, primaryZone string,
) error {
	// projectID := "your_project_id"
	// zone := "europe-west4-b"
	// diskName := "your_disk_name"
	// primaryDiskName := "your_disk_name2"
	// primaryZone := "europe-west2-b"

	ctx := context.Background()
	disksClient, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewDisksRESTClient: %w", err)
	}
	defer disksClient.Close()

	secondaryFullDiskName := fmt.Sprintf("projects/%s/zones/%s/disks/%s", projectID, zone, diskName)

	req := &computepb.StartAsyncReplicationDiskRequest{
		Project: projectID,
		Zone:    primaryZone,
		Disk:    primaryDiskName,
		DisksStartAsyncReplicationRequestResource: &computepb.DisksStartAsyncReplicationRequest{
			AsyncSecondaryDisk: &secondaryFullDiskName,
		},
	}

	op, err := disksClient.StartAsyncReplication(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to create disk: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Replication started\n")

	return nil
}

// [END compute_disk_start_replication]
