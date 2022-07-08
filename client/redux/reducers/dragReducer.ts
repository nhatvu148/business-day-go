import { ActionTypes, IDragState, TDragAction } from "redux/types";

const initialState: IDragState = {
  stopDragDialogAt: { x: 0, y: 0 },
};

const dragReducer = (state = initialState, action: TDragAction) => {
  switch (action.type) {
    case ActionTypes.DRAG_STOP_DIALOG:
      return {
        ...state,
        stopDragDialogAt: action.stopDragDialogAt,
      };

    default:
      return state;
  }
};

export default dragReducer;
