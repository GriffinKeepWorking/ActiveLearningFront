import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import queryString from 'query-string';
import './App.css';

import Image from './components/image'; //Component
import store from './store'; //store
import { Provider } from 'react-redux'; //provider

class Child extends Component {

  // componentDidMount() {
  //   var id = 0;
  //   console.log("Search = ", this.props.location.search)
  // }

  render() {
    const id = queryString.parse(this.props.location.search).id
    console.log("id = ", id)
    // console.log("Inside app.js")
    return (

      <React.Fragment>

      <div className="container-fluid">

          <Provider className = "h80" store = { store } >
            <Image data = { id } />
          </Provider>
        
      </div>
      </React.Fragment>
    );
  }
}

class App extends Component {

  render() {
    return (

        <Router>
          <Route path='/' component = {Child} />
        </Router>

      )
  }
  
}

export default App;


