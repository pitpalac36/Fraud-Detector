from pymongo import MongoClient


def read_from_db(url, db_name, col_name):
    client = MongoClient(url)
    db = client[db_name]
    collection = db[col_name]
    documents = collection.find({}).limit(9500)
    inputs = []
    outputs = []
    for each in documents:
        tran = []
        for i in range(1, 29):
            tran.append(each['v{}'.format(i)])
        tran.append(each['amount'])
        tran.append(each['class'])
        inputs.append([float(x) for x in tran[0:29]])
        outputs.append(int(tran[29]))
    return inputs, outputs
