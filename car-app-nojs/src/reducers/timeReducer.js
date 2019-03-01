import { INITIALISE_TIME, CHANGE_TIME } from '../actions/types';

const initialState = {
  time: Number(new Date())
};

export default function(state = initialState.time, action) {
  console.log("Change time reducer = ", action.type)
  switch (action.type) {
    case INITIALISE_TIME:
      // console.log("Returning 0 for time")
      return Number(new Date())
    case CHANGE_TIME: // Both are same for now. Designed to facilitate changes in future
      return Number(new Date())
    default:
      return state;
  }
}