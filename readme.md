# columbo

columbo - got them clues

## Usage

```
> columbo -r columbo.yaml -o ~/tmp/output <tarball>
```

## Description

Pretty straightforward, it parses a yaml specification of regexs and tries to
find errors/concerns within output files. It'll take any tarball with plain text
files and parse each file concurrently for matches. Results are stored as JSON
in the output directory.

## Rules Spec

An example rules file looks like:

```yaml

- id: python-tb-exception
  description: parses logs for tracebacks
  start_marker: "^Traceback.*"
  end_marker: "^.*Error|InvalidRequest:"
```

You can also match by line:

```yaml

- id: subprocess-exit-status
  description: pulls lines with an exit status in the text
  line_match: "exit status 1"
```

## Building Locally

Developed with Go version 1.14

```
> go build ./cmd/columbo.go
```

## AsciiCast

[![asciicast](https://asciinema.org/a/MUs0GdCUxsN89C3fDlRUEHfKI.svg)](https://asciinema.org/a/MUs0GdCUxsN89C3fDlRUEHfKI)

## More information

- [Website / Documentation](https://columbo.8op.org)
