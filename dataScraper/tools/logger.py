import logging

# Create a logger
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


# Create a StreamHandler to print log messages to stdout
stream_handler = logging.StreamHandler()
stream_handler.setLevel(logging.INFO)


# Create a formatter and set the format for the log messages
formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')
stream_handler.setFormatter(formatter)


# Add the StreamHandler to the logger
logger.addHandler(stream_handler)