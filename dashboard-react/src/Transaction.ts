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

export function jsonToTransaction(j : any) {
    return {
        source: j.Source,
        amount: j.Amount,
        v1: j.V1,
        v2: j.V2,
        v3: j.V3,
        v4: j.V4,
        v5: j.V5,
        v6: j.V6,
        v7: j.V7,
        v8: j.V8,
        v9: j.V9,
        v10: j.V10,
        v11: j.V11,
        v12: j.V12,
        v13: j.V13,
        v14: j.V14,
        v15: j.V15,
        v16: j.V16,
        v17: j.V17,
        v18: j.V18,
        v19: j.V19,
        v20: j.V20,
        v21: j.V21,
        v22: j.V22,
        v23: j.V23,
        v24: j.V24,
        v25: j.V25,
        v26: j.V26,
        v27: j.V27,
        v28: j.V28
    }
}