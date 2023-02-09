import pytest
import testinfra


def test_simple(host):
    c = host.run('A=AVALUE /env-mapper Z:A -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=AVALUE' in c.stdout


def test_multiple_mappings(host):
    c = host.run('A=AVALUE B=BVALUE /env-mapper Z:A  Y:B -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=AVALUE' in c.stdout
    assert 'Y=BVALUE' in c.stdout

def test_no_mapping(host):
    c = host.run('A=AVALUE /env-mapper -- /usr/bin/env')
    assert c.rc == 0
    assert 'A=AVALUE' in c.stdout

def test_custom_separator(host):
    c = host.run('A=AVALUE B=BVALUE /env-mapper --envSep=^ Z^A  Y^B -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=AVALUE' in c.stdout
    assert 'Y=BVALUE' in c.stdout


def test_help(host):
    c = host.run('/env-mapper')
    assert c.rc != 0
    assert 'Usage' in c.stderr


def test_missing_dashes(host):
    c = host.run('A=AVALUE /env-mapper Z:A')
    assert c.rc != 0
    assert 'Usage' in c.stderr


def test_missing_command(host):
    c = host.run('A=AVALUE /env-mapper Z:A --')
    assert c.rc != 0
    assert 'Usage' in c.stderr


def test_no_mapping(host):
    c = host.run('A=AVALUE /env-mapper -- /usr/bin/env')
    assert c.rc != 0


def test_invalid_command(host):
    c = host.run('A=AVALUE /env-mapper Z:A -- /nosuchcommand')
    assert c.rc != 0

def test_command_return_value(host):
    c = host.run('A=AVALUE /env-mapper Z:A -- /bin/false')
    assert c.rc != 0
