# Readability

This repository contains a [Vale-compatible][1] implementation of many popular
"readability" metrics.

## Getting Started

> :exclamation: Readability requires Vale >= **2.13.0**. :exclamation:

Download the [latest release][2], copy the "Readability" directory to your
`StylesPath`, and include it in your configuration file:

```ini
# This goes in a file named either `.vale.ini` or `_vale.ini`.
StylesPath = path/to/some/directory
MinAlertLevel = warning # suggestion, warning or error

# Only Markdown and .txt files; change to whatever you're using.
[*.{md,txt}]
# List of styles to load.
BasedOnStyles = Readability
```

See [Usage][3] for more information.

[1]: https://github.com/errata-ai/vale
[2]: https://github.com/errata-ai/readability/releases
[3]: https://github.com/errata-ai/vale/#usage