FROM python:3.8-alpine3.11

LABEL maintainer="Sebastian Ramirez <tiangolo@gmail.com>"

COPY install-nginx-alpine.sh /

RUN sh /install-nginx-alpine.sh

EXPOSE 80

# # Expose 443, in case of LTS / HTTPS
EXPOSE 443

# Install uWSGI
RUN apk add --no-cache uwsgi-python3

# Copy the base uWSGI ini file to enable default dynamic uwsgi process number
COPY uwsgi.ini /etc/uwsgi/

# Install Supervisord
RUN apk add --no-cache supervisor
# Custom Supervisord config
COPY supervisord-alpine.ini /etc/supervisor.d/supervisord.ini

# uWSGI Python plugin
# As an env var to re-use the config file
ENV UWSGI_PLUGIN python3

# Which uWSGI .ini file should be used, to make it customizable
ENV UWSGI_INI /app/uwsgi.ini

# By default, run 2 processes
ENV UWSGI_CHEAPER 2

# By default, when on demand, run up to 16 processes
ENV UWSGI_PROCESSES 16

# By default, allow unlimited file sizes, modify it to limit the file sizes
# To have a maximum of 1 MB (Nginx's default) change the line to:
# ENV NGINX_MAX_UPLOAD 1m
ENV NGINX_MAX_UPLOAD 0

# By default, Nginx will run a single worker process, setting it to auto
# will create a worker for each CPU core
ENV NGINX_WORKER_PROCESSES 1

# By default, Nginx listens on port 80.
# To modify this, change LISTEN_PORT environment variable.
# (in a Dockerfile or with an option for `docker run`)
ENV LISTEN_PORT 80

# Used by the entrypoint to explicitly add installed Python packages 
# and uWSGI Python packages to PYTHONPATH otherwise uWSGI can't import Flask
ENV ALPINEPYTHON python3.8

# Copy start.sh script that will check for a /app/prestart.sh script and run it before starting the app
COPY start.sh /start.sh
RUN chmod +x /start.sh

# Copy the entrypoint that will generate Nginx additional configs
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]

# Add demo app
COPY ./app /app
WORKDIR /app

# Run the start script, it will check for an /app/prestart.sh script (e.g. for migrations)
# And then will start Supervisor, which in turn will start Nginx and uWSGI
CMD ["/start.sh"]
