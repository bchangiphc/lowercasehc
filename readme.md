# Lowercase Traefik plugin

Lowercase is a plugin that checks the URL path for uppercase characters. If any are present the user will be redirected to the lowercase equivalent via a 301.

### Configuration

Enable the plugin in your Traefik configuration:

```
[experimental.plugins.lowercase]
  modulename = "github.com/trinnylondon/lowercase"
  version = "v0.1.0"
```

Create a Middleware. Note that this plugin does not need any configuration, however, values must be passed in for it to be accepted within Traefik.

```yaml
---
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: lowercase
spec:
  plugin:
    lowercase:
      unused: var
```
