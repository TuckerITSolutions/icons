package icons

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/a-h/templ"
)

// SVGMap holds icon names mapped to their SVG content (raw string literal).
type SVGMap map[string]string

var (
	iconContents = make(map[string]string)
	iconMutex    sync.RWMutex
	customSVGMap = make(SVGMap)
	defFamily    = MDI
)

const (
	BOX       = "box"
	BOOTSTRAP = "bootstrap"
	CUSTOM    = "custom"
	IONIC     = "ionic"
	LINE      = "line"
	MDI       = "material-design"
	HERO      = "hero"
)

func SetDefaultFamily(family string) error {
	switch family {
	case HERO, BOX, BOOTSTRAP, IONIC, LINE, MDI:
		defFamily = family
	default:
		return fmt.Errorf("invalid icon family: %s. \nusing default: %s", family, defFamily)

	}

	return nil
}

func Icon(name string, family ...string) IconFunc {
	var iconFamily = defFamily
	if len(family) != 0 {
		iconFamily = family[0]
	}
	return func(props ...SvgProps) templ.Component {
		var p SvgProps
		if len(props) > 0 {
			p = props[0]
		}
		iconMutex.Lock()
		defer iconMutex.Unlock()

		cacheKey := fmt.Sprintf("%s|f:%s|c:%s", name, iconFamily, props[0].Class)
		return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			iconMutex.RLock()
			svg, cached := iconContents[cacheKey]
			iconMutex.RUnlock()

			if cached {
				_, err := w.Write([]byte(svg))
				if err != nil {
					return err
				}
				return nil
			}

			// Not cached, generate it
			// The actual generation now happens once and is cached.
			generatedSvg, err := generateSVG(name, iconFamily, p) // p (Props) is passed to generateSVG
			if err != nil {
				// Provide more context in the error message
				return fmt.Errorf("failed to generate svg for icon '%s' with props %+v: %w", name, p, err)
			}
			// log.Println(generatedSvg)

			iconMutex.Lock()
			iconContents[cacheKey] = generatedSvg
			iconMutex.Unlock()

			_, err = w.Write([]byte(generatedSvg))
			return err
		})
	}

}

func generateSVG(name, family string, props SvgProps) (string, error) {
	content, err := getIconContent(name, family)
	if err != nil {
		return "", err
	}

	str := strings.Builder{}

	for _, attr := range props.Attrs.Items() {
		if attr.Value == "" {
			fmt.Fprintf(&str, " %s", attr.Key)
		} else {
			fmt.Fprintf(&str, " %s=\"%v\"", attr.Key, attr.Value)
		}

	}

	res := fmt.Sprintf(content, props.Class, str.String())
	return res, nil

}

func getIconContent(name string, family string) (string, error) {
	content := ""
	exists := false
	switch family {
	case HERO:
		content, exists = hero_iconSvgData[name]
	case BOX:
		content, exists = bxSVGMap[name]
	case BOOTSTRAP:
		content, exists = bs_iconSvgData[name]
	case MDI:
		content, exists = mdi_iconSvgData[name]
	case IONIC:
		content, exists = ionic_iconSvgData[name]
	case LINE:
		content, exists = line_iconSvgData[name]
	case CUSTOM:
		content, exists = customSVGMap[name]
	default:
		return "", fmt.Errorf("invalid icon family: %s", family)
	}
	if !exists {
		return "", fmt.Errorf("icon not found: %s", name)
	}
	return content, nil
}

func (s SVGMap) AddSvg(name, content string) error {
	if _, exists := s[name]; exists {
		return fmt.Errorf("svg already exists: %s", name)
	}

	customSVGMap[name] = content
	return nil

}

func GetIconCount() int {
	return len(hero_iconSvgData) + len(bxSVGMap) + len(bs_iconSvgData) + len(ionic_iconSvgData) + len(line_iconSvgData) + len(mdi_iconSvgData) + len(customSVGMap)

}

type MultiError []error

func (me MultiError) Error() string {
	if len(me) == 0 {
		return ""
	}
	// Join all individual error messages with a newline for readability
	s := make([]string, len(me))
	for i, err := range me {
		s[i] = err.Error()
	}
	return "Multiple errors occurred:\n" + strings.Join(s, "\n")
}

func GetAvailableIcons(families ...string) ([]string, error) {
	var icons []string

	var errs MultiError
	if len(families) != 0 {
		for _, family := range families {
			switch family {
			case HERO:
				for k := range hero_iconSvgData {
					icons = append(icons, k)
				}
			case BOX:
				for k := range bxSVGMap {
					icons = append(icons, k)
				}
			case BOOTSTRAP:
				for k := range bs_iconSvgData {
					icons = append(icons, k)
				}
			case IONIC:
				for k := range ionic_iconSvgData {
					icons = append(icons, k)
				}
			case LINE:
				for k := range line_iconSvgData {
					icons = append(icons, k)
				}
			case MDI:
				for k := range mdi_iconSvgData {
					icons = append(icons, k)
				}
			case CUSTOM:
				for k := range customSVGMap {
					icons = append(icons, k)
				}
			default:
				errs = append(errs, fmt.Errorf("invalid icon family: %s", family))

			}
		}

		return icons, errs
	}
	for k := range hero_iconSvgData {
		icons = append(icons, k)
	}
	for k := range bxSVGMap {
		icons = append(icons, k)
	}
	for k := range bs_iconSvgData {
		icons = append(icons, k)
	}

	for k := range ionic_iconSvgData {
		icons = append(icons, k)
	}
	for k := range line_iconSvgData {
		icons = append(icons, k)
	}
	for k := range mdi_iconSvgData {
		icons = append(icons, k)
	}
	for k := range customSVGMap {
		icons = append(icons, k)
	}
	return icons, errs

}

func GetCachedIcons() []string {
	var ics = make([]string, 0, len(iconContents))
	iconMutex.Lock()
	defer iconMutex.Unlock()
	for k := range iconContents {
		ics = append(ics, k)
	}
	return ics
}

type SvgProps struct {
	Class string
	Attrs templ.Attributes
}

type IconFunc func(...SvgProps) templ.Component
