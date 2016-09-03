package metadata

import "time"

// MSearchHelp ...
type MSearchHelp struct {
	Date       time.Time    `xml:"Date,attr"`
	Version    Version      `xml:"Version,attr"`
	Resource   string       `xml:"Resource,attr"`
	SearchHelp []SearchHelp `xml:"SearchHelp"`
}

// SearchHelp ...
type SearchHelp struct {
	MetadataEntryID string `xml:"MetadataEntryID"`
	SearchHelpID    string `xml:"SearchHelpID"`
	Value           string `xml:"Value"`
}
