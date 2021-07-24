import os

from dotenv import load_dotenv


def get_scaler_file():
    load_dotenv()
    return str(os.getenv('SCALER_FILE'))


def get_regressor_file():
    load_dotenv()
    return str(os.getenv('REGRESSOR_FILE'))