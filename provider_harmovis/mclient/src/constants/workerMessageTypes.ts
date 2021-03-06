export interface WorkerMessage {
    type: SocketMsgTypes;
}

export interface SocketMessage<T> extends WorkerMessage  {
    payload: T
}

export type SocketMsgTypes = 'RECEIVED_POINT'
     | 'CONNECTED'
     | 'RECEIVED_MAPBOX_TOKEN'
     | 'RECEIVED_ROUTERESPONSE'
     | 'RECEIVED_VEHICLESTATUSREQUEST'
     | 'RECEIVED_DISPATCHREQUEST'
     | 'RECEIVED_DISPATCHRESPONSE'
     | 'RECEIVED_ROUTEREQUEST'
     | 'RECEIVED_VEHICLESTATUSRESPONSE'
     | 'RECEIVED_DISABILITYINFO'
     | 'RECEIVED_DISABILITYRESPONSE'
     | 'RECEIVED_EVFLEETSUPPLY'
     | 'RECEIVED_VEHICLELIST'
     | 'RECEIVED_EVFLEETRESPONSE'
     | 'RECEIVED_DELIVERYPLANNINGPROVIDE'
     | 'RECEIVED_DISPDISPATCHRESPONSE'
     | 'RECEIVED_DELIVERYPLANNINGREQUEST'
     | 'RECEIVED_DELIVERYPLANADOPTION'
     | 'RECEIVED_DELIVERYPLANNINGRESPONSE'
;
