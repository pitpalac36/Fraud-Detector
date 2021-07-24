import os

from dotenv import load_dotenv


def get_db_info():
    load_dotenv()
    url = str(os.getenv('MONGO_URI'))
    db_name = str(os.getenv('DB'))
    col_name = str(os.getenv('COLLECTION'))
    return url, db_name, col_name


