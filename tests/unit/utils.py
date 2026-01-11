import os
import logging
from typing import Dict, List, Tuple

# Set up logging configuration
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def parse_file_path(file_path: str) -> Tuple[str, str]:
    """
    Parse the file path into a directory path and a file name.
    
    Args:
        file_path (str): The full file path.
    
    Returns:
        Tuple[str, str]: A tuple containing the directory path and the file name.
    """
    directory_path = os.path.dirname(file_path)
    file_name = os.path.basename(file_path)
    return directory_path, file_name

def load_config(config_file_path: str) -> Dict:
    """
    Load the configuration from a JSON file.
    
    Args:
        config_file_path (str): The path to the configuration file.
    
    Returns:
        Dict: A dictionary containing the configuration.
    """
    import json
    try:
        with open(config_file_path, 'r') as config_file:
            config = json.load(config_file)
            return config
    except FileNotFoundError:
        logger.error(f"Configuration file not found: {config_file_path}")
        raise

def split_list(input_list: List, chunk_size: int) -> List[List]:
    """
    Split a list into chunks of a specified size.
    
    Args:
        input_list (List): The list to be split.
        chunk_size (int): The size of each chunk.
    
    Returns:
        List[List]: A list of chunks.
    """
    return [input_list[i:i + chunk_size] for i in range(0, len(input_list), chunk_size)]

def main():
    # Example usage:
    file_path = "/path/to/example.txt"
    directory_path, file_name = parse_file_path(file_path)
    logger.info(f"Directory path: {directory_path}, File name: {file_name}")

    config_file_path = "/path/to/config.json"
    config = load_config(config_file_path)
    logger.info(f"Configuration: {config}")

    input_list = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    chunk_size = 3
    chunks = split_list(input_list, chunk_size)
    logger.info(f"Chunks: {chunks}")

if __name__ == "__main__":
    main()