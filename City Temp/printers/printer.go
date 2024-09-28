package printers

import (
	"city_temp/models"
	"fmt"
	"os"
	"text/tabwriter"
)

type Printer struct {
	writer *tabwriter.Writer
}

func New() *Printer {
	return &Printer{writer: tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)}
}

func (p *Printer) CityHeader() {
	fmt.Fprintln(p.writer, "Id\tName\tTempC\tTempF\tBeach ready?\tSki Ready?")
}

func (p *Printer) CityDetails(c models.CityTemp) {
	fmt.Fprintf(p.writer, "%v\t%v\t%v\t%v\t%v\t%v\n", c.Id(), c.Name(), c.TempC(), c.TempF(), c.BeachVacationReady(), c.SkiVacationReady())
}

func (p *Printer) Flush() {
	p.writer.Flush()
}
