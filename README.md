<h1 align="center">
  Dockerfile From Image (dfimage)
  <br>
</h1>

<p align="center">
<a href="https://goreportcard.com/report/github.com/zer0yu/dfimage"><img src="https://goreportcard.com/badge/github.com/zer0yu/dfimage"></a>
<a href="https://github.com/zer0yu/dfimage/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://github.com/zer0yu/dfimage/releases"><img src="https://img.shields.io/github.com/zer0yu/dfimage"></a>
</p>

<h4 align="center">Reverse-engineers a Dockerfile from a Docker image.</h4>

<p align="center">
  <a href="#usage">Usage</a> •
  <a href="#installation">Install</a> •
  <a href="#docker-example">Docker Example</a> •
  <a href="#references">References</a> •
  <a href="#license">License</a>
</p>


## usage

```sh
# dfimage  <image>
dfimage ruby:latest
```

## installation

Run the following command to install the latest version:

```sh
go install -v github.com/zer0yu/dfimage/main@latest
```

or you can download the binary from [releases](http://github.com/zer0yu/dfimage/releases) 

## docker-example

```sh
❯ docker pull ruby:latest

❯ dfimage ruby:latest
FROM ruby:latest
ADD file:3e9b6405f11dd24ce62105c033f1d8b931d9409298553f63b03af1b6dd1dda35 in /
CMD ["bash"]
...
ENV GEM_HOME=/usr/local/bundle
ENV BUNDLE_SILENCE_ROOT_WARNING=1 BUNDLE_APP_CONFIG=/usr/local/bundle
ENV PATH=/usr/local/bundle/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
RUN /bin/sh -c mkdir -p "$GEM_HOME" \
    && chmod 1777 "$GEM_HOME"
CMD ["irb"]

```

## References

1. [dfimage - python](https://github.com/LanikSJ/dfimage/)
2. [Inspiration](https://github.com/CenturyLinkLabs/dockerfile-from-image)

## License

[![MIT License](https://img.shields.io/badge/license-MIT-blue)](https://en.wikipedia.org/wiki/MIT_License)