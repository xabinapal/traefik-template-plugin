# Traefik Template Plugin

<div align="center">

  <img src="./.assets/icon.svg" width="128" alt="Traefik Template Plugin">

</div>

<div align="center">

**Made with ❤️ for the Traefik community**
<br>
⭐ [Star repo](https://github.com/xabinapal/traefik-template-plugin/stargazers)
🐛 [Report bug](https://github.com/xabinapal/traefik-template-plugin/issues)
💡 [Request feature](https://github.com/xabinapal/traefik-template-plugin/issues)

</div>

<div align="center">

**If this plugin helps you, consider supporting me**
<br>
[![Ko-fi](https://img.shields.io/badge/Ko--fi-Support%20me-ff5e5b?logo=ko-fi&logoColor=white)](https://ko-fi.com/xabinapal)

</div>

## What is this?

Traefik Template Plugin is an example middleware plugin for [Traefik](https://traefik.io/) that demonstrates how to add custom headers to requests forwarded to upstream services. It is designed to be a simple starting point for building your own Traefik plugins.

## Installation

Add the plugin to your Traefik static configuration using the experimental plugins feature:

**File**

```yaml
experimental:
  plugins:
    template:
      moduleName: "github.com/xabinapal/traefik-template-plugin"
      version: "v1.0.0"
```

**CLI**

```sh
--experimental.plugins.template.modulename=github.com/xabinapal/traefik-template-plugin
--experimental.plugins.template.version=v1.0.0
```

## Configuration

- `headers`: `map[string]string`, optional, default `{}` \
  A key-value map of headers to inject into requests forwarded to upstream services.

## Examples

### File YAML Provider

```yaml
http:
  middlewares:
    template:
      plugin:
        template:
          headers:
            X-Demo: test
            X-Url:  http://localhost

  routers:
    api:
      rule: Host(`api.example.com`)
      middlewares:
        - template
      service: api-service
```

### Kubernetes CRD Provider

**Middleware**

```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: template
spec:
  plugin:
    template:
      headers:
        X-Demo: test
        X-Url:  http://localhost
```

**IngressRoute**

```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: api-route
spec:
  entryPoints:
    - web
  routes:
    - match: Host(`api.example.com`)
      kind: Rule
      middlewares:
        - name: template
      services:
        - name: api-service
          port: 80
```

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details
