import os

from dotenv import load_dotenv


def maybe_load_dotenv():
    if os.getenv('PRODUCTION') is not '1':
        load_dotenv()


def get_scaler_file():
    maybe_load_dotenv()
    return str(os.getenv('SCALER_FILE'))


def get_regressor_file():
    maybe_load_dotenv()
    return str(os.getenv('REGRESSOR_FILE'))


def get_address():
    maybe_load_dotenv()
    return str(os.getenv('AI_ADDRESS'))


def get_port():
    maybe_load_dotenv()
    return str(os.getenv('AI_PORT'))


def get_aggregator_url():
    maybe_load_dotenv()
    return str(os.getenv('AGGREGATOR_URL'))
