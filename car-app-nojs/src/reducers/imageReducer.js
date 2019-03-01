import { FETCH_URL } from '../actions/types';

const initialState = {
  // count: 0,
	imageURL: [],
  // count: 0
};

export default function(state = initialState.imageURL, action) {
  console.log("Action = ", action.type)
  switch (action.type) {
    case FETCH_URL:
      console.log("Inside fetch url")
      let value = action.payload;
      // var result = Object.keys(value).map(function(key) {
      //     return [Number(key)];
      //   });
      // console.log("result type = ", typeof result)
      console.log("Incoming payload type = ", typeof value)
      console.log("state in reducer = ", state)
      let arr = []
      // arr = arr.push(state)
      arr = arr.unshift(value)
      console.log("Arr after pushing = ", arr)
      return value
      // break;
    default:
      return state;
      // break;
  }
}
