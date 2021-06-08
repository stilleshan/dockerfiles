# ga-proxy

[![Travis Status](https://img.shields.io/travis/com/giuem/ga-proxy.svg?style=flat-square)](https://travis-ci.com/giuem/ga-proxy)
[![Docker Build Status](https://img.shields.io/docker/build/giuem/ga-proxy.svg?style=flat-square)](https://hub.docker.com/r/giuem/ga-proxy/)
[![GitHub release](https://img.shields.io/github/release/giuem/ga-proxy.svg?style=flat-square)](https://github.com/giuem/ga-proxy/releases/latest)
[![Size](https://img.badgesize.io/https://unpkg.com/@giuem/ga-proxy/dist/ga.min.js?compression=gzip&style=flat-square)](https://unpkg.com/@giuem/ga-proxy/dist/)
[![](https://data.jsdelivr.com/v1/package/npm/@giuem/ga-proxy/badge)](https://www.jsdelivr.com/package/npm/@giuem/ga-proxy)

Accelerate Google Analytics.

## Get Start

### Run via Docker

```bash
docker pull giuem/ga-proxy
docker run -d -p <port>:80 --name <container_name> giuem/ga-proxy
```

### Run as you like

#### 1. Install 

Download binary from [release](https://github.com/giuem/ga-proxy/releases) or build yourself.

#### 2. Run

```
GIN_MODE=release ./ga-proxy [arguments]
```

options:

```
--ip IP, -i IP        IP to listen (default: "127.0.0.1") [$IP]
--port port, -p port  port to listen (default: "9080") [$PORT]
```

e.g.

```
./ga_proxy -i 0.0.0.0 -p 80
```

### 3. Insert script to your website

``` html
<script>
// replace following variables to your own
window.ga_tid = "UA-XXXXX-Y";
window.ga_url = "https://ga.giuem.com";
</script>
<script src="https://unpkg.com/@giuem/ga-proxy/dist/ga.min.js" async></script>
```

Note: `ga.giuem.com` is my own service, it do not promise any SLA and may shutdown at some day. You'd better deploy your own server.
