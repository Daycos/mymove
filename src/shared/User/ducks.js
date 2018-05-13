import * as Cookies from 'js-cookie';
import * as decode from 'jwt-decode';
import * as helpers from 'shared/ReduxHelpers';
import { GetLoggedInUser } from './api.js';

const GET_LOGGED_IN_USER = 'GET_LOGGED_IN_USER';

export const getUserTypes = helpers.generateAsyncActionTypes(
  GET_LOGGED_IN_USER,
);

const getLoggedInActions = helpers.generateAsyncActions(GET_LOGGED_IN_USER);
export const loadLoggedInUser = () => {
  return function(dispatch) {
    const userInfo = getUserInfo();
    if (!userInfo.isLoggedIn) return;
    dispatch(getLoggedInActions.start());
    GetLoggedInUser()
      .then(item => dispatch(getLoggedInActions.success(item)))
      .catch(error => dispatch(getLoggedInActions.error(error)));
  };
};

const generatedReducer = helpers.generateAsyncReducer(
  GET_LOGGED_IN_USER,
  u => ({ loggedInUser: u }),
);

export const loggedInUserReducer = (state, action) => {
  const mutatedState = generatedReducer(state, action);
  //we want the service member info in logged in user to be up to date.
  // In the long run we may want to change the server member reducer to work here
  switch (action.type) {
    case 'CREATE_SERVICE_MEMBER_SUCCESS':
    case 'UPDATE_SERVICE_MEMBER_SUCCESS':
    case 'GET_SERVICE_MEMBER_SUCCESS':
      return {
        ...mutatedState,
        loggedInUser: {
          ...mutatedState.loggedInUser,
          service_member: action.payload,
        },
      };
    default:
      return mutatedState;
  }
};

const loggedOutUser = {
  isLoggedIn: false,
  email: null,
  userId: null,
};

function getUserInfo() {
  const cookie = Cookies.get('user_session');
  if (!cookie) return loggedOutUser;
  const jwt = decode(cookie);
  return {
    email: jwt.email,
    userId: jwt.user_id,
    isLoggedIn: true,
  };
}

const userReducer = (state = getUserInfo(), action) => {
  return getUserInfo();
};

export default userReducer;
