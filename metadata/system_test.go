package metadata

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var raw = `<?xml version="1.0" encoding="utf-8"?>
<RETS ReplyCode="0" ReplyText="Operation Successful">
  <METADATA>
	<METADATA-SYSTEM Version="01.72.11597" Date="2016-07-21T20:49:13">
	  <SYSTEM SystemID="ABBA" SystemDescription="abba" TimeZoneOffset="-04:00">
		<METADATA-RESOURCE Version="01.72.11597" Date="2016-07-21T20:49:13" System="ABBA">
		  <Resource>
			<ResourceID>Property</ResourceID>
		  </Resource>
		</METADATA-RESOURCE>
	  </SYSTEM>
	</METADATA-SYSTEM>
</METADATA>
</RETS>`

func TestLoad(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader(raw))
	defer body.Close()
	parser := xml.NewDecoder(body)
	xml := RETSResponseWrapper{}

	err := parser.Decode(&xml)
	if err != io.EOF {
		assert.Nil(t, err)
	}
	assert.Equal(t, "Operation Successful", xml.ReplyText)

	assert.Equal(t, "ABBA", xml.Metadata.MSystem.System.ID)
	assert.Equal(t, "Property", string(xml.Metadata.MSystem.System.MResource.Resource[0].ResourceID))
}

func TestSystem(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader(raw))
	defer body.Close()

	extractor := &Extractor{Body: body}
	response, err := extractor.Open()

	assert.Nil(t, err)
	assert.Equal(t, "Operation Successful", response.ReplyText)

	xml := MSystem{}
	err = extractor.DecodeNext("METADATA-SYSTEM", &xml)
	if err != io.EOF {
		assert.Nil(t, err)
	}
	assert.Equal(t, "ABBA", xml.System.ID)
	assert.Equal(t, "Property", string(xml.System.MResource.Resource[0].ResourceID))
}
