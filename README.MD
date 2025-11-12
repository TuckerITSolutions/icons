# Icons

This package provides a collection of SVG icons for use in web applications. This module relies on github.com/a-h/templ for templating.

Icons will return as templ.Components.

## Installation

```bash
go get github.com/tuckeritsolutions/icons
```

## Usage

To use an icon, simply import the package and call the desired icon function. For example:

```go
package main

import (
	"fmt"
	"github.com/tuckeritsolutions/icons"
)

func main() {
	icons.Home()
}
```

You can always call the Icon function directly. Example:

```go
package main

import (
"fmt"
"github.com/tuckeritsolutions/icons"
)

func main() {
    icons.Icon("page-last", icons.SvgProps{Attrs: map[string]string{"width": "24", "height": "24"}})
}
```

This will output the SVG code for the home icon.

## Example Available Icons

There are currently 16034 icons available across all icon families (box, bootstrap, hero, ionic, line, lucide, mdi).

All Icons show as Material Design Icon by default.

*   `Home()`: A house icon.
*   `Settings()`: A gear icon.
*   `User()`: A person icon.
*   `Plus()`: A plus sign icon.
*   `Minus()`: A minus sign icon.
*   `Check()`: A checkmark icon.
*   `Close()`: A close/multiply icon.
*   `Info()`: An information icon.
*   `Warning()`: A warning icon.
*   `Error()`: An error icon.
*   `Chevron()`: A chevron icon that defaults to the down direction

## Variations 

Simply add the variation of icon you would like to display.

* `Chevron("up", "circle", "outline)`: Produces the chevron-up-circle-outline icon from material design icons.

## Contributing

If you'd like to contribute new icons or improve existing ones, please feel free to open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
