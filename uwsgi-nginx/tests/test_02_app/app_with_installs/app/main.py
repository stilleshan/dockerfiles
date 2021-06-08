import sys

from flask import Flask

application = Flask(__name__)


@application.route("/")
def hello():
    version = "{}.{}".format(sys.version_info.major, sys.version_info.minor)
    message = "Hello World from Nginx uWSGI Python {} app in a Docker container".format(
        version
    )
    return message
