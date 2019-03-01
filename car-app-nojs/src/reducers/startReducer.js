import { INITIALISE_START, CHANGE_START } from '../actions/types';

const initialState = {
  start: false
};

export default function(state = initialState.start, action) {
  console.log("Change count reducer = ", action.type)
  switch (action.type) {
    case INITIALISE_START:
      return false
    case CHANGE_START:
      return !state
    default:
      return state;
  }
}