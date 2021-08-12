import os

from dotenv import load_dotenv


def get_scaler_file():
    if os.getenv("PRODUCTION") != "1":
        load_dotenv()
    return str(os.getenv('SCALER_FILE'))


def get_address_and_port():
    if os.getenv("PRODUCTION") != "1":
        load_dotenv()
    return str(os.getenv('ADDRESS')), str(os.getenv('PORT'))
