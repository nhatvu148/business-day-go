import { Paper } from "@mui/material";
import React, { FC } from "react";
import Draggable from "react-draggable";
import { connect } from "react-redux";
import { AnyAction } from "redux";
import { ThunkDispatch } from "redux-thunk";
import { IAppState } from "redux/reducers";
import { getDrag } from "redux/selectors";
import { ActionTypes, IDragState } from "redux/types";

interface IStateProps {
  drag: IDragState;
}

interface IDispatchProps {
  dispatch: ThunkDispatch<{}, {}, AnyAction>;
}

type IProps = IStateProps & IDispatchProps;

const OpenDialogDragger: FC<IProps> = ({
  drag: { stopDragDialogAt },
  dispatch,
  ...props
}) => {
  return (
    <Draggable
      handle="#alert-dialog-title"
      position={{ x: stopDragDialogAt.x, y: stopDragDialogAt.y }}
      onStop={(e, data) => {
        dispatch({
          type: ActionTypes.DRAG_STOP_DIALOG,
          x: data.x,
          y: data.y,
        });
      }}
    >
      <Paper {...props} />
    </Draggable>
  );
};

const mapStateToProps = (state: IAppState): IStateProps => ({
  drag: getDrag(state),
});

const mapDispatchToProps = (dispatch: ThunkDispatch<{}, {}, AnyAction>) => {
  return {
    dispatch,
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(OpenDialogDragger);
