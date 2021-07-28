from dataclasses import dataclass
from dataclasses_json import dataclass_json


@dataclass_json
@dataclass
class ResultDTO:
    tran_id: str
    is_fraud: bool