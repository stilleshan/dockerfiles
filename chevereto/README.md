[cheveretourl]: https://chevereto.com/
[cheveretogithub]: https://github.com/Chevereto/Chevereto-Free

[![chevereto](http://chevereto.com/app/themes/v3/img/chevereto-blue.svg)][cheveretourl]

# nmtan/chevereto - Chevereto Free Docker

[Chevereto][cheveretourl] is a powerful and fast image hosting script that allows you to create your very own full featured image hosting website in just minutes.

Please note that the `installer` tag will provide you with the installer script
from Chevereto, which can be used to install either the [free Chevereto version][cheveretogithub]
or the paid version, while all the other tags will only provide the [free version][cheveretogithub].

## Supported tags and respective Dockerfile links

> Here we list only the most relevant and recent tags, for the full list, please
use the `tags` tab in [Dockerhub page](https://hub.docker.com/r/nmtan/chevereto)

* `latest` - Using latest commit of the branch `master` from [original repo][cheveretogithub] ([Dockerfile](https://github.com/tanmng/docker-chevereto/blob/master/latest/Dockerfile))
* `installer` - Using latest [installer script](https://cdn.rawgit.com/Chevereto/php-repo-installer/master/index.php) ([Dockerfile-installer](https://github.com/tanmng/docker-chevereto/blob/master/Dockerfile-installer)) - Once you start the container with this image, you will have to wait for it to download latest source code from [original repo][cheveretogithub]
* `1.3.0` - Using the [`1.3.0` release](https://github.com/Chevereto/Chevereto-Free/releases/tag/1.3.0) ([Dockerfile](https://github.com/tanmng/docker-chevereto/blob/master/Dockerfile))
* `1.2.3` - Using the [`1.2.3` release](https://github.com/Chevereto/Chevereto-Free/releases/tag/1.2.3) ([Dockerfile](https://github.com/tanmng/docker-chevereto/blob/master/Dockerfile))
* `1.2.2` - Using the [`1.2.2` release](https://github.com/Chevereto/Chevereto-Free/releases/tag/1.2.2) ([Dockerfile](https://github.com/tanmng/docker-chevereto/blob/master/Dockerfile))
* `1.2.1` - Using the [`1.2.1` release](https://github.com/Chevereto/Chevereto-Free/releases/tag/1.2.1) ([Dockerfile](https://github.com/tanmng/docker-chevereto/blob/master/Dockerfile))
* `1.2.0` - Using the [`1.2.0` release](https://github.com/Chevereto/Chevereto-Free/releases/tag/1.2.0) ([Dockerfile](https://github.com/tanmng/docker-chevereto/blob/master/Dockerfile))
* `1.1.4` - Using the [`1.1.4` release](https://github.com/Chevereto/Chevereto-Free/releases/tag/1.1.4) ([Dockerfile](https://github.com/tanmng/docker-chevereto/blob/master/Dockerfile))

> Note on PHP version: [Chevereto](cheveretourl) added support for PHP 7.4 since release `1.3.0`, PHP 7.3 since release `1.2.0`, PHP 7.2 since release `1.1.0`, and support for PHP 7.1 since release `1.0.6`, thus, the Docker images use the appropriate version of PHP to ensure the best performance and stability. You can check the [images' labels](https://docs.docker.com/config/labels-custom-metadata/) (by running `docker image inspect IMAGE_NAME`) for this information.

> Some older tag of Chevereto are no longer supported to save up on build resources

## Environment variables

The most essentials environments variables are listed below

* `CHEVERETO_DB_HOST` - Hostname of the Database machine that you wish to connect, default to `db`
* `CHEVERETO_DB_PORT` - The port of the Database machine to connect to, default to `3306`
* `CHEVERETO_DB_USERNAME` - Username to authenticate to MySQL database, default to `chevereto`
* `CHEVERETO_DB_PASSWORD` - Password of the user when connect to MySQL database, default to `chevereto`
* `CHEVERETO_DB_NAME` - Name of the database in MySQL server, default to `chevereto`
* `CHEVERETO_DB_PREFIX` - Table prefix (you can use this to run multiple instance of Chevereto using the same Database), default to `chv_`

> For other environment variables, please consult the file [`settings.php`](https://github.com/tanmng/docker-chevereto/blob/master/settings.php) and the section "Advanced configuration" below.

## Connection to database

[Chevereto][cheveretourl] requires an Mysql database to store its information. You can use a [Mysql](https://hub.docker.com/_/mysql/) or [MariaDB](https://hub.docker.com/_/mariadb/) container to host this.

Information on connection to database is provided to container via environment variables explained above.

## Persistent storage

[Chevereto][cheveretourl] stores images uploaded by users in `/var/www/html/images` directory within the container.

You can mount a [data volume](https://docs.docker.com/engine/tutorials/dockervolumes/#data-volumes) at this location to ensure that you don't lose your images if you relaunch/remove container.

## Max image size

By default, PHP allow a maximum file upload to be 2MB. You can change such behaviour by updating the `php.ini` in your container, either by bind-mount the file, or build a new image with the updated file, that way you can reuse the image on demand.

> Note that by default, Chevereto set a file upload limit of 10MB, so after you modify your `php.ini`, you should also update this settings in Chevereto settings page (available at CHEVERETO_URL/dashboard/settings/image-upload)

> The customized `php.ini` should set the values of `upload_max_filesize`, `post_max_size` and potentially `memory_limit`, as showed in [the discussion from Chevereto Forum](https://chevereto.com/community/threads/chevereto-supports-only-2mb-max-upload-size.4729/). Further details on these parameters are available from [PHP documentation](http://php.net/manual/en/ini.core.php)

An example of this is available in the [`examples/bigger-files` directory](examples/bigger-files)

## Advanced configuration

Chevereto is written using the [`G\` framework](https://g.chevereto.com/) which supports some [advanced configurations](https://github.com/Chevereto/Chevereto-Free/blob/master/lib/G/G.php). Ultimate, that means we can use the variable `$settings` from `settings.php` to add more configurations to the applications. I have added multiple environment variables to support these configurations, please consult [the file `settings.php` from the repo](https://github.c m/tanmng/docker-chevereto/blob/master/settings.php) for the list of these environment variables as well as their purpose. Among these settings, I would say the following two are the most useful

* `CHEVERETO_SESSION_SAVE_PATH` - Specify where Chevereto can save the user session. Combining this with a shared volume/EFS/File synching, you can set up a highly available cluster of Free Chevereto while ensure user login status are not lost if they reconnect to another container.
* `CHEVERETO_DEFAULT_TIMEZONE` - Specify the default timezone for any newly created/registered user, please make sure to use a [PHP supported value](https://www.php.net/manual/en/timezones.php) eg. `America/New_York`.


## Example Usage

I recommend you to use [Docker-compose](https://docs.docker.com/compose/) / [Docker swarm](https://docs.docker.com/engine/swarm/) to launch Chevereto in conjunction with a MySQL database. A sample of docker-compose.yaml can be found below.

### Docker compose

```yaml
version: '3'

services:
  db:
    image: mariadb
    volumes:
      - database:/var/lib/mysql:rw
    restart: always
    networks:
      - private
    environment:
      MYSQL_ROOT_PASSWORD: chevereto_root
      MYSQL_DATABASE: chevereto
      MYSQL_USER: chevereto
      MYSQL_PASSWORD: chevereto

  chevereto:
    depends_on:
      - db
    image: nmtan/chevereto
    restart: always
    networks:
      - private
    environment:
      CHEVERETO_DB_HOST: db
      CHEVERETO_DB_USERNAME: chevereto
      CHEVERETO_DB_PASSWORD: chevereto
      CHEVERETO_DB_NAME: chevereto
      CHEVERETO_DB_PREFIX: chv_
    volumes:
      - chevereto_images:/var/www/html/images:rw
    ports:
      - 8080:80

networks:
  private:
volumes:
  database:
  chevereto_images:
```

Once `docker-compose.yaml` is ready, you can run

```bash
docker-compose up
```

To run the service

### Standalone

```bash
docker run -it --name chevereto -d \
    --link mysql:mysql \
    -p 80:80 \
    -v "$PWD/images":/var/www/html/images \
    -e "CHEVERETO_DB_HOST=db" \
    -e "CHEVERETO_DB_USERNAME=chevereto" \
    -e "CHEVERETO_DB_PASSWORD=chevereto" \
    -e "CHEVERETO_DB_NAME=chevereto" \
    -e "CHEVERETO_DB_PREFIX=chv_" \
    nmtan/chevereto
```

## Note on multi platform

It is feasible to run a Docker container image on different architectures. For now, I don't yet have the time to work on this, but will make sure to include that in future releases.

## Contributions

For now, this project is being maintained solely by me, for any questions or suggestions of improvements, please feel free [to reach out](mailto:tan.mng90@gmail.com)

## License

The docker image is released under the [MIT license](LICENSE)

Please note that [Chevereto](cheveretourl) is a product of [Rodolfo Berrios](http://rodolfoberrios.com/), this project aims mainly at encapsulating the free version (released under AGPL v3 License) into a Docker container image, which can then be used easily.
