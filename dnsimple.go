package dnsimple

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	dnsimple "github.com/libdns/dnsimple"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *dnsimple.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.dnsimple",
		New: func() caddy.Module { return &Provider{new(dnsimple.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.AccountID = repl.ReplaceAll(p.Provider.AccountID, "")
	p.Provider.APIAccessToken = repl.ReplaceAll(p.Provider.APIAccessToken, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	dnsimple [<api_access_token>] {
//	    api_access_token <api_access_token>
//	    account_id <account_id>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIAccessToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_access_token":
				if p.Provider.APIAccessToken != "" {
					return d.Err("API token already set")
				}
				if d.NextArg() {
					p.Provider.APIAccessToken = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "account_id":
				if p.Provider.AccountID != "" {
					return d.Err("Account ID already set")
				}
				if d.NextArg() {
					p.Provider.AccountID = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIAccessToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
