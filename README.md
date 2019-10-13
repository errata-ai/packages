# Styles [![Build Status](https://travis-ci.org/errata-ai/styles.svg?branch=master)](https://travis-ci.org/errata-ai/styles)

<p align="center">
  <img width="75%" alt="dash" src="https://user-images.githubusercontent.com/8785025/60774954-0391b300-a0d1-11e9-8d92-1c97f5d07bf4.png">
</p>

This repository contains the library of all managed (i.e., searchable, installable, and updatable) styles for Vale Server.

## Requirements

All styles in this library must (1) be maintained in their own (dedicated) repository, (2) publish releases following [Semantic Versioning](https://semver.org/), and (3) include a `meta.json` file with the following structure:

```json5
{
  # Individual or organization maintaining the style:
  "author": "...",
  # A summary of the style's purpose:
  "description": "...",
  # An ATOM-formatted release feed (doesn't have to be hosted on GitHub)
  "feed": "https://github.com/<USER>/<REPO>/releases.atom",
  # A link to a website or repository for the style:
  "homepage": "...",
  # The name of the style's license -- e.g., "MIT":
  "license": "...",
  # A link to the latest release (doesn't have to be hosted on GitHub)
  "url": "https://github.com/<USER>/<REPO>/releases/latest/download/<NAME>.zip",
  # The minimum required Vale version:
  "vale_version": "v1.0.0"
}
```

## Submitting a style

Fork this repo, add an entry (in alphabetical order) to the [`library.json`](https://github.com/errata-ai/styles/blob/master/library.json) file, and submit a PR.
