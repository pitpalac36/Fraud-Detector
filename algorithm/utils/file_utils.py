import csv

def read_from_csv(fileName):
    data = []
    with open(fileName) as csv_file:
        csv_reader = csv.reader(csv_file, delimiter=',')
        line_count = 0
        for row in csv_reader:
            print(line_count)
            if line_count == 10000:
                break
            if line_count != 0:
                data.append(row)
            line_count += 1
    inputs = [[float(x) for x in data[i][1:30]] for i in range(len(data))]
    outputs = [int(data[i][30]) for i in range(len(data))]
    return inputs, outputs


def write_to_csv(filename, header, input):
    with open(filename, mode='x') as csv_file:
        csv_writer = csv.writer(csv_file, delimiter=',')
        csv_writer.writerow(header)
        for each in input:
            csv_writer.writerow(each)