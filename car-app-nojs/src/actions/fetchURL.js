import { FETCH_URL, INITIALISE_COUNT, INITIALISE_TIME, CHANGE_START } from './types';

export const fetchURL = ( id ) => dispatch => {
  console.log('Fetching from server')
  // fetch('https://ui.dv-api.com/codetest?ip=76.125.94.0')
  // fetch('https://small.localtunnel.me/fetchImages?id=' + id)
  fetch('http://localhost:8080/fetchImages?id=' + id)
    .then(res =>
      res.json()
      // console.log(res.json())
      // console.log(res)
    )
    .then(posts => 
	    {
        console.log("posts: ", posts.url)
            // console.log("Calling dispatch")
              dispatch ({
                type: FETCH_URL,
                payload: posts.url
              })
      }
    )
    .then(dispatch({
    	type: INITIALISE_COUNT,
    	payload: 0
    })).then(dispatch({
      type: INITIALISE_TIME,
      payload: 0
    })).then(dispatch({
      type: CHANGE_START,
      payload: 0
    }));
};

// export const createPost = postData => dispatch => {
//   fetch('https://jsonplaceholder.typicode.com/posts', {
//     method: 'POST',
//     headers: {
//       'content-type': 'application/json'
//     },
//     body: JSON.stringify(postData)
//   })
//     .then(res => res.json())
//     .then(post =>
//       dispatch({
//         type: NEW_POST,
//         payload: post
//       })
//     );
// };