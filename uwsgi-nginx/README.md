[![Test](https://github.com/tiangolo/uwsgi-nginx-docker/workflows/Test/badge.svg)](https://github.com/tiangolo/uwsgi-nginx-docker/actions?query=workflow%3ATest) [![Deploy](https://github.com/tiangolo/uwsgi-nginx-docker/workflows/Deploy/badge.svg)](https://github.com/tiangolo/uwsgi-nginx-docker/actions?query=workflow%3ADeploy)

## Supported tags and respective `Dockerfile` links

* [`python3.8`, `latest` _(Dockerfile)_](https://github.com/tiangolo/uwsgi-nginx-docker/blob/master/docker-images/python3.8.dockerfile)
* [`python3.8-alpine` _(Dockerfile)_](https://github.com/tiangolo/uwsgi-nginx-docker/blob/master/docker-images/python3.8-alpine.dockerfile)
* [`python3.7`, _(Dockerfile)_](https://github.com/tiangolo/uwsgi-nginx-docker/blob/master/docker-images/python3.7.dockerfile)
* [`python3.6` _(Dockerfile)_](https://github.com/tiangolo/uwsgi-nginx-docker/blob/master/docker-images/python3.6.dockerfile)
* [`python2.7` _(Dockerfile)_](https://github.com/tiangolo/uwsgi-nginx-docker/blob/master/docker-images/python2.7.dockerfile)

**Note**: Note: There are [tags for each build date](https://hub.docker.com/r/tiangolo/uwsgi-nginx/tags). If you need to "pin" the Docker image version you use, you can select one of those tags. E.g. `tiangolo/uwsgi-nginx:python3.7-2019-09-28`.

# uwsgi-nginx

**Docker** image with **uWSGI** and **Nginx** for web applications in **Python 3.6** and above, and **Python 2.7** (as **Flask**) in a single container. Optionally with Alpine Linux.

## Description

This [**Docker**](https://www.docker.com/) image allows you to create [**Python**](https://www.python.org/) web applications that run with [**uWSGI**](https://uwsgi-docs.readthedocs.org/en/latest/) and [**Nginx**](http://nginx.org/en/) in a single container.

The combination of uWSGI with Nginx is a [common way to deploy Python web applications like Flask and Django](http://flask.pocoo.org/docs/1.0/deploying/uwsgi/). It is widely used in the industry and would give you decent performance. (*)

There is also an Alpine version. If you want it, check the tags from above.

This image was created to be the base image for [**tiangolo/uwsgi-nginx-flask**](https://hub.docker.com/r/tiangolo/uwsgi-nginx-flask/) but could be used as the base image for any other (WSGI-based) Python web application, like Django.

### * Note on performance and features

If you are starting a new project, you might benefit from a newer and faster framework based on ASGI instead of WSGI (Flask and Django are WSGI-based).

You could use an ASGI framework like:

* [**FastAPI**](https://github.com/tiangolo/fastapi) (which is based on Starlette) with this Docker image: [**tiangolo/uvicorn-gunicorn-fastapi**](https://github.com/tiangolo/uvicorn-gunicorn-fastapi-docker).
* [**Starlette**](https://github.com/encode/starlette) directly, with this Docker image: [**tiangolo/uvicorn-gunicorn-starlette**](https://github.com/tiangolo/uvicorn-gunicorn-starlette-docker).
* Or any other ASGI framework with this Docker image: [**tiangolo/uvicorn-gunicorn**](https://github.com/tiangolo/uvicorn-gunicorn-docker).

FastAPI, or Starlette, would give you about 800% (8x) the performance achievable with this image (**tiangolo/uwsgi-nginx**). [You can see the third-party benchmarks here](https://www.techempower.com/benchmarks/#section=test&runid=a979de55-980d-4721-a46f-77298b3f3923&hw=ph&test=query&l=z8kflr-v&a=2).

Also, if you want to use new technologies like WebSockets it would be easier (and *possible*) with a newer framework based on ASGI, like FastAPI or Starlette. As the standard ASGI was designed to be able to handle asynchronous code like the one needed for WebSockets.

#### If you need a WSGI-based application (like Flask or Django)

If you need to use an older WSGI-based framework like Flask or Django (instead of something based on ASGI) and you need to have the best performance possible, you can use the alternative image: [**tiangolo/meinheld-gunicorn**](https://github.com/tiangolo/meinheld-gunicorn-docker).

**tiangolo/meinheld-gunicorn** will give you about 400% (4x) the performance of this image.

---

**GitHub repo**: [https://github.com/tiangolo/uwsgi-nginx-docker](https://github.com/tiangolo/uwsgi-nginx-docker)

**Docker Hub image**: [https://hub.docker.com/r/tiangolo/uwsgi-nginx/](https://hub.docker.com/r/tiangolo/uwsgi-nginx/)

## How to use

* You shouldn't have to clone the GitHub repo. You should use it as a base image for other images, using this in your `Dockerfile`:

```Dockerfile
FROM tiangolo/uwsgi-nginx:python3.8

# Your Dockerfile code...
```

* But, if you need Python 2.7 that line would have to be `FROM tiangolo/uwsgi-nginx:python2.7`.

* By default it will try to find a uWSGI config file in `/app/uwsgi.ini`.

* That `uwsgi.ini` file will make it try to run a Python file in `/app/main.py`.

If you are building a **Flask** web application you should use instead [**tiangolo/uwsgi-nginx-flask**](https://hub.docker.com/r/tiangolo/uwsgi-nginx-flask/).

## Advanced usage

### Custom app directory

If you need to use a directory for your app different than `/app`, you can override the uWSGI config file path with an environment variable `UWSGI_INI`, and put your custom `uwsgi.ini` file there.

For example, if you needed to have your application directory in `/application` instead of `/app`, your `Dockerfile` would look like:

```Dockerfile
FROM tiangolo/uwsgi-nginx:python3.8

ENV UWSGI_INI /application/uwsgi.ini

COPY ./application /application
WORKDIR /appapplication
```

And your `uwsgi.ini` file in `./application/uwsgi.ini` would contain:

```ini
[uwsgi]
wsgi-file=/application/main.py
```

**Note**: it's important to include the `WORKDIR` option, otherwise uWSGI will start the application in `/app`.

### Custom uWSGI process number

By default, the image starts with 2 uWSGI processes running. When the server is experiencing a high load, it creates up to 16 uWSGI processes to handle it on demand.

If you need to configure these numbers you can use environment variables.

The starting number of uWSGI processes is controlled by the variable `UWSGI_CHEAPER`, by default set to `2`.

The maximum number of uWSGI processes is controlled by the variable `UWSGI_PROCESSES`, by default set to `16`.

Have in mind that `UWSGI_CHEAPER` must be lower than `UWSGI_PROCESSES`.

So, if, for example, you need to start with 4 processes and grow to a maximum of 64, your `Dockerfile` could look like:

```Dockerfile
FROM tiangolo/uwsgi-nginx:python3.8

ENV UWSGI_CHEAPER 4
ENV UWSGI_PROCESSES 64

COPY ./app /app
```

### Custom max upload size

In this image, Nginx is configured to allow unlimited upload file sizes. This is done because by default a simple Python server would allow that, so that's the simplest behavior a developer would expect.

If you need to restrict the maximum upload size in Nginx, you can add an environment variable `NGINX_MAX_UPLOAD` and assign a value corresponding to the [standard Nginx config `client_max_body_size`](http://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size).

For example, if you wanted to set the maximum upload file size to 1 MB (the default in a normal Nginx installation), you would need to set the `NGINX_MAX_UPLOAD` environment variable to the value `1m`. Then the image would take care of adding the corresponding configuration file (this is done by the `entrypoint.sh`).

So, your `Dockerfile` would look something like:

```Dockerfile
FROM tiangolo/uwsgi-nginx:python3.8

ENV NGINX_MAX_UPLOAD 1m

COPY ./app /app
```

### Custom listen port

By default, the container made from this image will listen on port 80.

To change this behavior, set the `LISTEN_PORT` environment variable.

You might also need to create the respective `EXPOSE` Docker instruction.

You can do that in your `Dockerfile`, it would look something like:

```Dockerfile
FROM tiangolo/uwsgi-nginx:python3.8

ENV LISTEN_PORT 8080

EXPOSE 8080

COPY ./app /app
```

### Custom `/app/prestart.sh`

If you need to run anything before starting the app, you can add a file `prestart.sh` to the directory `/app`. The image will automatically detect and run it before starting everything.

For example, if you want to add database migrations that are run on startup (e.g. with Alembic, or Django migrations), before starting the app, you could create a `./app/prestart.sh` file in your code directory (that will be copied by your `Dockerfile`) with:

```bash
#! /usr/bin/env bash

# Let the DB start
sleep 10;
# Run migrations
alembic upgrade head
```

and it would wait 10 seconds to give the database some time to start and then run that `alembic` command (you could update that to run Django migrations or any other tool you need).

If you need to run a Python script before starting the app, you could make the `/app/prestart.sh` file run your Python script, with something like:

```bash
#! /usr/bin/env bash

# Run custom Python script before starting
python /app/my_custom_prestart_script.py
```

**Note**: The image uses `.` to run the script (as in `. /app/prestart.sh`), so for example, environment variables would persist. If you don't understand the previous sentence, you probably don't need it.

### Custom Nginx processes number

By default, Nginx will start one "worker process".

If you want to set a different number of Nginx worker processes you can use the environment variable `NGINX_WORKER_PROCESSES`.

You can use a specific single number, e.g.:

```Dockerfile
ENV NGINX_WORKER_PROCESSES 2
```

or you can set it to the keyword `auto` and it will try to autodetect the number of CPUs available and use that for the number of workers.

For example, using `auto`, your Dockerfile could look like:

```Dockerfile
FROM tiangolo/uwsgi-nginx:python3.8

ENV NGINX_WORKER_PROCESSES auto

COPY ./app /app
```

### Custom Nginx maximum connections per worker

By default, Nginx will start with a maximum limit of 1024 connections per worker.

If you want to set a different number you can use the environment variable `NGINX_WORKER_CONNECTIONS`, e.g:

```Dockerfile
ENV NGINX_WORKER_CONNECTIONS 2048
```

It cannot exceed the current limit on the maximum number of open files. See how to configure it in the next section.

### Custom Nginx maximum open files

The number connections per Nginx worker cannot exceed the limit on the maximum number of open files.

You can change the limit of open files with the environment variable `NGINX_WORKER_OPEN_FILES`, e.g.:

```Dockerfile
ENV NGINX_WORKER_OPEN_FILES 2048
```

### Customizing Nginx additional configurations

If you need to configure Nginx further, you can add `*.conf` files to `/etc/nginx/conf.d/` in your `Dockerfile`.

Just have in mind that the default configurations are created during startup in a file at `/etc/nginx/conf.d/nginx.conf` and `/etc/nginx/conf.d/upload.conf`. So you shouldn't overwrite them. You should name your `*.conf` file with something different than `nginx.conf` or `upload.conf`, for example: `custom.conf`.

**Note**: if you are customizing Nginx, maybe copying configurations from a blog or a StackOverflow answer, have in mind that you probably need to use the [configurations specific to uWSGI](http://nginx.org/en/docs/http/ngx_http_uwsgi_module.html), instead of those for other modules, like for example, `ngx_http_fastcgi_module`.

### Overriding Nginx configuration completely

If you need to configure Nginx even further, completely overriding the defaults, you can add a custom Nginx configuration to `/app/nginx.conf`.

It will be copied to `/etc/nginx/nginx.conf` and used instead of the generated one.

Have in mind that, in that case, this image won't generate any of the Nginx configurations, it will only copy and use your configuration file.

That means that all the environment variables described above that are specific to Nginx won't be used.

It also means that it won't use additional configurations from files in `/etc/nginx/conf.d/*.conf`, unless you explicitly have a section in your custom file `/app/nginx.conf` with:

```conf
include /etc/nginx/conf.d/*.conf;
```

If you want to add a custom `/app/nginx.conf` file but don't know where to start from, you can use [the `nginx.conf` used for the tests](https://github.com/tiangolo/uwsgi-nginx-docker/blob/master/tests/test_02_app/custom_nginx_app/app/nginx.conf) and customize it or modify it further.

## Technical details

The combination of uWSGI with Nginx is a [common way to deploy Python web applications](http://flask.pocoo.org/docs/1.0/deploying/uwsgi/).

Roughly:

* **Nginx** is a web server, it takes care of the HTTP connections and also can serve static files directly and more efficiently.

* **uWSGI** is an application server, that's what runs your Python code and it talks with Nginx.

* **Your Python code** has the actual web application, and is run by uWSGI.

This image takes advantage of already slim and optimized existing Docker images (based on Debian as [recommended by Docker](https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/)) and implements Docker best practices.

It uses the official Python Docker image, installs uWSGI and on top of that, with the least amount of modifications, adds the official Nginx image (as of 2016-02-14).

And it controls all these processes with Supervisord.

---

There's the rule of thumb that you should have "one process per container".

That helps, for example, isolating an app and its database in different containers.

But if you want to have a "micro-services" approach you may want to [have more than one process in one container](https://valdhaus.co/writings/docker-misconceptions/) if they are all related to the same "service", and you may want to include your Flask code, uWSGI and Nginx in the same container (and maybe run another container with your database).

That's the approach taken in this image.

---

This image has a default sample "Hello World" app in the container's `/app` directory using the example in the [uWSGI documentation](http://uwsgi-docs.readthedocs.org/en/latest/WSGIquickstart.html).

You probably want to override it or delete it in your project.

It is there in case you run this image by itself and not as a base image for your own `Dockerfile`, so that you get a sample app without errors.

## Tests

All the image tags, configurations, environment variables and application options are tested.

## Updates

Updates are announced in the releases.

You can click the "watch" button at the top right and select "Releases only" to receive an email notification when there's a new release.

## Release Notes

### Latest Changes

* 🐛 Fix broken link to TechEmpower benchmarks. PR [#96](https://github.com/tiangolo/uwsgi-nginx-docker/pull/96) by [@tiangolo](https://github.com/tiangolo).
* 👷 Add GitHub Action latest-changes, update issue-manager. PR [#92](https://github.com/tiangolo/uwsgi-nginx-docker/pull/92) by [@tiangolo](https://github.com/tiangolo).
* Fix Python 3.8 Alpine environment for installed packages. PR [#84](https://github.com/tiangolo/uwsgi-nginx-docker/pull/84).

### 1.4.0

* Add [GitHub Sponsors](https://github.com/sponsors/tiangolo) button.
* Add new image for Python 3.8, and new image for Python 3.8 on Alpine. PR [#83](https://github.com/tiangolo/uwsgi-nginx-docker/pull/83).
* Upgrade Nginx to latest version, `1.17.10`, based on latest Debian, Buster. PR [#82](https://github.com/tiangolo/uwsgi-nginx-docker/pull/82).
* Remove support for Python 3.5. PR [#81](https://github.com/tiangolo/uwsgi-nginx-docker/pull/81).

### 1.3.0

* This is the last version to support:
    * Debian Stretch (before upgrading to Buster).
    * Python 3.5.
    * Alpine 3.7, 3.8, 3.9 (before upgrading to Alpine 3.11).
    * Alpine in older versions of Python, 2.7 and 3.6 (Before upgrading to Python 3.8).
    * If you need any of those, make sure to use a tag for the build date `2020-05-04`.
* Refactor build set up:
    * Re-use code and configs.
    * Migrate to GitHub Actions.
    * Simplify tests.
    * PR [#78](https://github.com/tiangolo/uwsgi-nginx-docker/pull/78).
* Migrate Travis to .com, update badge. PR [#77](https://github.com/tiangolo/uwsgi-nginx-docker/pull/77).

### 1.2.0

* 2019-10-14:
    * Refactor and simplify test scripts. PR [#66](https://github.com/tiangolo/uwsgi-nginx-docker/pull/66).
* 2019-09-28:
    * Refactor build scripts and add image tags for each build date, like `tiangolo/uwsgi-nginx:python3.7-2019-09-28`. PR [#65](https://github.com/tiangolo/uwsgi-nginx-docker/pull/65).

* Upgrade Travis. PR [#60](https://github.com/tiangolo/uwsgi-nginx-docker/pull/60).

### 1.1.0

* Added support for `/app/prestart.sh` script to run arbitrary code before starting the app (for example, Alembic - SQLAlchemy migrations). The [documentation for the `/app/prestart.sh` is in the main README](https://github.com/tiangolo/uwsgi-nginx-docker#custom-appprestartsh). [PR #59](https://github.com/tiangolo/uwsgi-nginx-docker/pull/59).

### 1.0.0

* 2019-05-04:
    * Add Alpine Linux 3.9. PR [#55](https://github.com/tiangolo/uwsgi-nginx-docker/pull/55) by [evilgoldfish](https://github.com/evilgoldfish).
    * Build images using Travis matrix to improve development/testing speed. Needed for some recent PRs. [PR #58](https://github.com/tiangolo/uwsgi-nginx-docker/pull/58).

* 2019-02-02:
    * The Nginx configurations are generated dynamically from the entrypoint, instead of modifying pre-existing files. [PR #50](https://github.com/tiangolo/uwsgi-nginx-docker/pull/50).
    * Support for a completely custom `/app/nginx.conf` file that overrides the generated one. [PR #51](https://github.com/tiangolo/uwsgi-nginx-docker/pull/51).

* 2018-11-23: New Alpine 3.8 images for Python 2.7, Python 3.6 and Python 3.7 (Python 3.7 temporarily disabled). Thanks to [philippfreyer](https://github.com/philippfreyer) in [PR #45](https://github.com/tiangolo/uwsgi-nginx-docker/pull/45)

* 2018-09-22: New Python 3.7 versions, standard and Alpine based. Thanks to [desaintmartin](https://github.com/desaintmartin) in [this PR](https://github.com/tiangolo/uwsgi-nginx-docker/pull/39).

* 2018-06-22: You can now use `NGINX_WORKER_CONNECTIONS` to set the maximum number of Nginx worker connections and `NGINX_WORKER_OPEN_FILES` to set the maximum number of open files. Thanks to [ronlut](https://github.com/ronlut) in [this PR](https://github.com/tiangolo/uwsgi-nginx-docker/pull/26).

* 2018-06-22: Make uWSGI require an app to run, instead of going in "full dynamic mode" while there was an error. Supervisord doesn't terminate itself but tries to restart uWSGI and shows the errors. Uses `need-app` as suggested by [luckydonald](https://github.com/luckydonald) in [this comment](https://github.com/tiangolo/uwsgi-nginx-flask-docker/issues/3#issuecomment-321991279).

* 2018-06-22: Correctly handled graceful shutdown of uWSGI and Nginx. Thanks to [desaintmartin](https://github.com/desaintmartin) in [this PR](https://github.com/tiangolo/uwsgi-nginx-docker/pull/30).

* 2018-02-04: It's now possible to set the number of Nginx worker processes with the environment variable `NGINX_WORKER_PROCESSES`. Thanks to [naktinis](https://github.com/naktinis) in [this PR](https://github.com/tiangolo/uwsgi-nginx-docker/pull/22).

* 2018-01-14: There are now two Alpine based versions, `python2.7-alpine3.7` and `python3.6-alpine3.7`.

* 2017-12-08: Now you can configure which port the container should listen on, using the environment variable `LISTEN_PORT` thanks to [tmshn](https://github.com/tmshn) in [this PR](https://github.com/tiangolo/uwsgi-nginx-docker/pull/16).

* 2017-08-09: You can set a custom maximum upload file size using an environment variable `NGINX_MAX_UPLOAD`, by default it has a value of `0`, that allows unlimited upload file sizes. This differs from Nginx's default value of 1 MB. It's configured this way because that's the simplest experience a developer that is not expert in Nginx would expect.

* 2017-08-09: Now you can override where to look for the `uwsgi.ini` file, and with that, change the default directory from `/app` to something else, using the envirnoment variable `UWSGI_INI`.

* 2017-08-08: There's a new `latest` tag image, just to show a warning for those still using `latest` for Python 2.7 web applications. As of now, [everyone](https://www.python.org/dev/peps/pep-0373/) [should be](http://flask.pocoo.org/docs/0.12/python3/#python3-support) [using Python 3](https://docs.djangoproject.com/en/1.11/faq/install/#what-python-version-should-i-use-with-django).

* 2017-08-08: Supervisord now terminates uWSGI on `SIGTERM`, so if you run `docker stop` or something similar, it will actually stop everything, instead of waiting for Docker's timeout to kill the container.

* 2017-07-31: There's now an image tag for Python 3.6, based on the official image for Python 3.6 thanks to [jrd](https://github.com/jrd) in [this PR](https://github.com/tiangolo/uwsgi-nginx-docker/pull/6).

* 2016-10-01: Now you can override default `uwsgi.ini` parameters from the file in `/app/uwsgi.ini`.

* 2016-08-16: There's now an image tag for Python 3.5, based on the official image for Python 3.5. So now you can use this image for your projects in Python 2.7 and Python 3.5.

* 2016-08-16: Use dynamic a number of worker processes for uWSGI, from 2 to 16 depending on load. This should work for most cases. This helps especially when there are some responses that are slow and take some time to be generated, this change allows all the other responses to keep fast (in a new process) without having to wait for the first (slow) one to finish.

* Also, it now uses a base `uwsgi.ini` file under `/etc/uwsgi/` with most of the general configurations, so, the `uwsgi.ini` inside `/app` (the one you could need to modify) is now a lot simpler.

* 2016-04-05: Nginx and uWSGI logs are now redirected to stdout, allowing to use `docker logs`.

## License

This project is licensed under the terms of the Apache license.
