# Styles [![Build Status](https://travis-ci.org/errata-ai/styles.svg?branch=master)](https://travis-ci.org/errata-ai/styles)

This repository contains a library of all officially supported styles for Vale and Vale Server.

<table>
    <tr>
        <th>Vale</th>
        <th>Vale Server</th>
    </tr>
    <tr>
        <td width="50%">
          <a href="https://user-images.githubusercontent.com/8785025/63803049-d7bccd80-c8c8-11e9-97fd-169631f80be9.png">
                <img src="https://user-images.githubusercontent.com/8785025/63803049-d7bccd80-c8c8-11e9-97fd-169631f80be9.png" width="100%">
            </a>
        </td>
        <td width="50%">
            <a href="https://user-images.githubusercontent.com/8785025/60774954-0391b300-a0d1-11e9-8d92-1c97f5d07bf4.png">
                <img src="https://user-images.githubusercontent.com/8785025/60774954-0391b300-a0d1-11e9-8d92-1c97f5d07bf4.png" width="100%">
            </a>
        </td>
    </tr>
    <tr>
        <td width="50%">
          <a href="https://github.com/errata-ai/vale">Vale</a> is an open-source, command-line linter for prose. It's fast, syntax-aware, and extensible.
        </td>
        <td width="50%">
          <a href="https://github.com/errata-ai/vale-server">Vale Server</a> was a commercial desktop application (macOS and Windows) that enhanced and refined the Vale experience.
    </tr>
</table>

The benefits of using these styles over their original implementations include:

- [X] [Improved support for markup](https://vale.sh/docs/topics/scoping/), including the ability to ignore code and target only certain sections of text (e.g., checking headers for a specific capitalization style).
- [X] No need to install and configure npm (Node.js), pip (Python), or other language-specific tools. With Vale, you get all the functionality in a single, standalone binary available for Windows, macOS, and Linux.
- [X] Easily combine, mismatch, or otherwise customize each style.

## Available styles

<dl>
  <dt><a href="https://github.com/errata-ai/Microsoft"><code>Microsoft</code></a></dt>
  <dd>An implementation of the <a href="https://docs.microsoft.com/en-us/style-guide/welcome/"><i>Microsoft Writing Style Guide</i></a>.</dd>

  <dt><a href="https://github.com/errata-ai/Google"><code>Google</code></a></dt>
  <dd>An implementation of the <a href="https://developers.google.com/style/"><i>Google Developer Documentation Style Guide</i></a>.</dd>

  <dt><a href="https://github.com/errata-ai/write-good"><code>write-good</code></a></dt>
  <dd>An implementation of the guidelines enforced by the <a href="https://github.com/btford/write-good"><code>write-good</code></a> linter.</dd>

  <dt><a href="https://github.com/errata-ai/proselint"><code>proselint</code></a></dt>
  <dd>An implementation of the guidelines enforced by the <a href="https://github.com/amperser/proselint/"><code>proselint</code></a> linter.</dd>

  <dt><a href="https://github.com/errata-ai/Joblint"><code>Joblint</code></a></dt>
  <dd>An implementation of the guidelines enforced by the <a href="https://github.com/rowanmanning/joblint"><code>Joblint</code></a> linter.</dd>

  <dt><a href="https://github.com/errata-ai/alex"><code>alex</code></a></dt>
  <dd>An implementation of the guidelines enforced by the <a href="https://github.com/get-alex/alex"><code>alex</code></a> linter.</dd>

  <dt><a href="https://github.com/errata-ai/readability"><code>Readability</code></a></dt>
  <dd>An implementations of many popular "readability" metrics.</dd>
</dl>

## Requirements

All styles in this library must (1) be maintained in their own (dedicated) repository, (2) publish releases following [Semantic Versioning](https://semver.org/), and (3) include a `meta.json` file with the following structure:

```json
{
  "feed": "...",
  "vale_version": "..."
}
```

where `feed` is an Atom-formatted release feed (e.g., `https://github.com/<USER>/<REPO>/releases.atom`) and `vale_version` is the minimum required Vale version (e.g., `v1.0.0`).

## Submitting a style

Fork this repo, add an entry (in alphabetical order) to the [`library.json`](https://github.com/errata-ai/styles/blob/master/library.json) file, and submit a PR.
