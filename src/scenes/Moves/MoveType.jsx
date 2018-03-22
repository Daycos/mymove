import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import PropTypes from 'prop-types';

import { setPendingMoveType } from './ducks';
import carGray from 'shared/icon/car-gray.svg';
import trailerGray from 'shared/icon/trailer-gray.svg';
import truckGray from 'shared/icon/truck-gray.svg';
import './MoveType.css';

export class BigButton extends Component {
  render() {
    let className = 'ppm-size-button';
    if (this.props.selected) {
      className += ' selected';
    }
    return (
      <div className={className} onClick={this.props.onButtonClick}>
        <p>{this.props.firstLine}</p>
        <p>{this.props.secondLine}</p>
        <img className="icon" src={this.props.icon} alt={this.props.altTag} />
      </div>
    );
  }
}

BigButton.propTypes = {
  firstLine: PropTypes.string.isRequired,
  secondLine: PropTypes.string.isRequired,
  icon: PropTypes.string.isRequired,
  altTag: PropTypes.string.isRequired,
  selected: PropTypes.bool,
  onButtonClick: PropTypes.func,
};

export class BigButtonGroup extends Component {
  render() {
    var createButton = (value, firstLine, secondLine, icon, altTag) => {
      var onButtonClick = () => {
        this.props.onMoveTypeSelected(value);
      };
      return (
        <BigButton
          value={value}
          firstLine={firstLine}
          secondLine={secondLine}
          icon={icon}
          altTag={altTag}
          selected={this.props.selectedOption === value}
          onButtonClick={onButtonClick}
        />
      );
    };
    var small = createButton(
      'S',
      'A few items in your car?',
      '(approx 100 - 800 lbs)',
      carGray,
      'car-gray',
    );
    var medium = createButton(
      'M',
      'A trailer full of household goods?',
      '(approx 400 - 1,200 lbs)',
      trailerGray,
      'trailer-gray',
    );
    var large = createButton(
      'L',
      'A moving truck that you rent yourself?',
      '(approx 1,000 - 5,000 lbs)',
      truckGray,
      'truck-gray',
    );

    return (
      <div>
        <div className="usa-width-one-third">{small}</div>
        <div className="usa-width-one-third">{medium}</div>
        <div className="usa-width-one-third">{large}</div>
      </div>
    );
  }
}
BigButtonGroup.propTypes = {
  selectedOption: PropTypes.string,
  onMoveTypeSelected: PropTypes.func,
};

export class MoveType extends Component {
  componentDidMount() {
    document.title = 'Transcom PPP: Move Type Selection';
  }

  onMoveTypeSelected = value => {
    this.props.setPendingMoveType(value);
  };
  render() {
    const { pendingMoveType, currentMove } = this.props;
    const selectedOption =
      pendingMoveType || (currentMove && currentMove.selected_move_type);
    return (
      <div className="usa-grid-full">
        <BigButtonGroup
          selectedOption={selectedOption}
          onMoveTypeSelected={this.onMoveTypeSelected}
        />
      </div>
    );
  }
}

MoveType.propTypes = {
  pendingMoveType: PropTypes.string,
  currentMove: PropTypes.shape({
    id: PropTypes.string,
    size: PropTypes.string,
  }),
  setPendingMoveType: PropTypes.func.isRequired,
};

function mapStateToProps(state) {
  return state.submittedMoves;
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators({ setPendingMoveType }, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(MoveType);
