# frozen_string_literal: true
require "google_drive"

session = GoogleDrive::Session.from_service_account_key("./credentials/little-joe-86fe37486bb4.json")
ws = session.spreadsheet_by_key("1wz5gZ2TaHvzxDfvdOJxpuQ0dyAo7iCLEBKRFj87oJAY").worksheets[0]
puts ws[1, 1]
ws.insert_rows(2, [[Time.now.to_s, "16.6"]])
ws.save