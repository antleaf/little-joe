# CLI utility for sending data to Google sheet

This utility depends on several environment settings:

- When executing the binary, these must be exported as normal in the OS's environment
- When developing with the source code, these can either be exported as normal in the OS's environment, or specified in the `.env` file. 

Settings from the operating system environment **take precedence** over anything set in `.env`.

This utility requires a separate Google credentials JSON file at a path specified in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable

To execute the binary:

1. Provide a Google credentials JSON file somewhere in the local file system
2. Export an environment variable called `GOOGLE_APPLICATION_CREDENTIALS` with the path to the Google credentials JSON file
3. Export the other environment variables and values detailed in `.env`
4. Run the executable

## Example
```bash
cd binaries/linux

export GOOGLE_APPLICATION_CREDENTIALS="<SOME_PATH>/little-joe-google-api-service-key.json" && \
  export LITTLE_JOE_GOOGLE_SHEET_ID="1wz5gZ2TaHvzxDfvdOJxpuQ0dyAo7iCLEBKRFj87oJAY" && \
  export LITTLE_JOE_GOOGLE_SHEET_WORKSHEET="Soil Moisture" && \
  ./littlejoe updateSoilMoistureData --debug=true --reading=15.7
```


The Google sheet is here:
https://docs.google.com/spreadsheets/d/1wz5gZ2TaHvzxDfvdOJxpuQ0dyAo7iCLEBKRFj87oJAY/edit?gid=0#gid=0
