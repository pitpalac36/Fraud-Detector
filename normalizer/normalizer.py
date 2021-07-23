from sklearn.preprocessing import StandardScaler


def normalization(data):
    scaler = StandardScaler()
    if not isinstance(data[0], list):
        xData = [[d] for d in data]
        scaler.fit(xData)
        normalisedData = scaler.transform(xData)
        xNormalisedData = [el[0] for el in normalisedData]
    else:
        xNormalisedData = scaler.fit_transform(data)
    return xNormalisedData