package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ray1729/gota/dataframe"
	"github.com/ray1729/gota/series"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s PATH\n", os.Args[0])
	}
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	df := dataframe.ReadCSV(file)
	fmt.Println(df)

	male := df.Filter(dataframe.F{
		Colname:    "Gender",
		Comparator: series.Eq,
		Comparando: "Male",
	})

	female := df.Filter(dataframe.F{
		Colname:    "Gender",
		Comparator: series.Eq,
		Comparando: "Female",
	})

	fmt.Println(df.Describe())
	fmt.Println(male.Describe())
	fmt.Println(female.Describe())

}
