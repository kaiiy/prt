# prt

prt replaces full-width characters in a tex file.

| src     | dest |
| ------- | ---- |
| , or 、 | ，   |
| . or 。 | ．   |

## Installation

`go install github.com/kaiiy/prt@latest`

## Usages

### Execute `prt` command

`prt <src-file-path>`

e.g. `prt ./src/main.tex`

### Use prt api with golang

```go
import (
    prt "github.com/kaiiy/prt/lib"
)
var srcText string;

destText := prt.Parse(srcText)
```

## Conversion Rules

The lower the rule, the higher the priority.

1. The conversion targets are the characters in `\begin{document} ..... \end{document}`.
2. The comment line, `% ......` isn't the target.
3. The command line, `\ ......` isn't the target except `caption`, `cite`.
4. The command block, `\begin{} ...... \end{}` isn't the target except `enumerate`, `itemize`.
5. A period, `.` following an alphabet or a number isn't the target.