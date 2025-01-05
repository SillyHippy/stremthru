[![GitHub Workflow Status: CI [SDK/PY]](https://img.shields.io/github/actions/workflow/status/SillyHippy/stremthru/ci-sdk-py.yml?branch=main&label=CI%20%5BSDK%2FPY%5D&style=for-the-badge)](https://github.com/SillyHippy/stremthru/actions/workflows/ci-sdk-py.yml)
[![PyPI - Version](https://img.shields.io/pypi/v/stremthru?style=for-the-badge)](https://pypi.org/project/stremthru/)
[![License](https://img.shields.io/github/license/SillyHippy/stremthru?style=for-the-badge)](https://github.com/SillyHippy/stremthru/blob/main/sdk/py/LICENSE)

# StremThru - Python SDK

## Installation

```sh
pip install stremthru
# or
poetry add stremthru
```

## Usage

**Basic Usage:**

```py
from stremthru import StremThru;

st = StremThru(base_url="http://127.0.0.1:8080", auth="user:pass")
```

## License

Licensed under the MIT License. Check the [LICENSE](./LICENSE) file for details.
