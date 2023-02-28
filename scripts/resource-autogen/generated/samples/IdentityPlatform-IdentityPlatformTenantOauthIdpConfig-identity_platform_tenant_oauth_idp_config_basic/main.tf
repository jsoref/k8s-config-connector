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
resource "google_identity_platform_tenant" "tenant" {
  display_name  = "tenant"
}

resource "google_identity_platform_tenant_oauth_idp_config" "tenant_oauth_idp_config" {
  name          = "oidc.oauth-idp-config"
  tenant        = google_identity_platform_tenant.tenant.name
  display_name  = "Display Name"
  client_id     = "client-id"
  issuer        = "issuer"
  enabled       = true
  client_secret = "secret"
}
```
