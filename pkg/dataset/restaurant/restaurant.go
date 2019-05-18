package restaurant

import (
	"strings"

	"github.com/ray1729/goml/pkg/csvreader"
	"github.com/ray1729/goml/pkg/dataset"
)

const DATA = `Alt,Bar,Fri,Hun,Pat,Price,Rain,Res,Type,Est,Wait
Yes,No,No,Yes,Some,$$$,No,Yes,French,0-10,Yes
Yes,No,No,Yes,Full,$,No,No,Thai,30-60,No
No,Yes,Yes,No,Some,$,No,No,Burger,0-10,Yes
Yes,No,No,Yes,Full,$,Yes,No,Thai,10-30,Yes
Yes,No,No,No,Full,$$$,No,Yes,French,>60,No
No,Yes,Yes,Yes,Some,$$,Yes,Yes,Italian,0-10,Yes
No,Yes,Yes,No,None,$,Yes,No,Burger,0-10,No
No,No,No,Yes,Some,$$,Yes,Yes,Thai,0-10,Yes
No,Yes,Yes,No,Full,$,Yes,No,Burger,>60,No
Yes,Yes,Yes,Yes,Full,$$$,No,Yes,Italian,10-30,No
No,No,No,No,None,$,No,No,Thai,0-10,No
Yes,Yes,Yes,Yes,Full,$,No,No,Burger,30-60,Yes`

func ReadDataset() (*dataset.Dataset, error) {
	r := strings.NewReader(DATA)
	yesno := csvreader.NewEnumerator("No", "Yes")
	spec := map[string]csvreader.ValCoercer{
		"Alt":   yesno,
		"Bar":   yesno,
		"Fri":   yesno,
		"Hun":   yesno,
		"Pat":   csvreader.NewEnumerator("None", "Some", "Full"),
		"Price": csvreader.NewEnumerator("$", "$$", "$$$"),
		"Rain":  yesno,
		"Res":   yesno,
		"Type":  csvreader.NewEnumerator("French", "Thai", "Burger", "Italian"),
		"Est":   csvreader.NewEnumerator("0-10", "10-30", "30-60", ">60"),
		"Wait":  yesno,
	}
	return csvreader.ReadCSV(r, csvreader.NewRowCoercer(spec))
}

func MustReadDataset() *dataset.Dataset {
	ds, err := ReadDataset()
	if err != nil {
		panic(err)
	}
	return ds
}
