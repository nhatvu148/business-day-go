import { ThunkDispatch } from "redux-thunk";
import { AnyAction } from "redux";
import { ActionTypes, ICoordinate } from "redux/types";

export const setDragStopDialog =
  (newData: ICoordinate) => (dispatch: ThunkDispatch<{}, {}, AnyAction>) => {
    try {
      dispatch({
        type: ActionTypes.DRAG_STOP_DIALOG,
        stopDragDialogAt: newData,
      });
    } catch (error) {
      console.log(error);
    }
  };
