import subprocess
import os

def update_soil_moisture_data(reading: float, debug: bool = True):
    """
    Executes the little_joe binary to update soil moisture data.

    Args:
        reading (float): The soil moisture reading to be passed to the binary.
        debug (bool): Whether to enable debug mode. Defaults to True.

    Raises:
        FileNotFoundError: If the binary file is not found.
        subprocess.CalledProcessError: If the binary execution fails.
    """
    # Read binary path from environment variable
    binary_path = os.getenv("LITTLE_JOE_BINARY_PATH")
    if binary_path is None:
        # If the environment variable is not set, use a default path
        print("LITTLE_JOE_BINARY_PATH not set, using default path.")
        binary_path = "./little_joe"
    else:
        print(f"LITTLE_JOE_BINARY_PATH set to: {binary_path}")  

    # Check if the binary path is valid
    if not os.path.isfile(binary_path):
        raise FileNotFoundError(f"Invalid binary path: {binary_path}")

    # Check if the binary file exists
    try:
        with open(binary_path, 'rb') as f:
            pass
    except FileNotFoundError:
        raise FileNotFoundError(f"Binary file not found: {binary_path}")

    # Execute the binary file
    try:
        debug_flag = "--debug=true" if debug else "--debug=false"
        cmd = f"{binary_path} updateSoilMoistureData {debug_flag} --reading={reading}"
        print(f"Executing command: {cmd}")
        result = subprocess.run(cmd, shell=True, check=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        print("Output:", result.stdout)
        return result.stdout
    except subprocess.CalledProcessError as e:
        print("Error:", e.stderr)
        raise e