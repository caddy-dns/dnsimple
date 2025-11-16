# DNSimple module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with DNSimple.

## Caddy module name

```
dns.providers.dnsimple
```

## DNSimple API access tokens

An `account_id` must be provided when using this module with a DNSimple user token. The `account_id` can be omitted when using an account token since the account token already defines the account being used.

See the article on [account tokens vs user tokens](https://support.dnsimple.com/articles/api-access-token/#account-tokens-vs-user-tokens) for more information.


## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "dnsimple",
        "account_id": "YOUR_ACCOUNT_ID",
        "api_access_token": "YOUR_API_ACCESS_TOKEN"
      }
    }
  }
}
```

or with the Caddyfile:

### Global setting

```Caddyfile
{
	acme_dns dnsimple {$DNSIMPLE_API_ACCESS_TOKEN}
}

example.com {
	...
}
```

### Single site setting
```Caddyfile
example.com {
	tls {
		dns dnsimple {$DNSIMPLE_API_ACCESS_TOKEN}
	}
	...
}
```

### Provide account ID
```Caddyfile
example.com {
	tls {
		dns dnsimple {
			account_id {$DNSIMPLE_ACCOUNT_ID}
			api_access_token {$DNSIMPLE_API_ACCESS_TOKEN}
		}
	}
	...
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
