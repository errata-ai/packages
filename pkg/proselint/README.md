# proselint [![Build Status](https://travis-ci.org/errata-ai/proselint.svg?branch=master)](https://travis-ci.org/errata-ai/proselint) ![Vale version](https://img.shields.io/badge/vale-%3E%3D%20v1.7.0-blue.svg) ![license](https://img.shields.io/github/license/mashape/apistatus.svg)

> [`proselint`](https://github.com/amperser/proselint/) places the world’s greatest writers and editors by your side, where they whisper suggestions on how to improve your prose.

This repository contains a [Vale-compatible](https://github.com/errata-ai/vale) implementation of the guidelines enforced by the `proselint` ([LICENSE](https://github.com/amperser/proselint/blob/master/LICENSE.md)) linter.

## Getting Started

> :exclamation: proselint requires Vale >= **1.7.0**. :exclamation:

Download the [latest release](https://github.com/errata-ai/proselint/releases), copy the "proselint" directory to your `StylesPath`, and include it in your configuration file:

```ini
# This goes in a file named either `.vale.ini` or `_vale.ini`.
StylesPath = path/to/some/directory
MinAlertLevel = warning # suggestion, warning or error

# Only Markdown and .txt files; change to whatever you're using.
[*.{md,txt}]
# List of styles to load.
BasedOnStyles = proselint
```

See [Usage](https://github.com/errata-ai/vale/#usage) for more information.
