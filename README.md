# Bulldog

A simple CLI Golang application which you can use to scrape your kubernetes
cluster for resources you are looking for.

This tool was built with a goal to quickly take a live snapshot of the resources
and their contents by querying the kubernetes API Server and storing it in text
files. This app requires a scrapeconfig file to be provided to it and then uses
the default kubeconfig file present on an instance to fetch the specified 
resources from the cluster which is specified in the kubeconfig file.

## Using bulldog

You can either download the source code and build the application on your own 
machine or download the prebuilt binaries present on github releases and then
execute them. Building the binary is as simple as 

```sh
git clone https://github.com/gat786/bulldog
cd app && go build . -o bulldog
```

If you decide to download a prebuilt binary you would either need to add it to
path (install) or execute it from the directory in which it is downloaded. 

You can install the downloaded binary like below 

```sh
# change the binary name according to your downloaded choice
mv bulldog-linux-amd64 bulldog
sudo install -m 755 bulldog /usr/local/bin
```

Doing this would add the binary to an executable directory and make it such that 
it will then be available to be executed from terminal automatically.

Executing the Bulldog CLI would then show you the supported options

```sh
$ bulldog

2026/04/27 18:26:59 info Initialising Kubernetes Resource Scraper
NAME:
   bulldog - A new cli application

USAGE:
   bulldog [global options] [command [command options]]

COMMANDS:
   scrape
   print-config
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

There are two options as of now, 

1. scrape

    When ran the Scrape option will read the default kubeconfig file present 
    and start scraping the resources present on that cluster and saving it as a
    text output in the `exported_data` directory.


> !TODO - configure a way to export to a custom user specified directory


2. print-config
  
    Inorder to scrape data from a cluster, this tool requires a scrapeconfig, 
    which if not specified manually, the tool considers that you just want 
    everything to scraped and stored. You can provide a custom scrapeconfig by 
    specifying the `CONFIG_FILE_PATH` file path which when specified the yaml 
    file on that path will be used as the scrapeconfig.

    If you are not sure what items the tool will scrape and want to see what 
    config the tool is using you can run this command along with the environment
    variable specified or not and the tool will print out the config it is 
    using to scrape the data.

## Scrapeconfig

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
