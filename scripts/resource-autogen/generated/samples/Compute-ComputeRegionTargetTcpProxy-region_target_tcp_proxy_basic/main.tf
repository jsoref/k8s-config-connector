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
resource "google_compute_region_target_tcp_proxy" "default" {
  name            = "test-proxy"
  region          = "europe-west4"
  backend_service = google_compute_region_backend_service.default.id
}

resource "google_compute_region_backend_service" "default" {
  name        = "backend-service"
  protocol    = "TCP"
  timeout_sec = 10
  region      = "europe-west4"

  health_checks         = [google_compute_region_health_check.default.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
}

resource "google_compute_region_health_check" "default" {
  name               = "health-check"
  region             = "europe-west4"
  timeout_sec        = 1
  check_interval_sec = 1
  
  tcp_health_check {
    port = "80"
  }
}
```
