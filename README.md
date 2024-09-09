# DNSimple module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with DNSimple.

## Caddy module name

```
dns.providers.dnsimple
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "dnsimple",
        "api_access_token": "YOUR_API_ACCESS_TOKEN"
      }
    }
  }
}
```

or with the Caddyfile:

```Caddyfile
{
	acme_dns dnsimple {$DNSIMPLE_OAUTH_TOKEN}
}

example.com {
	...
}
```

```Caddyfile
example.com {
	tls {
		dns dnsimple {$DNSIMPLE_OAUTH_TOKEN}
	}

	...
}
```

```Caddyfile
# syntax
dnsimple [<api_access_token>] {
	api_access_token <api_access_token>
}
```

## License

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
