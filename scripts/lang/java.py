import os
from xml.etree import ElementTree

from lang.base import LanguageBase


_POM = "pom"
_POM_XML = "{0}/pom.xml"
_WORKIVA = "com.workiva"
_NS = {_POM: "http://maven.apache.org/POM/4.0.0"}


class Java(LanguageBase):
    """
    Java implementation of LanguageBase. Uses xml.tree.ElementTree to parse all
    pom.xml's.
    """

    def update_frugal(self, version):
        """Update the java version."""
        # Update library pom
        self._update_maven_version("lib/java", version)

        # Update example pom
        self._update_maven_version("examples/java", version)
        self._update_maven_dep("examples/java", _WORKIVA, "frugal", version)

        # Update integration tests
        self._update_maven_dep(
            "test/integration/java/frugal-integration-test", _WORKIVA, "frugal", version
        )

    def _update_maven_version(self, where, version):
        """Update the project version in the current directory's pom.xml."""
        pwd = _POM_XML.format(where)
        tree = ElementTree.parse(pwd)
        ver = tree.getroot().find("{0}:version".format(_POM), _NS)
        ver.text = version
        tree.write(pwd, default_namespace=_NS[_POM])

    def _update_maven_dep(self, where, group, artifact, version):
        """Update a maven dependency in the current directory's pom.xml."""
        pwd = _POM_XML.format(where)
        tree = ElementTree.parse(pwd)
        for dep in tree.getroot().find("{0}:dependencies".format(_POM), _NS):
            g = dep.find("{0}:groupId".format(_POM), _NS)
            a = dep.find("{0}:artifactId".format(_POM), _NS)
            if g.text == group and a.text == artifact:
                dep.find("{0}:version".format(_POM), _NS).text = version
        tree.write(pwd, default_namespace=_NS[_POM])
