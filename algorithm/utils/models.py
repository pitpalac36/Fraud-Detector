from dataclasses import dataclass

@dataclass
class ResultDTO:
    tran_id: str
    is_fraud: bool