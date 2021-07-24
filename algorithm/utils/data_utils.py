import pickle
import random
import pandas as pd
from imblearn.over_sampling import SMOTE
from matplotlib import pyplot as plt
from sklearn.preprocessing import StandardScaler


def generate_frauds(trainInputs, trainOutputs):
    sm = SMOTE(random_state=2)
    trainInputs, trainOutputs = sm.fit_resample(trainInputs, trainOutputs)
    return trainInputs, trainOutputs


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
    # save scaler information to pickle file
    with open('scaler_info.pickle', 'wb') as handle:
        pickle.dump(scaler, handle, protocol=pickle.HIGHEST_PROTOCOL)
    return xNormalisedData


def split_data(inputs, outputs):
    random.seed(5)
    indexes = [i for i in range(len(inputs))]
    trainSample = random.choices(indexes, k=int(0.5 * len(inputs)))
    testSample = [i for i in indexes if not i in trainSample]
    trainInputs = [inputs[i] for i in trainSample]
    trainOutputs = [outputs[i] for i in trainSample]
    testInputs = [inputs[i] for i in testSample]
    testOutputs = [outputs[i] for i in testSample]
    return trainInputs, trainOutputs, testInputs, testOutputs


def balance_histogram(outputs):
    pd.value_counts(outputs).plot.bar()
    plt.title('Fraud class histogram')
    plt.xlabel('Class')
    plt.ylabel('Frequency')
    plt.show()
