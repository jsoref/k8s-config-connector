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
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "secret-name"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "secret-data"
}

resource "google_network_services_edge_cache_origin" "default" {
  name           = "my-origin"
  origin_address = "gs://media-edge-default"
  description    = "The default bucket for V4 authentication"
  aws_v4_authentication {
    access_key_id             = "ACCESSKEYID"
    secret_access_key_version = google_secret_manager_secret_version.secret-version-basic.id
    origin_region             = "auto"
  }
}
```
