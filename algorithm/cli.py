from sklearn.linear_model import LogisticRegression
from utils.db import read_from_db
from utils.data import normalization, split_data, balance_histogram, generate_frauds
from utils.env import get_db_info
from common.env import get_regressor_file
from utils.metrics import get_confusion_matrix
import pickle


def train_model(trainInputs, trainOutputs):
    print("Before OverSampling, frauds: {}".format(sum(trainOutputs[i] == 1 for i in range(len(trainOutputs)))))
    print("Before OverSampling, normal: {} \n".format(sum(trainOutputs[i] == 0 for i in range(len(trainOutputs)))))

    # generate frauds
    finalTrainInputs, finalTrainOutputs = generate_frauds(trainInputs, trainOutputs)

    print(
        "After OverSampling, frauds: {}".format(sum(finalTrainOutputs[i] == 1 for i in range(len(finalTrainOutputs)))))
    print("After OverSampling, normal: {} \n".format(
        sum(finalTrainOutputs[i] == 0 for i in range(len(finalTrainOutputs)))))

    # view data balance after oversampling
    balance_histogram(finalTrainOutputs)
    # initialize logistic regressor with l2 penalty as cost function
    lr = LogisticRegression(penalty='l2')
    # train regressor
    lr.fit(finalTrainInputs, finalTrainOutputs)

    # save regressor information to pickle file
    lr_file = get_regressor_file()
    with open(lr_file, 'wb') as handle:
        pickle.dump(lr, handle, protocol=pickle.HIGHEST_PROTOCOL)


def test_model(testInputs, testOutputs):
    # get regressor info from pickle file
    lr_file = get_regressor_file()
    with open(lr_file, 'rb') as handle:
        lr = pickle.load(handle)

    # get confusion matrix elements
    false_negative, false_positive, true_positive, true_negative = get_confusion_matrix(lr, testInputs, testOutputs)

    # view performance resume
    print("test data size: " + str(len(testInputs)))
    print("predicted correctly: " + str(len(testInputs) - false_positive - false_negative))
    print("predicted wrong: " + str(false_positive + false_negative))
    print("false positive: " + str(false_positive))
    print("false negative: " + str(false_negative))

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
    # get db info from env file
    url, db_name, col_name = get_db_info()
    # read data from db
    inputs, outputs = read_from_db(url, db_name, col_name)
    # view data balance before oversampling
    balance_histogram(outputs)
    # normalise data
    normalisedInputs = normalization(inputs)
    # split data into train inputs and train outputs
    trainInputs, trainOutputs, testInputs, testOutputs = split_data(normalisedInputs, outputs)

    train_model(trainInputs, trainOutputs)
    test_model(testInputs, testOutputs)
