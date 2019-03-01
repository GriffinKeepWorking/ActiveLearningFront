# Active Learning System - Frontend Interface

**Daniel and Griffin Version of Visual Factor Graph**

## Install
1. Frameworks and Technologies: React, Redux, Webpack, GoLang, Python
2. run "apt install nodejs npm"
3. run "./dependencies.txt"

## Run GoLang Server for local testing
1. cd React/car-app-nojs
2. run "go get github.com/gorilla/mux"
3. run "go run server.go" 
4. go to "localhost:8080" in your browser and you should see:   

<img src="https://i.ibb.co/b2hGc4X/test.png" width="700" height="400">

5. lt --subdomain small --port 8080

## current implementation workflow
<img src="https://i.ibb.co/RQgwrXL/flow.png" width="700" height="70">

## Uploading image to interface
```
run python3 prepareData.py -i <foldername> 
```
prepare the image list from local address here. It randomly divides all image into tasks with 20 images each, upload them to S3 instance, and generate a tasks folder with txt files that have all the urls.
```
run python3 create_hit_boto3.py --task <tasknumber>
```
It sends an email with link to the HIT to user's email address when the task is published.
access key problem waiting to be solved here.

## Approve the HIT
```
run python3 approve_hits.py -t <threshold>
approve the hit if the validation accuracy is above the threshold.
```


# Development Documentation
high level implementation   
<img src="https://i.ibb.co/dmJzXGk/localserver-Logics.png" width="700" height="400">   

low level code explaination:   
<img src="https://i.ibb.co/L1FjVJT/codes.png" width="700" height="500">    
   
./src:   
*.css: stylesheet contains style information of classes including their location, sizes or colors.
index.html: contains UI components including the time counter, sidebar, header and React component. It also has mturk form(returned in the interation with mturk api).   
app.js: contains the "image.js" component and the store that brings corresponding action and reducer together.   
store.js: creates and initializes the store info.   
   
./components:   
image.js: loads all images and displays them in a centered self-adjusting style, logics in interation with other actions.   
   
./actions:   
fetchURL.js: featches URL from server and handles format.   
keyPress.js: detects input key presses and updates time when input comes.   
postURL.js: wraps urls into json format.   
types/js(trivial): defines global constant.   
      
./reducers:   
countReducer.js: updates processed images count for dashboard purpose.   
imageReducer.js: updates image states in the store when fetching urls is occuring.   
startReducer.js: stores a start time so time spent on image can be later calculated by time difference.   
timeReducer.js: interprates time frame and designs to facilitate changes in the future.   
index.js(trivial): combines above reducers into one like a header file.   
   
graph visualization of actions, reducers and store:   
<img src="https://i.ibb.co/2MvcV8S/actions-and-store.png" width="700" height="400">

extra: dashboard implementation is still incomplete, but basically it takes the states in redux store in real-time and shows graphs based on them in another React app component.

## Code style guides

* [Javascript Style Guide](https://google.github.io/styleguide/javascriptguide.xml)
* [Golang Style Guide](https://github.com/golang/go/wiki/CodeReviewComments)
* [Python Style Guide](https://www.python.org/dev/peps/pep-0008/)
