package xml

import "github.com/clbanning/mxj/v2"

type ConversorXML struct{}

func (x ConversorXML) Conversor(data string) (map[string]any, error) {
	xml, error := mxj.NewMapXml([]byte(data))

	if error != nil {
		return nil, error
	}

	return xml.Old(), nil
}

func (x ConversorXML) Formato() string {
	return "XML"
}
