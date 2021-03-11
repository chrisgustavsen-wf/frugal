import os

from lang.base import LanguageBase


VERSION_FILE = "lib/python/frugal/version.py"


class Python(LanguageBase):
    """
    Python implementation of LanguageBase.
    """

    def update_frugal(self, version):
        """Update the Python version."""

        old_stuff = ""
        with open(VERSION_FILE, "r") as f:
            for line in f:
                if "__version__" in line:
                    old_stuff += "__version__ = '{0}'\n".format(version)
                    break
                old_stuff += line

        with open(VERSION_FILE, "w") as f:
            f.write(old_stuff)
