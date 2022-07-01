export interface IAppState {
  drag: IDragState;
}

export enum ActionTypes {
  DRAG_STOP_DIALOG,
}

export interface IDragState {
  stopDragDialogAt: { x: number; y: number };
}

export type TDragAction = {
  type: ActionTypes.DRAG_STOP_DIALOG;
  x: number;
  y: number;
};
