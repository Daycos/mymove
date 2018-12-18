// eslint-disable-next-line no-unused-vars
import React from 'react';
import PropTypes from 'prop-types';
import './index.css';
import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faSpinner from '@fortawesome/fontawesome-free-solid/faSpinner';

//this is taken from https://designsystem.digital.gov/components/alerts/
const Alert = props => (
  <div className={`usa-alert usa-alert-${props.type}`}>
    <div className="usa-alert-body">
      <div className="body--heading">
        {props.type === 'loading' ? (
          <div className="heading--icon">
            <FontAwesomeIcon icon={faSpinner} spin pulse size="2x" />
          </div>
        ) : null}
        <div>
          {props.heading && <h3 className="usa-alert-heading">{props.heading}</h3>}
          <div className="usa-alert-text">{props.children}</div>
        </div>
      </div>
    </div>
  </div>
);

Alert.propTypes = {
  heading: PropTypes.string.isRequired,
  children: PropTypes.node,
  type: PropTypes.oneOf(['error', 'warning', 'info', 'success', 'loading']),
};
export default Alert;
