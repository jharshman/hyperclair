

# hyperclair

[![Build Status](https://travis-ci.org/jharshman/hyperclair.svg?branch=master)](https://travis-ci.org/jharshman/hyperclair)
> Tracking container vulnerabilities, that's should be *Hyperclair*

Tracking vulnerabilities in your container images, it's easy with CoreOS Clair.
Integrate it inside your CI/CD pipeline is easier with Hyperclair.

Hyperclair is a lightweight command-line tool doing the bridge between Registries as Docker Hub, Docker Registry or Quay.io, and the CoreOS vulnerability tracker, Clair.
It's easily integrated in your CI/CD pipeline and Hyperclair will play as reverse proxy for authentication.


> The Registry is a stateless, highly scalable server side application that stores and lets you distribute Docker images. The Registry is open-source, under the permissive Apache license.
>
>*From https://docs.docker.com/registry/*

> Clair is a container vulnerability analysis service. It provides a list of vulnerabilities that threaten a container, and can notify users when new vulnerabilities that affect existing containers become known.
>
>*From https://github.com/coreos/clair*

hyperclair is tool to make the link between the Docker Registry and the CoreOS Clair tool.

![hyperclair](https://cloud.githubusercontent.com/assets/3304363/14174675/348bc190-f746-11e5-9edd-9e736ec38b0e.png)

# Installation

```bash
# sudo curl -L -o /usr/local/bin/hyperclair  https://github.com/wemanity-belgium/hyperclair/releases/download/0.5.0/hyperclair-{OS}-{ARCH}
sudo curl -L -o /usr/local/bin/hyperclair  https://github.com/wemanity-belgium/hyperclair/releases/download/0.5.0/hyperclair-linux-amd64
sudo chmod +x /usr/local/bin/hyperclair
```

# Usage

[![asciicast](https://asciinema.org/a/41461.png)](https://asciinema.org/a/41461)

# Notification
2. On-Demand: the CLI tool is used to pull image from Registry then push it to Clair

# Reporting

**hyperclair** get vulnerabilities report from Clair and generate HTML report

See [Sample](http://htmlpreview.github.io/?https://github.com/wemanity-belgium/hyperclair/blob/master/report-sample.html)

hyperclair can be used for Docker Hub and self-hosted Registry

# Command

```
Analyse your docker image with Clair, directly from your registry.

Usage:
  hyperclair [command]

Available Commands:
  analyse     Analyse Docker image
  health      Get Health of Hyperclair and underlying services
  login       Log in to a Docker registry
  logout      Log out from a Docker registry
  pull        Pull Docker image information
  push        Push Docker image to Clair
  report      Generate Docker Image vulnerabilities report
  version     Get Versions of Hyperclair and underlying services

Flags:
      --config string      config file (default is ./.hyperclair.yml)
      --log-level string   log level [Panic,Fatal,Error,Warn,Info,Debug]

Use "hyperclair [command] --help" for more information about a command.
```

# Optional Configuration

```yaml
clair:
  port: 6060
  healthPort: 6061
  uri: http://clair
  priority: Low
  report:
    path: ./reports
    format: html
```

# Contribution and Test

Go to /contrib folder
