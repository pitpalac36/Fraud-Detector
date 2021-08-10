export default interface Transaction {
    source: string;
    amount: number;
    v1: number;
    v2: number;
    v3: number;
    v4: number;
    v5: number;
    v6: number;
    v7: number;
    v8: number;
    v9: number;
    v10: number;
    v11: number;
    v12: number;
    v13: number;
    v14: number;
    v15: number;
    v16: number;
    v17: number;
    v18: number;
    v19: number;
    v20: number;
    v21: number;
    v22: number;
    v23: number;
    v24: number;
    v25: number;
    v26: number;
    v27: number;
    v28: number;
}

export function processData(t: any) {
    t.amount = Math.round(t.amount * 100) / 100;
    var decoded = atob(t.source).split("@", 2);
    t.source = decoded[0];
    t.timestamp = decoded[1];
    return t;
}