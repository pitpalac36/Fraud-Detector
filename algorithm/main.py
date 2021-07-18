from sklearn.linear_model import LogisticRegression
from utils.file_utils import read_from_csv, write_to_csv
from utils.data_utils import normalization, split_data, balance_histogram
from imblearn.over_sampling import SMOTE
from pymongo import MongoClient
from pprint import pprint

from utils.metrics_utils import get_confusion_matrix


def generate_frauds(trainInputs, trainOutputs):
    sm = SMOTE(random_state=2)
    trainInputs, trainOutputs = sm.fit_resample(trainInputs, trainOutputs)
    header = ['V1', 'V2','V3','V4','V5','V6','V7','V8','V9','V10','V11','V12','V13','V14','V15','V16','V17','V18','V19','V20','V21','V22','V23','V24','V25','V26','V27','V28','Amount', 'Class']
    # write_to_csv('data/generated.csv', header, trainInputs)
    return trainInputs, trainOutputs


def main():
    # read data from csv/db
    inputs, outputs = read_from_csv('data/creditcard.csv')
    # view data balance before oversampling
    balance_histogram(outputs)
    # normalise data
    normalisedInputs = normalization(inputs)

    print(inputs[0])
    print(normalization(inputs[0]))

    # split data into train inputs and train outputs
    trainInputs, trainOutputs, testInputs, testOutputs = split_data(normalisedInputs, outputs)

    print("Before OverSampling, frauds: {}".format(sum(trainOutputs[i] == 1 for i in range(len(trainOutputs)))))
    print("Before OverSampling, normal: {} \n".format(sum(trainOutputs[i] == 0 for i in range(len(trainOutputs)))))

    # generate frauds
    finalTrainInputs, finalTrainOutputs = generate_frauds(trainInputs, trainOutputs)

    print("After OverSampling, frauds: {}".format(sum(finalTrainOutputs[i] == 1 for i in range(len(finalTrainOutputs)))))
    print("After OverSampling, normal: {} \n".format(sum(finalTrainOutputs[i] == 0 for i in range(len(finalTrainOutputs)))))

    # view data balance after oversampling
    balance_histogram(finalTrainOutputs)
    # initialize logistic regressor with l2 penalty as cost function
    lr = LogisticRegression(penalty='none')
    # train regressor
    lr.fit(finalTrainInputs, finalTrainOutputs)

    # get confusion matrix elements
    false_negative, false_positive, true_positive, true_negative = get_confusion_matrix(lr, testInputs, testOutputs)
    # view performance resume
    print("test data size: " + str(len(testInputs)))
    print("predicted correctly: " + str(len(testInputs) - false_positive - false_negative))
    print("predicted wrong: " + str(false_positive + false_negative))
    print("false positive: " + str(false_positive))
    print("false negative (very bad!): " + str(false_negative))
    # view performance metrics
    # accuracy (weak metric; assumes equal costs for both kinds of errors)
    accuracy = (true_positive + true_negative) / len(testInputs)
    print("accuracy: " + str(accuracy))
    # precision (what proportion of predicted positives is truly positive - important)
    precision = true_positive / (true_positive + false_positive)
    print("precision: " + str(precision))
    # recall (what proportion of actual positives is correctly classified - very important)
    recall = (true_positive) / (true_positive + false_negative)
    print("recall: " + str(recall))
    # f1_score (a tradeoff between precision and recall)
    f1_score = 2 * (precision * recall) / (precision + recall)
    print("f1 score: " + str(f1_score))

if __name__ == '__main__':
    main()






























