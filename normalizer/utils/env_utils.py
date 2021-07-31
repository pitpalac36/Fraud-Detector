import os

from dotenv import load_dotenv


def get_scaler_file():
    load_dotenv()
    return str(os.getenv('SCALER_FILE'))


def get_address_and_port():
    load_dotenv()
    return str(os.getenv('ADDRESS')), str(os.getenv('PORT'))
