from numpy import asarray


def denormalization(scaler, data):
    if not isinstance(data[0], list):
        denormalisedData = scaler.inverse_transform(asarray([data]))
    else:
        denormalisedData = scaler.inverse_transform(data)
    return denormalisedData[0].tolist()
