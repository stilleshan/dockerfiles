import os

from docker.client import DockerClient
from docker.errors import NotFound
from docker.models.containers import Container

CONTAINER_NAME = "uwsgi-nginx-test"


def get_logs(container: Container) -> str:
    logs = container.logs()
    return logs.decode("utf-8")


def get_nginx_config(container: Container) -> str:
    result = container.exec_run(f"/usr/sbin/nginx -T")
    return result.output.decode()


def remove_previous_container(client: DockerClient) -> None:
    try:
        previous = client.containers.get(CONTAINER_NAME)
        previous.stop()
        previous.remove()
    except NotFound:
        return None


def get_response_text1() -> str:
    python_version = os.getenv("PYTHON_VERSION")
    return f"Hello World from a default Nginx uWSGI Python {python_version} app in a Docker container (default)"


def get_response_text2() -> str:
    python_version = os.getenv("PYTHON_VERSION")
    return f"Hello World from Nginx uWSGI Python {python_version} app in a Docker container"


def generate_dockerfile_content_custom_app(name: str) -> str:
    content = f"FROM tiangolo/uwsgi-nginx:{name}\n"
    content += "COPY ./application /application\n"
    content += "COPY ./prestart.sh /app/prestart.sh\n"
    content += "WORKDIR /application\n"
    return content


def generate_dockerfile_content_custom_nginx_app(name: str) -> str:
    content = f"FROM tiangolo/uwsgi-nginx:{name}\n"
    content += "COPY app /app\n"
    return content


def generate_dockerfile_content_simple_app(name: str) -> str:
    content = f"FROM tiangolo/uwsgi-nginx:{name}\n"
    content += "COPY ./app/main.py /app/main.py\n"
    return content


def generate_dockerfile_content_app_with_installs(name: str) -> str:
    content = f"FROM tiangolo/uwsgi-nginx:{name}\n"
    content += "RUN pip install flask\n"
    content += "COPY ./app/main.py /app/main.py\n"
    return content
