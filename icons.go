package icons

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/a-h/templ"
)

// SVGMap holds icon names mapped to their SVG content (raw string literal).
type SVGMap map[string]string

type csSVGMap SVGMap

var (
	iconContents = make(map[string]string)
	iconMutex    sync.RWMutex
	customSVGMap = make(csSVGMap)
	defFamily    = MDI
)

const (
	BOX       = "box"
	BOOTSTRAP = "bootstrap"
	CUSTOM    = "custom"
	FA        = "fa"
	IONIC     = "ionic"
	LINE      = "line"
	LUCIDE    = "lucide"
	MDI       = "material-design"
	HERO      = "hero"
)

var iconfamilyMap = map[string]any{
	BOOTSTRAP: bs_iconSvgData,
	BOX:       bxSVGMap,
	FA:        fa_iconSvgData,
	HERO:      hero_iconSvgData,
	IONIC:     ionic_iconSvgData,
	LINE:      line_iconSvgData,
	LUCIDE:    lucide_iconSvgData,
	MDI:       mdi_iconSvgData,
	CUSTOM:    customSVGMap,
}

func SetDefaultFamily(family string) error {
	switch family {
	case BOX, BOOTSTRAP, FA, HERO, IONIC, LINE, LUCIDE, MDI:
		defFamily = family
	default:
		return fmt.Errorf("invalid icon family: %s. \nusing default: %s", family, defFamily)

	}

	return nil
}

type IconProps struct {
	Variant string
	Family  string
	Class   string
	Attrs   templ.Attributes
}

func Icon(name string, props ...IconProps) templ.Component {
	var iconFamily = defFamily

	p := IconProps{}
	if len(props) > 0 {
		p = props[0]
	}

	if p.Family != "" {
		iconFamily = p.Family
	}

	iconMutex.Lock()
	defer iconMutex.Unlock()

	cacheKey := fmt.Sprintf("%s|f:%s|c:%s", name, iconFamily, p.Class)
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		iconMutex.RLock()
		svg, cached := iconContents[cacheKey]
		iconMutex.RUnlock()

		if cached {
			_, err := w.Write([]byte(svg))
			return err
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

func generateSVG(name, family string, props IconProps) (string, error) {

	content, exist := iconfamilyMap[family].(SVGMap)[name]

	if !exist {
		content, exist = iconfamilyMap[MDI].(SVGMap)[name]
		if !exist {
			return "", fmt.Errorf("icon not found: %s", name)
		}

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

func AddCustomSvg(name, content string) error {
	if _, exist := customSVGMap[name]; exist {
		return fmt.Errorf("svg already exist: %s", name)
	}

	customSVGMap[name] = content
	return nil

}

func GetAvailableIconCount() int {
	count := 0
	for _, m := range iconfamilyMap {
		count += len(m.(SVGMap))
	}
	return count

}

func GetSVGMapFromJSON(jsonStr string) error {
	err := json.Unmarshal([]byte(jsonStr), &customSVGMap)
	if err != nil {
		return err
	}
	return nil
}

func SaveCustomSvgMapToFile(filePath string) error {
	// 1. Marshal the data into a JSON byte slice
	jsonData, err := json.MarshalIndent(customSVGMap, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// 2. Write the JSON byte slice to the file
	// 0644 sets standard read/write permissions for the file owner.
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}
	return nil
}

func LoadCustomSvgMapFromFile(filePath string) error {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	if err := json.Unmarshal(data, &customSVGMap); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

func GetAvailableIcons(families ...string) ([]string, error) {
	var icons []string

	if len(families) != 0 {
		for _, family := range families {
			f, exist := iconfamilyMap[family]
			if !exist {

				return nil, fmt.Errorf("invalid icon family: %s", family)
			}

			for k := range f.(SVGMap) {
				icons = append(icons, k)
			}
		}

		return icons, nil
	}

	for _, m := range iconfamilyMap {
		for k := range m.(SVGMap) {
			icons = append(icons, k)
		}
	}

	return icons, nil

}

func GetCachedIcons() []string {
	var ics = make([]string, 0, len(iconContents))
	iconMutex.RLock()
	defer iconMutex.Unlock()
	for k := range iconContents {
		ics = append(ics, k)
	}
	return ics
}

type Component struct{}

func (c Component) Render(ctx context.Context, w io.Writer) error {
	return nil
}
