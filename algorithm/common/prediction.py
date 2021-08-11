def predict(lr, sample):
    return lr.predict(sample.reshape(1, -1))


def predict2(lr, sample):
    return lr.predict([sample])
