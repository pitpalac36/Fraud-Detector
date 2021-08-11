import pickle

from utils.env_utils import get_scaler_file
from numpy import asarray


def normalization(scaler, data):
    if not isinstance(data[0], list):
        normalisedData = scaler.transform(asarray([data]))
    else:
        normalisedData = scaler.transform(data)
    return normalisedData[0].tolist()
