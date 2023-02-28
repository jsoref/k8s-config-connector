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
resource "random_id" "url_signature" {
  byte_length = 16
}

resource "google_compute_backend_bucket_signed_url_key" "backend_key" {
  name           = "test-key"
  key_value      = random_id.url_signature.b64_url
  backend_bucket = google_compute_backend_bucket.test_backend.name
}

resource "google_compute_backend_bucket" "test_backend" {
  name        = "test-signed-backend-bucket"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "bucket" {
  name     = "test-storage-bucket"
  location = "EU"
}
```
