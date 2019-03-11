package plugin

import (
	"html/template"
	"time"
)

// PrettyTime formats time.Time formats to a prettier format for displaying in HTML
// Usage: {{PRETTYTIME .Deadline}}
func PrettyTime() template.FuncMap {
	f := make(template.FuncMap)

	f["PRETTYTIME"] = func(t time.Time) string {
		return t.Format("15:04 02/01/2006") // Norwegian format hh:mm dd/MM/yyyy
	}

	return f
}
