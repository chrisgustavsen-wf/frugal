import os

from yaml import dump, safe_load

from lang.base import LanguageBase


class Dart(LanguageBase):
    """
    Dart implementation of LanguageBase. Uses PyYAML to parse all
    pubspec.yaml's.
    """

    def update_frugal(self, version):
        """Update the dart version."""
        # Update libary pubspec
        def update_lib(data):
            data["version"] = version

        self._update("lib/dart", update_lib, "Dart lib")

        # Update example pubspec
        def update_example(data):
            data["dependencies"]["frugal"]["version"] = "^{0}".format(version)

        self._update("examples/dart", update_example, "Dart example")

    def _update(self, where, update, prefix):
        """
        Update pubspec.yaml in current directory using the given update
        function.
        """
        pubspec = where + "/pubspec.yaml"
        with open(pubspec, "r") as f:
            data = safe_load(f.read())
            update(data)
        with open(pubspec, "w") as f:
            dump(data, f, default_flow_style=False)
