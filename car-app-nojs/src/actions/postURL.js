import { CHANGE_COUNT } from './types'

export const postURL = postData => dispatch => {
  fetch('https://jsonplaceholder.typicode.com/posts', {
    method: 'POST',
    headers: {
      'content-type': 'application/json'
    },
    body: JSON.stringify(postData)
  });
};