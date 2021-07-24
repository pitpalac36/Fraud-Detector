import pickle

from utils.env_utils import get_scaler_file
from numpy import asarray


def normalization(data):
    scaler_file = get_scaler_file()
    with open(scaler_file, 'rb') as handle:
        scaler = pickle.load(handle)
    if not isinstance(data[0], list):
        normalisedData = scaler.transform(asarray([data]))
    else:
        normalisedData = scaler.transform(data)
    return normalisedData[0].tolist()
