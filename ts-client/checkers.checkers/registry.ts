import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/checkers/checkers/query";
import { QueryParamsResponse } from "./types/checkers/checkers/query";
import { MsgUpdateParamsResponse } from "./types/checkers/checkers/tx";
import { GenesisState } from "./types/checkers/checkers/genesis";
import { MsgUpdateParams } from "./types/checkers/checkers/tx";
import { Params } from "./types/checkers/checkers/params";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/checkers.checkers.QueryParamsRequest", QueryParamsRequest],
    ["/checkers.checkers.QueryParamsResponse", QueryParamsResponse],
    ["/checkers.checkers.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/checkers.checkers.GenesisState", GenesisState],
    ["/checkers.checkers.MsgUpdateParams", MsgUpdateParams],
    ["/checkers.checkers.Params", Params],
    
];

export { msgTypes }