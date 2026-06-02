package output

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

func Save(filename string, data interface{}) error {

	if filename == "" {
		return nil
	}

	ext := filepath.Ext(filename)

	switch ext {

	case ".json":
		return saveJSON(filename, data)

	case ".xml":
		return saveXML(filename, data)

	case ".html":
		return saveHTML(filename, data)

	default:
		if ext == "" {
			filename += ".txt"
		}
		return saveTXT(filename, data)
	}
}

func saveTXT(filename string, data interface{}) error {
	return os.WriteFile(filename, []byte(fmt.Sprintf("%v", data)), 0644)
}

func saveJSON(filename string, data interface{}) error {

	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, content, 0644)
}

func saveXML(filename string, data interface{}) error {

	content, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, content, 0644)
}

func saveHTML(filename string, data interface{}) error {

	html := fmt.Sprintf(`
<html>
<head>
<title>JS Intel Report</title>
</head>
<body>
<pre>%v</pre>
</body>
</html>
`, data)

	return os.WriteFile(filename, []byte(html), 0644)
}
