# Reverse proxy based on a header's value 

(This code is based on the code (and takes inspiration) from: https://github.com/acouvreur/sablier)
This middleware can be used to manage the automatic deployment and scaling of services in Kubernetes, for the purpose of setting up a personal environment for the logged in user.

Compared to the original application (Sablier), this is a much more specialized, focused plugin for a very specific task.

It main purpose: based on a header field, check to see if there exists a specific service. If not, copy-n-modify a generic example service to provide the requested service.


## Configuration

### Static: 
Production:

```toml
[pilot]
  token = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

[experimental]
  [experimental.plugins]
    [experimental.plugins.my-plugin-name]
      moduleName = "github.com/dobots/service-duplicator
      version = "v0.0.2"
```

or if you're using [devMode](https://doc.traefik.io/traefik-pilot/plugins/plugin-dev/#developer-mode):

```toml
[pilot]
  token = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  
[experimental.devPlugin]
  goPath = "/plugins/go"
  moduleName = "github.com/dobots/service-duplicator"
  # Plugin will be loaded from '/plugins/go/src/github.com/dobots/service-duplicator'
```

### Dynamic:

Production:

```toml
[http]
  [http.middlewares]
    [http.middlewares.my-middleware-name.plugin.my-plugin-name]
      header  = "X-Forwarded-User"
      namespace = "myk8snamespace"

  [http.routers]
    [http.routers.my-router-name]
      entryPoints = ["websecure"]
      rule = "Host(`my-service-name.domain.tld`)"
      middlewares = ["traefik-forward-auth@docker", "my-middleware-name@file"]
      
      # if no matches are found this is the service that we forward the request to
      service = "noop@internal"
      
      [http.routers.my-router-name.tls]
        certResolver = "letsencrypt"
```

or if you're using [devMode](https://doc.traefik.io/traefik-pilot/plugins/plugin-dev/#developer-mode):

```toml
[http]
  [http.middlewares]
    [http.middlewares.my-middleware-name.plugin.dev]
      header  = "X-Forwarded-User"
      namespace = "myk8snamespace"

  [http.routers]
    [http.routers.my-router-name]
      entryPoints = ["websecure"]
      rule = "Host(`my-service-name.domain.tld`)"
      middlewares = ["traefik-forward-auth@docker", "my-middleware-name@file"]
      
      # if no matches are found this is the service that we forward the request to
      service = "noop@internal"
      
      [http.routers.my-router-name.tls]
        certResolver = "letsencrypt"
```

## License
This software is released under the Apache 2.0 License.
