import * as Cookies from 'js-cookie';
import * as decode from 'jwt-decode';
import * as helpers from 'shared/ReduxHelpers';
import { GetLoggedInUser } from './api.js';

const getLoggedInUserType = 'GET_LOGGED_IN_USER';

export const GET_LOGGED_IN_USER = helpers.generateAsyncActionTypes(
  getLoggedInUserType,
);

const getLoggedInActions = helpers.generateAsyncActions(getLoggedInUserType);
export const loadLoggedInUser = () => {
  return function(dispatch) {
    const userInfo = getUserInfo();
    if (!userInfo.isLoggedIn) return Promise.resolve();
    dispatch(getLoggedInActions.start());
    return GetLoggedInUser()
      .then(item => dispatch(getLoggedInActions.success(item)))
      .catch(error => dispatch(getLoggedInActions.error(error)));
  };
};

// the results of the api call will be handled by other reducers. This just lets us know app has loaded initial data
export const loggedInUserReducer = (state = {}, action) => {
  switch (action.type) {
    case GET_LOGGED_IN_USER.start:
      return {
        ...state,
        hasSucceeded: false,
        hasErrored: false,
        isLoading: true,
      };
    case GET_LOGGED_IN_USER.success:
      return {
        ...state,
        isLoading: false,
        hasErrored: false,
        hasSucceeded: true,
      };
    case GET_LOGGED_IN_USER.error:
      return {
        ...state,
        isLoading: false,
        hasErrored: true,
        hasSucceeded: false,
        error: action.error,
      };
    default:
      return state;
  }
};

const loggedOutUser = {
  isLoggedIn: false,
  email: null,
  userId: null,
};

function getUserInfo() {
  const cookie = Cookies.get('session_token');
  if (!cookie) return loggedOutUser;
  const jwt = decode(cookie);
  return {
    email: jwt.SessionValue.Email,
    userId: jwt.SessionValue.UserID,
    isLoggedIn: true,
  };
}

const userReducer = (state = getUserInfo(), action) => {
  return getUserInfo();
};

export default userReducer;
