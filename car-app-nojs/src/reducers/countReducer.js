import { INITIALISE_COUNT, CHANGE_COUNT } from '../actions/types';

const initialState = {
  count: 0
};

export default function(state = initialState.count, action) {
  console.log("Change count reducer = ", action.type)
  switch (action.type) {
    case INITIALISE_COUNT:
      console.log("Returning 0 for count")
      return 0
    case CHANGE_COUNT:
      return state + 1
    default:
      return state;
  }
}