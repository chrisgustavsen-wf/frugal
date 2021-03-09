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

_VERSION_MATCH = '.*?..*?..*?'


def main(args):
    root = os.getcwd().rstrip('/')
    print(LANGUAGES.keys())
    if args.version:
        update_frugal_version(args.version.strip('v'), root)
        update_expected_tests(root)


def update_frugal_version(version, root):
    """Update the frugal version."""
    # TODO: Implement dry run
    print(f"Updating frugal to version {version} for {', '.join(LANGUAGES.keys())}")
    update_compiler(version, root)
    for lang in LANGUAGES.values():
        lang.update_frugal(version, root)
    update_tests(version, root)
    update_examples(version, root)


def update_compiler(version, root):
    """Update the frugal compiler."""
    # Update the global version
    os.chdir('{0}/compiler/globals'.format(root))
    base_str = 'const Version = \"{0}\"'
    sub_str = base_str.format(_VERSION_MATCH)
    ver_str = base_str.format(version)
    glob = 'globals.go'
    s = ''
    with open(glob, 'r') as f:
        s = re.sub(sub_str, ver_str, f.read())
    with open(glob, 'w') as f:
        f.write(s)
    # Install the binary with the updated version
    os.chdir(root)
    if subprocess.call(['go', 'mod', 'download']) != 0:
        raise Exception('downloading modules failed')
    if subprocess.call(['go', 'install']) != 0:
        raise Exception('installing frugal binary failed')


def update_tests(version, root):
    """Update the frugal generation tests."""
    os.chdir('{0}/test'.format(root))
    if subprocess.call(['go', 'get', 'github.com/stretchr/testify/assert/...']) != 0:
        raise Exception('Failed to get testify dependency')
    if subprocess.call(['go', 'test', '--copy-files']) != 0:
        raise Exception('Failed to update generated tests')
    if subprocess.call(['frugal', '--gen', 'dart:use_enums=true', '-r', '--out=../test/integration/dart/gen-dart', '../test/integration/frugalTest.frugal']):
        raise Exception('Failed to generate Dart test code')


def update_examples(version, root):
    """Update the examples."""
    os.chdir('{0}/examples'.format(root))
    # TODO: Replace the generate example shell script
    if subprocess.call(['make', 'generate'], stdout=subprocess.DEVNULL) != 0:
        raise Exception('Failed to generate example code')


def update_expected_tests(root):
    for key, value in LANGUAGES.items():
        print(f"Updating expected tests for {key}")
        value.update_expected_tests(root)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description='Update version'
    )
    parser.add_argument('--version', dest='version', type=str)
    args = parser.parse_args()
    main(args)
