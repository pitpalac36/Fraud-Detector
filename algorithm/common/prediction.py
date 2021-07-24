

def predict(lr, sample):
    return lr.predict(sample.reshape(1, -1))