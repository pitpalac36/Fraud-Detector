from common.prediction import predict


def get_confusion_matrix(lr, testInputs, testOutputs):
    false_negative = 0
    false_positive = 0
    true_positive = 0
    true_negative = 0
    for i in range(len(testInputs)):
        real = 'ok' if testOutputs[i] == 0 else 'fraud'
        computed = 'ok' if predict(lr, testInputs[i]) == [0] else 'fraud'
        if computed == real:
            print("computed : " + str(computed) + "     real : " + real)
            if real == 'fraud':
                true_positive += 1
            else:
                true_negative += 1
        else:
            print("computed : " + str(computed) + "     real : " + real + "   WRONG")
            if computed == 'ok' and real == 'fraud':  # false negative
                false_negative += 1
            else:
                false_positive += 1
    return false_negative, false_positive, true_positive, true_negative