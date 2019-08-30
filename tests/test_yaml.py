import sys
import logging
import pytest
import pathlib
import yaml
import json
from jsonschema import validate, draft7_format_checker

logging.basicConfig(stream=sys.stdout, level=logging.INFO)
ROOT = pathlib.Path(__file__).absolute().parent.parent
SCHEMA_FILE = ROOT / 'tests' / 'schema.json'
SCHEMA_DATA = json.loads(SCHEMA_FILE.read_bytes())


@pytest.fixture
def pocs():
    return [file for file in ROOT.glob('pocs/*.yml')]


def check_field(f: pathlib.Path):
    logging.info("check for %s", f.name)
    data = yaml.safe_load(f.read_bytes())

    validate(instance=data, schema=SCHEMA_DATA, format_checker=draft7_format_checker)


def check_poc_name(name):
    pass


def test_yaml_parse(pocs):
    for f in pocs:
        check_field(f)
