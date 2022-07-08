import { combineReducers } from "redux";
import dragReducer from "redux/reducers/dragReducer";
import { IDragState } from "redux/types";

export interface IAppState {
  drag: IDragState;
}

export default combineReducers<IAppState>({
  drag: dragReducer,
});
