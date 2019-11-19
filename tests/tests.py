import json
import os
import pathlib
import re
import requests
import subprocess

import yaml
from jsonschema import validate, draft7_format_checker
from jsonschema.exceptions import ValidationError

ROOT = pathlib.Path(__file__).absolute().parent.parent
POCNAME_PATTERN = re.compile(r"\A(?!-)[a-z0-9\-]+(?<!-)\.yml\Z")
SCHEMA_FILE = ROOT / "tests" / "schema.json"
SCHEMA_DATA = json.loads(SCHEMA_FILE.read_bytes())

file_check_failed_details = {}
schema_check_failed_details = {}
yaml_lint_output = ""


def filenames():
    diff = subprocess.check_output(["git", "diff", "--name-only", "origin/master"], cwd=str(ROOT))
    if diff:
        return [filename.strip() for filename in diff.decode().split("\n") if filename.strip()]
    else:
        return []


def check_file():
    for filename in filenames():
        if filename not in file_check_failed_details:
            file_check_failed_details[filename] = []
        poc_file = pathlib.Path(filename)
        if not filename.startswith("poc-"):
            file_check_failed_details[filename].append("文件名必须以 poc- 开头")
        if not poc_file.parent.absolute() == ROOT.joinpath("pocs").absolute():
            file_check_failed_details[filename].append("poc 文件都必须在 pocs 文件夹中")
        if not poc_file.suffix == ".yml":
            file_check_failed_details[filename].append("文件名必须是 .yml 拓展名")
        if not POCNAME_PATTERN.match(poc_file.name):
            file_check_failed_details[filename].append("文件名必须符合正则规范")


def check_field(f: pathlib.Path):
    data = yaml.safe_load(f.read_bytes())
    validate(instance=data, schema=SCHEMA_DATA, format_checker=draft7_format_checker)


def check_yaml_schema():
    for f in [file for file in ROOT.glob("pocs/*.yml")]:
        try:
            check_field(f)
        except ValidationError as e:
            schema_check_failed_details[f.name] = str(e)


def check_yaml_lint():
    try:
        subprocess.check_output(["yamllint", "-c", "tests/yamllint.yml", "-f", "parsable", "pocs/"], cwd=str(ROOT),
                                stderr=subprocess.STDOUT)
    except subprocess.CalledProcessError as e:
        global yaml_lint_output
        yaml_lint_output = [item for item in e.stdout.decode("utf-8").split("\n") if item.strip()]


check_file()
check_yaml_schema()
check_yaml_lint()

if file_check_failed_details or schema_check_failed_details or yaml_lint_output:
    msg = ""
    if file_check_failed_details:
        msg += "## poc 路径检查结果\n\n"
        for k, v in file_check_failed_details.items():
            msg += k + "\n"
            for item in v:
                msg += "  - " + item + "\n"
        msg += "\n\n"

    if schema_check_failed_details:
        msg += "## poc 格式检查结果\n\n"
        for k, v in schema_check_failed_details.items():
            msg += k + "\n\n" + "```\n" + v + "\n```\n\n"

    if yaml_lint_output:
        msg += "## yaml 规范检查结果\n\n"
        for item in yaml_lint_output:
            msg += "  - " + item + "\n"
    print(msg)

    pr_id = os.environ.get("TRAVIS_PULL_REQUEST")
    if pr_id and pr_id != "false":
        resp = requests.post("https://api.github.com/repos/chaitin/xray/issues/" + pr_id + "/comments", json={"body": msg}, headers={"Authorization": os.environ.get("github_basic_auth")})
        print(resp.json())
    if msg:
        exit(1)
