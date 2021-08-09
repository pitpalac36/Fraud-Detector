import os

from dotenv import load_dotenv


def get_scaler_file():
    load_dotenv()
    return str(os.getenv('SCALER_FILE'))


def get_regressor_file():
    load_dotenv()
    return str(os.getenv('REGRESSOR_FILE'))


def get_address():
    load_dotenv()
    return str(os.getenv('AI_ADDRESS'))


def get_port():
    load_dotenv()
    return str(os.getenv('AI_PORT'))


def get_aggregator_url():
    load_dotenv()
    return str(os.getenv('AGGREGATOR_URL'))
