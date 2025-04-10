# frozen_string_literal: true
require "google_drive"

spreadsheet_id = "1wz5gZ2TaHvzxDfvdOJxpuQ0dyAo7iCLEBKRFj87oJAY"
google_api_service_key_file_path = "/Users/paulwalk/Library/CloudStorage/Dropbox/WorkingDocuments/AntLeaf/Projects/Little Joe/credentials/little-joe-google-api-service-key.json"
worksheet_index = 0

session = GoogleDrive::Session.from_service_account_key(google_api_service_key_file_path)
ws = session.spreadsheet_by_key(spreadsheet_id).worksheets[worksheet_index]
ws.insert_rows(2, [[Time.now.to_s, "17.2"]])
ws.save