import re
import pytest
import subprocess
import pathlib

ROOT = pathlib.Path(__file__).absolute().parent.parent
POCNAME_PATTERN = re.compile(r'\A(?!-)[a-z0-9\-]+(?<!-)\.yml\Z')


@pytest.fixture
def filenames():
    diff = subprocess.check_output(['git', 'diff', '--name-only', 'master'], cwd=str(ROOT))
    if diff:
        return [filename.strip() for filename in diff.decode().split('\n')]
    else:
        return []


def test_filename(filenames):
    for filename in filenames:
        poc_file = pathlib.Path(filename)
        assert poc_file.parent.absolute() == ROOT.absolute(), 'POC must be in pocs/ folder, without subfolder'
        assert poc_file.suffix == '.yml', 'POC extension must be .yml'
        assert POCNAME_PATTERN.match(poc_file.name), 'filename format is wrong'
