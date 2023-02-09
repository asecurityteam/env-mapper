import pytest
import testinfra


def test_simple(host):
    c = host.run('A=AVALUE /env-mapper -complex "Z:||A||" -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=AVALUE' in c.stdout

def test_prefix(host):
    c = host.run('A=AVALUE /env-mapper -complex "Z:a||A||" -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=aAVALUE' in c.stdout

def test_postfix(host):
    c = host.run('A=AVALUE /env-mapper -complex "Z:||A||a" -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=AVALUEa' in c.stdout

def test_both(host):
    c = host.run('A=AVALUE /env-mapper -complex "Z:pre||A||post" -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=preAVALUEpost' in c.stdout

def test_multiple_mappings(host):
    c = host.run('A=AVALUE B=BVALUE /env-mapper -complex "Z:||A||"  "Y:pre||B||post" -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=AVALUE' in c.stdout
    assert 'Y=preBVALUEpost' in c.stdout

def test_no_mapping(host):
    c = host.run('A=AVALUE /env-mapper -complex -- /usr/bin/env')
    assert c.rc == 0
    assert 'A=AVALUE' in c.stdout

def test_custom_separator(host):
    c = host.run('A=AVALUE B=BVALUE /env-mapper -complex --envSep=^ "Z^||A||"  "Y^pre||B||post" -- /usr/bin/env')
    assert c.rc == 0
    assert 'Z=AVALUE' in c.stdout
    assert 'Y=preBVALUEpost' in c.stdout

def test_missing_dashes(host):
    c = host.run('A=AVALUE /env-mapper -complex Z:A')
    assert c.rc != 0
    assert 'Usage' in c.stderr


def test_missing_command(host):
    c = host.run('A=AVALUE /env-mapper -complex Z:A --')
    assert c.rc != 0
    assert 'Usage' in c.stderr


def test_no_mapping(host):
    c = host.run('A=AVALUE /env-mapper -complex -- /usr/bin/env')
    assert c.rc != 0


def test_invalid_command(host):
    c = host.run('A=AVALUE /env-mapper -complex "Z:||A||" -- /nosuchcommand')
    assert c.rc != 0

def test_command_return_value(host):
    c = host.run('A=AVALUE /env-mapper -complex "Z:||A||" -- /bin/false')
    assert c.rc != 0
