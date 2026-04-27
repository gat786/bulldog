# Bulldog

A simple CLI Golang application which you can use to scrape your kubernetes
cluster for resources you are looking for.

This tool was built with a goal to quickly take a live snapshot of the resources
and their contents by querying the kubernetes API Server and storing it in text
files. This app requires a scrapeconfig file to be provided to it and then uses
the default kubeconfig file present on an instance to fetch the specified 
resources from the cluster which is specified in the kubeconfig file.

An example scrapeconfig file looks like this

```yaml
namespaces:
  - default
resources:
  - Group: ""
    ResourceNames:
      - pods
      - serviceaccounts
    Version: v1
savefullmanifest: true
```

There are three fields which it requires to be defined in the scrapeconfig.

1. Namespaces: `namespaces`
2. Resources: `resources`
3. SaveFullManifest: `savefullmanifest`

As the names suggest these are `namespaces` from which you want the resources
to be scraped, `resources` are the actual resources that you want to be scraped
from these namespaces and the `savefullmanifest` flag actually asks if you want to
store the yaml file currently present on api server as a backup or not. By default
the scraper only lists things and will only tell you how many resources as well as
what are their names in a particular cluster or namespace.
