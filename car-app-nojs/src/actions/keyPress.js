import { CHANGE_COUNT, CHANGE_TIME } from './types.js'

export const keyPress = (key, image, time, id) => dispatch => {

	let pile = ""
	switch(key) {
		case "ArrowLeft":
			pile = "yes"
			break;
		case "ArrowRight":
			pile = "no";
			break;
	}

	if (pile !== "") {
		console.log('Sending Image: ', image)
		console.log('Pile: ', pile)
		console.log("Time: ", new Date() - time)
		console.log("ID: ", id)

		let data = {
			url: image,
			type: pile,
			time: new Date() - time,
			id: parseInt(id)
		}
		
	  // fetch('https://ui.dv-api.com/codetest?ip=76.125.94.0')

	  fetch('http://localhost:8080/postImages', {
	    method: 'POST',
	    headers: {
	      'content-type': 'application/json'
	    },
	    body: JSON.stringify(data),
	    mode: 'no-cors'
	  })
	    .then(
	    	dispatch({
	    		type: CHANGE_COUNT,
	    		payload: 1
	    	})
	    ).then(dispatch({
	    	type: CHANGE_TIME
	    }));


	  // fetch('http://localhost:8080/fetchImages')
	  //   .then(res => res.json())
	  //   .then(posts => {
		 //    console.log("posts: ", posts.url)
		 //      dispatch({
		 //        type: FETCH_URL,
		 //        payload: posts.url
		 //      })
		 //  }
	  //   ).then(dispatch({
	  //   	type: INITIALISE_COUNT,
	  //   	payload: 0
	  //   }));
	}
};