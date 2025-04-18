# Little Joe Python Wrapper

This repository provides a Python wrapper for the `little_joe_iot_go` binary, enabling seamless integration of its functionality into Python-based projects.

## Overview

The wrapper acts as an interface to the `little_joe_iot_go` binary, requiring the same environment variables as the binary itself, along with an additional variable to specify the binary's location.

## Requirements

- Python 3.6 or higher
- The `little_joe_iot_go` binary
- Required environment variables for `little_joe_iot_go`

## Environment Variables

Ensure the following environment variables are set:

- **Binary Location**: `LITTLE_JOE_BINARY_PATH`  
   Path to the `little_joe_iot_go` binary.

- **Other Variables**: All environment variables required by the `little_joe_iot_go` binary.

## Usage

Run the script

```python
python main.py
```

## Install as SystemD service

To install `main.py` as a SystemD service, follow these steps:

1. Copy the service and timer files from the `systemd` folder to the SystemD directory:

   ```bash
   sudo cp systemd/little_joe.service /etc/systemd/system/
   sudo cp systemd/little_joe.timer /etc/systemd/system/
   ```

2. Reload the SystemD daemon to recognize the new service:

   ```bash
   sudo systemctl daemon-reload
   ```

3. Enable the service to start on boot:

   ```bash
   sudo systemctl enable little_joe.timer
   ```

4. Start the service:

   ```bash
   sudo systemctl start little_joe.timer
   ```

5. Check the status of the service to ensure it is running:
   ```bash
   sudo systemctl list-timers --all
   ```

Make sure the `LITTLE_JOE_BINARY_PATH` and other required environment variables are properly set in the service file or the environment.
