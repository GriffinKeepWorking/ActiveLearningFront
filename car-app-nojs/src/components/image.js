import React, { Component } from 'react';
import { connect } from 'react-redux';
import { fetchURL } from '../actions/fetchURL'
import { keyPress } from '../actions/keyPress'
import PropTypes from 'prop-types';
import './image.css'

class Image extends Component {

  componentWillMount() {
    console.log('Mounting');
    // this.props.fetchURL(this.props.data);
  }

  handleKeyPress = (event) => {
    console.log("Inside event")
    console.log("Count: ", this.props.count)
    if (!this.props.start && (event.key === "ArrowLeft" || event.key === "ArrowRight")) {
      console.log("Changing start for event", event.key);
      this.props.fetchURL(this.props.data);
    }
         
    if (this.props.start && this.props.count < this.props.imageURL.length)
      this.props.keyPress(event.key, 
        this.props.imageURL[this.props.count], 
        this.props.time,
        this.props.data);
    else 
      console.log("Limit exceeded")
  }

  importAll(r) {
    let images = {};
    r.keys().map((item, index) => { images[item.replace('./', '')] = r(item); });
    return images;
  }

  render() {

    console.log("Props in component: ")

    console.log("Direct object = ", this.props)
    
    return (

      <div className ="row"
        onKeyDown = {this.handleKeyPress} 
        tabIndex = "-1">

        {
          this.props.start ? 

            this.props.count < this.props.imageURL.length ?
              <React.Fragment>
                
                <div 
                  className = "middlepane">
                  <img 
                    className = "image centered"
                    src = {this.props.imageURL[this.props.count]} 
                  />
                </div>
              </React.Fragment>
              :
              <h1 className = "centered"> Thank you for participating </h1>

              :

              <h1 className = "centered"> Click the left or the right arrows to start</h1>
        }

      </div>
    );
  }

}

Image.propTypes = {
  fetchURL: PropTypes.func.isRequired,
  keyPress: PropTypes.func.isRequired,
  imageURL: PropTypes.array.isRequired,
  count: PropTypes.number.isRequired,
  time: PropTypes.number.isRequired,
  start: PropTypes.bool.isRequired
};

const mapStateToProps = state => ({
  imageURL: state.imageURL,
  count: state.count,
  time: state.time,
  start: state.start
});

export default connect(mapStateToProps, { fetchURL, keyPress })(Image);
