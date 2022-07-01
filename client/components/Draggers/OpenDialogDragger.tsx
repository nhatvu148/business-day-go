import { Paper } from "@mui/material";
import React, { FC } from "react";
import Draggable from "react-draggable";
import { connect } from "react-redux";
import { AnyAction, bindActionCreators } from "redux";
import { ThunkDispatch } from "redux-thunk";
import { setDragStopDialog } from "redux/actions/dragActions";
import { IAppState } from "redux/reducers";
import { getDrag } from "redux/selectors";
import { ICoordinate, IDragState } from "redux/types";

interface IStateProps {
  drag: IDragState;
}

interface IDispatchProps {
  dispatch: ThunkDispatch<{}, {}, AnyAction>;
  setDragStopDialog: (
    newData: ICoordinate
  ) => (dispatch: ThunkDispatch<{}, {}, AnyAction>) => void;
}

type IProps = IStateProps & IDispatchProps;

const OpenDialogDragger: FC<IProps> = ({
  drag: { stopDragDialogAt },
  dispatch,
  setDragStopDialog,
  ...props
}) => {
  return (
    <Draggable
      handle="#alert-dialog-title"
      position={{ x: stopDragDialogAt.x, y: stopDragDialogAt.y }}
      onStop={(e, data) => {
        setDragStopDialog({ x: data.x, y: data.y });
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
    ...bindActionCreators(
      {
        setDragStopDialog,
      },
      dispatch
    ),
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(OpenDialogDragger);
