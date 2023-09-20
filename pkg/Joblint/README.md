# Joblint [![Build Status](https://travis-ci.org/errata-ai/Joblint.svg?branch=master)](https://travis-ci.org/errata-ai/Joblint) ![Vale version](https://img.shields.io/badge/vale-%3E%3D%20v1.7.0-blue.svg) ![license](https://img.shields.io/github/license/mashape/apistatus.svg)

> [`Joblint`](https://github.com/rowanmanning/joblint): Test tech job posts for issues with sexism, culture, expectations, and recruiter fails.

This repository contains a [Vale-compatible](https://github.com/errata-ai/vale) implementation of the guidelines enforced by the `Joblint` ([LICENSE](https://github.com/rowanmanning/joblint/blob/master/LICENSE)) linter.

## Getting Started

> :exclamation: Joblint requires Vale >= **1.7.0**. :exclamation:

Download the [latest release](https://github.com/errata-ai/Joblint/releases), copy the "Joblint" directory to your `StylesPath`, and include it in your configuration file:

```ini
# This goes in a file named either `.vale.ini` or `_vale.ini`.
StylesPath = path/to/some/directory
MinAlertLevel = warning # suggestion, warning or error

# Only Markdown and .txt files; change to whatever you're using.
[*.{md,txt}]
# List of styles to load.
BasedOnStyles = Joblint
```

See [Usage](https://github.com/errata-ai/vale/#usage) for more information.

## Using the Template

This repository serves as a [template](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) for creating your own style. The following files/folders should be edited with your own infomation:

- [`Joblint`](https://github.com/errata-ai/Joblint/tree/master/Joblint): This is where your actual style should be implemented. Re-name the folder and replace its content with your `.yml` rule files.

- [`features`](https://github.com/errata-ai/Joblint/tree/master/features) + [`fixtures`](https://github.com/errata-ai/Joblint/tree/master/fixtures): This is where your tests should reside.

- [`LICENSE`](https://github.com/errata-ai/Joblint/blob/master/LICENSE): replace with your own license information.

- [`.travis.yml`](https://github.com/errata-ai/Joblint/blob/master/.travis.yml): Replace all instances of "Joblint" with the name of your style. You'll also need to add your own [`$GITHUB_TOKEN`](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line) to your [Travis CI repository settings](https://docs.travis-ci.com/user/environment-variables/#defining-variables-in-repository-settings).
