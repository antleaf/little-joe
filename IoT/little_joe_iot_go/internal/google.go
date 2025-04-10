package internal

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"
)

func NewGoogleSheetsService() (*sheets.Service, error) {
	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx, sheets.SpreadsheetsScope)
	if err != nil {
		Log.Fatalf("Unable to find default credentials: %v", err)
	}
	service, err := sheets.NewService(ctx, option.WithCredentials(credentials))
	if err != nil {
		Log.Fatalf("Unable to create Sheets service: %v", err)
	}
	return service, err
}

func WriteToGoogleSheet(srv *sheets.Service, spreadsheetId string, writeRange string, values [][]interface{}) error {
	valueRange := &sheets.ValueRange{
		Values: values,
	}
	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, valueRange).ValueInputOption("USER_ENTERED").Do()
	return err
}
