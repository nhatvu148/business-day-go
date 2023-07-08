export enum ActionTypes {
  DRAG_STOP_DIALOG,
}

export interface ICoordinate {
  x: number;
  y: number;
}

export interface IDragState {
  stopDragDialogAt: ICoordinate;
}

export type TDragAction = {
  type: ActionTypes.DRAG_STOP_DIALOG;
  stopDragDialogAt: ICoordinate;
};
