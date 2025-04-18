import update_sheet

def main():
    reading = 1.0
    debug = True
    # Update soil moisture data
    update_sheet.update_soil_moisture_data(reading, debug)

if __name__ == "__main__":
    main()