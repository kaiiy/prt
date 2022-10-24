# pretex

pretex replaces full-width characters in a tex file.

| src     | dest |
| ------- | ---- |
| , or 、 | ，   |
| . or 。 | ．   |

## Installation

`go install github.com/kaiiy/pretex@latest`

## Usages

### Execute `pretex` command

`pretex <src-file-path>`

e.g. `pretex ./src/main.tex`

### Use pretex api with golang

```go
import (
    pretex "github.com/kaiiy/pretex/lib"
)
var srcText string;

destText := pretex.Parse(srcText)
```

## Conversion Rules

The lower the rule, the higher the priority.

1. The conversion targets are the characters in `\begin{document} ..... \end{document}`.
2. The comment line, `% ......` isn't the target.
3. The command line, `\ ......` isn't the target except `caption`, `cite`.
4. The command block, `\begin{} ...... \end{}` isn't the target except `enumerate`, `itemize`.
5. A period, `.` following an alphabet or a number isn't the target.