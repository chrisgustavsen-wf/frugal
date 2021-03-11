#! /usr/bin/env python3

import argparse
import os
import re
import subprocess

from lang import Dart, Go, Java, Python


LANGUAGES = {
    "dart": Dart(),
    "go": Go(),
    "java": Java(),
    "python": Python(),
}

_VERSION_MATCH = ".*?..*?..*?"


def main(args):
    update_frugal_version(args.version.strip("v"))


def update_frugal_version(version):
    """Update the frugal version."""
    # TODO: Implement dry run
    print(f"Updating frugal to version {version} for {', '.join(LANGUAGES.keys())}")
    update_compiler(version)
    install_frugal()
    for lang in LANGUAGES.values():
        lang.update_frugal(version)
    update_tests()
    update_examples()


def update_compiler(version):
    """Update the frugal compiler."""
    # Update the global version
    base_str = 'const Version = "{0}"'
    sub_str = base_str.format(_VERSION_MATCH)
    ver_str = base_str.format(version)
    glob = "compiler/globals/globals.go"
    s = ""
    with open(glob, "r") as f:
        s = re.sub(sub_str, ver_str, f.read())
    with open(glob, "w") as f:
        f.write(s)


def install_frugal():
    # Install the binary with the updated version
    if subprocess.call(["go", "mod", "download"]) != 0:
        raise Exception("downloading modules failed")
    if subprocess.call(["go", "install"]) != 0:
        raise Exception("installing frugal binary failed")


def update_tests():
    """Update the frugal generation tests."""
    if subprocess.call(["go", "test", "./...", "--copy-files"]) != 0:
        raise Exception("Failed to update generated tests")
    if subprocess.call(
        [
            "frugal",
            "--gen",
            "dart:use_enums=true",
            "-r",
            "--out=dart/gen-dart",
            "frugalTest.frugal",
        ],
        cwd="./test/integration",
    ):
        raise Exception("Failed to generate Dart test code")


def update_examples():
    """Update the examples."""
    # TODO: Replace the generate example shell script
    if subprocess.call(
        ["make", "generate"], stdout=subprocess.DEVNULL, cwd="./examples",
    ):
        raise Exception("Failed to generate example code")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Update version")
    parser.add_argument("--version", dest="version", type=str, required=True)
    args = parser.parse_args()
    main(args)
