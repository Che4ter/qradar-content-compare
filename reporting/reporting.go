package reporting

import (
	"fmt"
	"os"
	"qradar-content-compare/types"
	"strconv"
	"time"
)

func ReportToTerminal(reports []types.Report) {
	var separator = "=================="
	for _, report := range reports {
		fmt.Println(separator)
		fmt.Println("Report for: ", report.ElementType)
		fmt.Println("Elements Ok: " + strconv.Itoa(report.SameCount))
		if len(report.MissingRecords) > 0 {
			fmt.Println("Elements missing in new QRadar: ")
			for _, missingElement := range report.MissingRecords {
				fmt.Println(missingElement)
			}
		} else {
			fmt.Println("Elements missing in new QRadar: 0")
		}
		fmt.Println(separator)

		if len(report.DifferentRecords) > 0 {
			fmt.Println("Elements different in new QRadar: ")
			for _, differentRecord := range report.DifferentRecords {
				fmt.Println(differentRecord.RecordName)
				for _, differentElement := range differentRecord.DifferentElements {
					fmt.Println("Element: " + differentElement.Name)
					fmt.Println("Old Value: ", differentElement.OldValue)
					fmt.Println("New Value: ", differentElement.NewValue)
				}
				fmt.Println(separator)
			}
		} else {
			fmt.Println("Elements different in new QRadar: 0")
		}
	}
}

func ReportToFiles(reports []types.Report) error {
	separator := "***************************"
	folderName := "qradar_compare_report_"+ time.Now().Format("02_01_2006")+ "/"
	fmt.Println("write to folder "+ folderName)

	for _, report := range reports {
		fileName := folderName + report.ElementType + ".txt"
		if _, err := os.Stat( fileName); os.IsNotExist(err) {
			os.MkdirAll(folderName, 0700)
			// Create your file
		}
		file, err := os.Create(fileName)
		if err != nil{
			return err
		}

		fmt.Fprintln(file, "Report for: ", report.ElementType)
		fmt.Fprintln(file,"Records OK: " + strconv.Itoa(report.SameCount))

		if len(report.MissingRecords) > 0 {
			fmt.Fprintln(file, "")
			fmt.Fprintln(file,"Records missing in new QRadar: ")
			fmt.Fprintln(file, separator)
			for _, missingElement := range report.MissingRecords {
				fmt.Fprintln(file, missingElement)
			}
		} else {
			fmt.Fprintln(file, "")
			fmt.Fprintln(file,"Records missing in new QRadar: 0")
			fmt.Fprintln(file, separator)
		}


		if len(report.DifferentRecords) > 0 {
			fmt.Fprintln(file, "")
			fmt.Fprintln(file,"Records different in new QRadar: ")
			fmt.Fprintln(file, separator)
			for _, differentRecord := range report.DifferentRecords {
				fmt.Fprintln(file,differentRecord.RecordName)
				for _, differentElement := range differentRecord.DifferentElements {
					fmt.Fprintln(file,"Element: " + differentElement.Name)
					fmt.Fprintln(file,"Old Value: ", differentElement.OldValue)
					fmt.Fprintln(file,"New Value: ", differentElement.NewValue)
				}
				fmt.Fprintln(file, "")
			}
		} else {
			fmt.Fprintln(file, "")
			fmt.Fprintln(file,"Records different in new QRadar: 0")
			fmt.Fprintln(file, separator)
		}
		if err := file.Close(); err != nil {
			return err
		}
	}

	return nil
}
