/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
resource "google_dataproc_cluster" "basic" {
  name     = "dataproc-policy"
  region   = "us-central1"

  cluster_config {
    autoscaling_config {
      policy_uri = google_dataproc_autoscaling_policy.asp.name
    }
  }
}

resource "google_dataproc_autoscaling_policy" "asp" {
  policy_id = "dataproc-policy"
  location  = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"

      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}
```
