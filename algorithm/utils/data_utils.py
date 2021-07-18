import random
import pandas as pd
from matplotlib import pyplot as plt
from sklearn.preprocessing import StandardScaler


def normalization(data):
    scaler = StandardScaler()
    if not isinstance(data[0], list):
        xData = [[d] for d in data]
        scaler.fit(xData)
        normalisedData = scaler.transform(xData)
        xNormalisedData = [el[0] for el in normalisedData]
    else:
        scaler.fit(data)
        xNormalisedData = scaler.transform(data)
    return xNormalisedData


def split_data(inputs, outputs):
    random.seed(5)
    indexes = [i for i in range(len(inputs))]
    trainSample = random.choices(indexes, k=int(0.8 * len(inputs)))
    testSample = [i for i in indexes if not i in trainSample]
    trainInputs = [inputs[i] for i in trainSample]
    trainOutputs = [outputs[i] for i in trainSample]
    testInputs = [inputs[i] for i in testSample]
    testOutputs = [outputs[i] for i in testSample]
    return trainInputs, trainOutputs, testInputs, testOutputs


def split_into_batches(input, no_batches):
    for i in range(0, len(input), no_batches):
        yield input[i:i + no_batches]


def balance_histogram(outputs):
    pd.value_counts(outputs).plot.bar()
    plt.title('Fraud class histogram')
    plt.xlabel('Class')
    plt.ylabel('Frequency')
    plt.show()