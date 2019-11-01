package main

import (
	"fmt"
	"log"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

// parseServiceAreas: parser for: 1b) Service Areas
var parseServiceAreas processXlsxSheet = func(params paramConfig, sheetIndex int, tableFromSliceCreator services.TableFromSliceCreator) error {
	// XLSX Sheet consts
	const xlsxDataSheetNum int = 4          // 1b) Service Areas
	const serviceAreaRowIndexStart int = 10 // start at row 10 to get the rates
	const basePointCityColumn int = 2
	const stateColumn int = 3
	const serviceAreaNumberColumn int = 4
	const zip3sColumn int = 5
	const internationalRateAreaColumn int = 9
	const rateAreaIDColumn int = 10

	if xlsxDataSheetNum != sheetIndex {
		return fmt.Errorf("parseServiceAreas expected to process sheet %d, but received sheetIndex %d", xlsxDataSheetNum, sheetIndex)
	}

	log.Println("Parsing Domestic Service Areas")
	// Create CSV writer to save data to CSV file, returns nil if params.saveToFile=false
	csvWriter := createCsvWriter(params.saveToFile, sheetIndex, params.runTime)
	if csvWriter != nil {
		defer csvWriter.close()
	}

	var domServAreas []models.StageDomesticServiceArea
	dataRows := params.xlsxFile.Sheets[xlsxDataSheetNum].Rows[serviceAreaRowIndexStart:]
	for _, row := range dataRows {
		domServArea := models.StageDomesticServiceArea{
			BasePointCity:     getCell(row.Cells, basePointCityColumn),
			State:             getCell(row.Cells, stateColumn),
			ServiceAreaNumber: getCell(row.Cells, serviceAreaNumberColumn),
			Zip3s:             getCell(row.Cells, zip3sColumn),
		}
		// All the rows are consecutive, if we get to a blank one we're done
		if domServArea.BasePointCity == "" {
			break
		}

		if params.showOutput == true {
			log.Printf("%v\n", domServArea)
		}
		domServAreas = append(domServAreas, domServArea)
	}

	// TODO: Move these two things out of here
	if csvWriter != nil {
		if err := gocsv.MarshalFile(domServAreas, csvWriter.csvFile); err != nil {
			return errors.Wrap(err, "Could not marshal CSV file for domestic service areas")
		}
	}
	if err := tableFromSliceCreator.CreateTableFromSlice(domServAreas); err != nil {
		return errors.Wrap(err, "Could not create temp table for domestic service areas")
	}

	log.Println("Parsing International Service Areas")

	var intlServAreas []models.StageInternationalServiceArea
	for _, row := range dataRows {
		intlServArea := models.StageInternationalServiceArea{
			RateArea:   getCell(row.Cells, internationalRateAreaColumn),
			RateAreaID: getCell(row.Cells, rateAreaIDColumn),
		}
		// All the rows are consecutive, if we get to a blank one we're done
		if intlServArea.RateArea == "" {
			break
		}

		if params.showOutput == true {
			log.Printf("%v\n", intlServArea)
		}
		intlServAreas = append(intlServAreas, intlServArea)
	}

	// TODO: Move these two things out of here
	if csvWriter != nil {
		if err := gocsv.MarshalFile(intlServAreas, csvWriter.csvFile); err != nil {
			return errors.Wrap(err, "Could not marshal CSV file for international service areas")
		}
	}
	if err := tableFromSliceCreator.CreateTableFromSlice(intlServAreas); err != nil {
		return errors.Wrap(err, "Could not create temp table for international service areas")
	}

	return nil
}

// verifyServiceAreas: verification for: 1b) Service Areas
var verifyServiceAreas verifyXlsxSheet = func(params paramConfig, sheetIndex int) error {
	log.Println("TODO verifyServiceAreas() not implemented")
	return nil
}