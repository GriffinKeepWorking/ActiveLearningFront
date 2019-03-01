import { combineReducers } from 'redux'

import imageUrlReducer from './imageReducer'
import countReducer from './countReducer'
import timeReducer from './timeReducer'
import startReducer from './startReducer'

export default combineReducers({
	imageURL: imageUrlReducer,
	count: countReducer,
	time: timeReducer,
  start: startReducer
});


  // const reducers = combineReducers({
  //   count: countReducer,
  //   yesPile: yesPileReducer,
  //   noPile: noPileReducer,

  //   timeArray: timeArrayReducer,

  //   timeBackEnd: timeBackEndReducer,
    
  // })
